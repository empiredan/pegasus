/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#pragma once

#include <array>
#include <cstddef>
#include <cstdint>
#include <memory>
#include <vector>

#include "replica/replica_base.h"
#include "utils/metrics.h"

namespace dsn {
class blob;

namespace apps {
class check_and_mutate_request;
class check_and_mutate_response;
class check_and_set_request;
class check_and_set_response;
class duplicate_request;
class duplicate_response;
class incr_request;
class incr_response;
class multi_put_request;
class multi_remove_request;
class multi_remove_response;
class update_request;
class update_response;
} // namespace apps
namespace replication {
class ingestion_request;
class ingestion_response;
} // namespace replication
} // namespace dsn

namespace pegasus::server {

// The context of an mutation to the database.
struct db_write_context
{
    // the mutation decree
    int64_t decree{0};

    // The time when this mutation is generated.
    // This is used to calculate the new timetag.
    uint64_t timestamp{0};

    // timetag of the remote write, 0 if it's not from remote.
    uint64_t remote_timetag{0};

    // Whether to compare the timetag of old value with the new write's.
    // - If true, it requires a read to the DB before write. If the old record has a larger timetag
    // than `remote_timetag`, the write will be ignored, otherwise it will be applied.
    // - If false, no overhead for the write but the eventual consistency on duplication
    // is not guaranteed.
    //
    // This is for duplicated write only. Because under casual consistency, the former
    // duplicated write must **happen before** the latest local write, regardless whether
    // its timestamp is larger. This relationship can be easily proved:
    // ```
    //      T1(put "a") > T2(duplicated put "b") > T3(put "b" in remote cluster)
    // ```
    // However write conflict may still result in inconsistent data in different clusters,
    // though those "versions" can all be considered as latest. User who requires consistency
    // can read from one main cluster instead of reading from multiple.
    bool verify_timetag{false};

    static inline db_write_context empty(int64_t d) { return create(d, 0); }

    // Creates a context for normal write.
    static inline db_write_context create(int64_t decree, uint64_t timestamp)
    {
        db_write_context ctx;
        ctx.decree = decree;
        ctx.timestamp = timestamp;
        return ctx;
    }

    // Creates a context for duplicated write.
    static inline db_write_context
    create_duplicate(int64_t decree, uint64_t remote_timetag, bool verify_timetag)
    {
        db_write_context ctx;
        ctx.decree = decree;
        ctx.remote_timetag = remote_timetag;
        ctx.verify_timetag = verify_timetag;
        return ctx;
    }

    bool is_duplicated_write() const { return remote_timetag > 0; }
};

class capacity_unit_calculator;
class pegasus_server_impl;

/// Handle the write requests.
/// As the signatures imply, this class is not responsible for replying the rpc,
/// the caller(pegasus_server_write) should do.
/// \see pegasus::server::pegasus_server_write::on_batched_write_requests
class pegasus_write_service : dsn::replication::replica_base
{
public:
    explicit pegasus_write_service(pegasus_server_impl *server);

    ~pegasus_write_service();

    // Write empty record.
    // See this document (https://pegasus.apache.org/zh/2018/03/07/last_flushed_decree.html)
    // to know why we must have empty write.
    int empty_put(int64_t decree);

    // Write MULTI_PUT record.
    int multi_put(const db_write_context &ctx,
                  const dsn::apps::multi_put_request &update,
                  dsn::apps::update_response &resp);

    // Write MULTI_REMOVE record.
    int multi_remove(int64_t decree,
                     const dsn::apps::multi_remove_request &update,
                     dsn::apps::multi_remove_response &resp);

    // Translate an INCR request into an idempotent PUT request. Only called by primary
    // replicas.
    int make_idempotent(const dsn::apps::incr_request &req,
                        dsn::apps::incr_response &err_resp,
                        std::vector<dsn::apps::update_request> &updates);

    // Write an idempotent INCR record (i.e. a PUT record) and reply to the client with INCR
    // response. Only called by primary replicas.
    int put(const db_write_context &ctx,
            const std::vector<dsn::apps::update_request> &updates,
            const dsn::apps::incr_request &req,
            dsn::apps::incr_response &resp);

    // Write a non-idempotent INCR record.
    int incr(int64_t decree, const dsn::apps::incr_request &update, dsn::apps::incr_response &resp);

    // Translate a CHECK_AND_SET request into an idempotent PUT request. Only called by
    // primary replicas.
    int make_idempotent(const dsn::apps::check_and_set_request &req,
                        dsn::apps::check_and_set_response &err_resp,
                        std::vector<dsn::apps::update_request> &updates);

    // Write an idempotent CHECK_AND_SET record (i.e. a PUT record) and reply to the client
    // with CHECK_AND_SET response. Only called by primary replicas.
    int put(const db_write_context &ctx,
            const std::vector<dsn::apps::update_request> &updates,
            const dsn::apps::check_and_set_request &req,
            dsn::apps::check_and_set_response &resp);

    // Write CHECK_AND_SET record.
    int check_and_set(int64_t decree,
                      const dsn::apps::check_and_set_request &update,
                      dsn::apps::check_and_set_response &resp);

    // Translate a CHECK_AND_MUTATE request into multiple idempotent PUT requests, which are
    // shared by both single-put and single-remove operations. Only called by primary replicas.
    int make_idempotent(const dsn::apps::check_and_mutate_request &req,
                        dsn::apps::check_and_mutate_response &err_resp,
                        std::vector<dsn::apps::update_request> &updates);

