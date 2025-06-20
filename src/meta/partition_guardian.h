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

#include <stdint.h>
#include <map>
#include <memory>
#include <set>
#include <string>
#include <utility>
#include <vector>

#include "common/gpid.h"
#include "dsn.layer2_types.h"
#include "meta_admin_types.h"
#include "meta_data.h"
#include "rpc/rpc_host_port.h"
#include "utils/command_manager.h"
#include "utils/ports.h"
#include "utils/zlocks.h"

namespace dsn::replication {

class meta_service;

class partition_guardian
{
public:
    template <typename T>
    static partition_guardian *create(meta_service *svc)
    {
        return new T(svc);
    }
    typedef partition_guardian *(*factory)(meta_service *svc);

    explicit partition_guardian(meta_service *svc);
    virtual ~partition_guardian() = default;

    virtual pc_status
    cure(meta_view view, const dsn::gpid &gpid, configuration_proposal_action &action);
    void reconfig(meta_view view, const configuration_update_request &request);
    void register_ctrl_commands();
    void get_ddd_partitions(const gpid &pid, std::vector<ddd_partition_info> &partitions) const;
    void clear_ddd_partitions()
    {
        zauto_lock l(_ddd_partitions_lock);
        _ddd_partitions.clear();
    }

private:
    bool
    from_proposals(meta_view &view, const dsn::gpid &gpid, configuration_proposal_action &action);
    pc_status on_missing_primary(meta_view &view, const dsn::gpid &gpid);
    pc_status on_missing_secondary(meta_view &view, const dsn::gpid &gpid);
    pc_status on_redundant_secondary(meta_view &view, const dsn::gpid &gpid);
    // if a proposal is generated by cure, meta will record the POSSIBLE PARTITION COUNT
    // IN FUTURE of a node with module "newly_partitions".
    // the side effect should be eliminated when a proposal is finished, no matter
    // successfully or unsuccessfully
    void finish_cure_proposal(meta_view &view,
                              const dsn::gpid &gpid,
                              const configuration_proposal_action &action);
    std::string ctrl_assign_secondary_black_list(const std::vector<std::string> &args);

    void set_ddd_partition(ddd_partition_info &&partition)
    {
        zauto_lock l(_ddd_partitions_lock);
        _ddd_partitions[partition.config.pid] = std::move(partition);
    }

    bool in_black_list(const dsn::host_port &hp)
    {
        dsn::zauto_read_lock l(_black_list_lock);
        return _assign_secondary_black_list.count(hp) != 0;
    }

    meta_service *_svc;

    mutable zlock _ddd_partitions_lock; // [
    std::map<gpid, ddd_partition_info> _ddd_partitions;
    // ]

    // NOTICE: the command handler is called in THREADPOOL_DEFAULT
    // but when adding secondary, the black list is accessed in THREADPOOL_META_STATE
    // so we need a lock to protect it
    dsn::zrwlock_nr _black_list_lock; // [
    std::set<dsn::host_port> _assign_secondary_black_list;
    // ]

    std::vector<std::unique_ptr<command_deregister>> _cmds;
    int64_t _replica_assign_delay_ms_for_dropouts;

    friend class meta_partition_guardian_test;

    DISALLOW_COPY_AND_ASSIGN(partition_guardian);
    DISALLOW_MOVE_AND_ASSIGN(partition_guardian);
};

} // namespace dsn::replication
