// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

#pragma once

#include "utils/errors.h"
#include "meta_admin_types.h"
#include "partition_split_types.h"
#include "duplication_types.h"
#include "bulk_load_types.h"
#include "backup_types.h"
#include "consensus_types.h"
#include "replica_admin_types.h"
#include "replica/replica_base.h"
#include "runtime/pipeline.h"

namespace dsn {
namespace replication {

/// \brief Each of the mutation is a tuple made up of
/// <timestamp, task_code, dsn::blob, decree, index>.
/// dsn::blob is the content of the mutation.
typedef std::tuple<uint64_t, task_code, blob, int64_t, int64_t> mutation_tuple;

// Elements in mutation_tuple_set are sorted by the decree of a mutation and the index
// of an update in this mutation.
struct mutation_tuple_cmp
{
    inline bool operator()(const mutation_tuple &lhs, const mutation_tuple &rhs) const
    {
        // Compare the decrees of two mutations.
        const int64_t lhs_decree = std::get<3>(lhs);
        const int64_t rhs_decree = std::get<3>(rhs);
        if (lhs_decree < rhs_decree) {
            return true;
        }

        if (lhs_decree > rhs_decree) {
            return false;
        }

        // Compare the indexes of 2 updates in the same mutation.
        return std::get<4>(lhs) < std::get<4>(rhs);
    }
};
typedef std::set<mutation_tuple, mutation_tuple_cmp> mutation_tuple_set;

/// \brief This is an interface for handling the mutation logs intended to
/// be duplicated to remote cluster.
/// \see dsn::replication::replica_duplicator
class mutation_duplicator : public replica_base
{
public:
    typedef std::function<void(size_t /*total_shipped_size*/)> callback;

    /// Duplicate the provided mutations to the remote cluster.
    /// The implementation must be non-blocking.
    ///
    /// \param cb: Call it when all the given mutations were sent successfully
    virtual void duplicate(mutation_tuple_set mutations, callback cb) = 0;

    // Singleton creator of mutation_duplicator.
    static std::function<std::unique_ptr<mutation_duplicator>(
        replica_base *, absl::string_view /*remote cluster*/, absl::string_view /*app name*/)>
        creator;

    explicit mutation_duplicator(replica_base *r) : replica_base(r) {}

    virtual ~mutation_duplicator() = default;

    void set_task_environment(pipeline::environment *env) { _env = *env; }

protected:
    friend class replica_duplicator_test;

    pipeline::environment _env;
};

inline std::unique_ptr<mutation_duplicator> new_mutation_duplicator(
    replica_base *r, absl::string_view remote_cluster_address, absl::string_view app)
{
    return mutation_duplicator::creator(r, remote_cluster_address, app);
}

} // namespace replication
} // namespace dsn