    // Write an idempotent CHECK_AND_MUTATE record (i.e. batched PUT records including both
    // single-put and single-remove operations) and reply to the client with CHECK_AND_MUTATE
    // response. Only called by primary replicas.
    int put(const db_write_context &ctx,
            const std::vector<dsn::apps::update_request> &updates,
            const dsn::apps::check_and_mutate_request &req,
            dsn::apps::check_and_mutate_response &resp);

    // Write CHECK_AND_MUTATE record.
    int check_and_mutate(int64_t decree,
                         const dsn::apps::check_and_mutate_request &update,
                         dsn::apps::check_and_mutate_response &resp);

    // Handles DUPLICATE duplicated from remote.
    int duplicate(int64_t decree,
                  const dsn::apps::duplicate_request &update,
                  dsn::apps::duplicate_response &resp);

    // Execute bulk load ingestion
    int ingest_files(int64_t decree,
                     const dsn::replication::ingestion_request &req,
                     dsn::replication::ingestion_response &resp);

    /// For batch write.

    // Prepare batch write.
    void batch_prepare(int64_t decree);

    // Add PUT record in batch write.
    // \returns rocksdb::Status::Code.
    // NOTE that `resp` should not be moved or freed while the batch is not committed.
    // When called by secondary replicas, this put request may be translated from an incr
    // request for idempotence.
    int batch_put(const db_write_context &ctx,
                  const dsn::apps::update_request &update,
                  dsn::apps::update_response &resp);

    // Add REMOVE record in batch write.
    // \returns rocksdb::Status::Code.
    // NOTE that `resp` should not be moved or freed while the batch is not committed.
    int batch_remove(int64_t decree, const dsn::blob &key, dsn::apps::update_response &resp);

    // Commit batch write.
    // \returns rocksdb::Status::Code.
    // NOTE that if the batch contains no updates, rocksdb::Status::kOk is returned.
    int batch_commit(int64_t decree);

    // Abort batch write.
    void batch_abort(int64_t decree, int err);

    void set_default_ttl(uint32_t ttl);

private:
    // Finish batch write with metrics such as latencies calculated and some states cleared.
    void batch_finish();

    // Used to store the batch size for each type of write into an array, see comments for
    // `_batch_sizes` for details.
    enum class batch_write_type : uint32_t
    {
        put = 0,
        remove,
        incr,
        check_and_set,
        check_and_mutate,
        COUNT,
    };

    // Read/change the batch size of the writes with the given type.
    uint32_t &batch_size(batch_write_type type)
    {
        return _batch_sizes.at(static_cast<uint32_t>(type));
    }

    uint32_t &put_batch_size() { return batch_size(batch_write_type::put); }

    uint32_t &remove_batch_size() { return batch_size(batch_write_type::remove); }

    friend class pegasus_write_service_test;
    friend class PegasusWriteServiceImplTest;
    friend class pegasus_server_write_test;
    friend class rocksdb_wrapper_test;

    pegasus_server_impl *_server;

    class impl;

    std::unique_ptr<impl> _impl;

    uint64_t _batch_start_time;

    capacity_unit_calculator *_cu_calculator;

    // To calculate the metrics such as the number of requests and latency for the writes
    // allowed to be batched, measure the size of requests in batch applied into RocksDB
    // for single put, single remove, incr, check_and_set and check_and_mutate, all of which
    // are contained in batch_write_type. In fact, incr, check_and_set and check_and_mutate
    // are not batched; the reason why they are contained in batch_write_type is because
    // all of them are not actually incr, check_and_set, or check_and_mutate themselves, but
    // rather single put operations that have been transformed from these operations to be
    // idempotent. Therefore, they essentially all appear in the form of single puts with an
    // extra field indicating what the original request is.
    //
    // Each request of single put, single remove, incr and check_and_set contains only one
    // write operation, while check_and_mutate may contain multiple operations of single
    // puts and removes.
    std::array<uint32_t, static_cast<size_t>(batch_write_type::COUNT)> _batch_sizes{};

    METRIC_VAR_DECLARE_counter(put_requests);
    METRIC_VAR_DECLARE_counter(multi_put_requests);
    METRIC_VAR_DECLARE_counter(remove_requests);
    METRIC_VAR_DECLARE_counter(multi_remove_requests);
    METRIC_VAR_DECLARE_counter(incr_requests);
    METRIC_VAR_DECLARE_counter(check_and_set_requests);
    METRIC_VAR_DECLARE_counter(check_and_mutate_requests);

    METRIC_VAR_DECLARE_percentile_int64(make_incr_idempotent_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(make_check_and_set_idempotent_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(make_check_and_mutate_idempotent_latency_ns);

    METRIC_VAR_DECLARE_percentile_int64(put_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(multi_put_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(remove_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(multi_remove_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(incr_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(check_and_set_latency_ns);
    METRIC_VAR_DECLARE_percentile_int64(check_and_mutate_latency_ns);

    METRIC_VAR_DECLARE_counter(dup_requests);
    METRIC_VAR_DECLARE_percentile_int64(dup_time_lag_ms);
    METRIC_VAR_DECLARE_counter(dup_lagging_writes);

    // TODO(wutao1): add metrics for failed rpc.
};

} // namespace pegasus::server
