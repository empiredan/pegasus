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

#ifdef DSN_USE_JEMALLOC

#include "utils/je_ctl.h"

#include <cstring>
#include <utility>
#include <vector>

#include <boost/algorithm/string/join.hpp>
#include <fmt/format.h>
#include <jemalloc/jemalloc.h>

#include "utils/api_utilities.h"
#include "utils/fmt_logging.h"
#include "utils/ports.h"
#include "utils/safe_strerror_posix.h"

#define RETURN_ARRAY_ELEM_BY_ENUM_TYPE(type, array)                                                \
    do {                                                                                           \
        const auto index = static_cast<size_t>(type);                                              \
        CHECK_LT(index, sizeof(array) / sizeof(array[0]));                                         \
        return array[index];                                                                       \
    } while (0);

namespace dsn {

namespace {

void je_stats_cb(void *opaque, const char *str)
{
    if (str == nullptr) {
        return;
    }

    auto stats = reinterpret_cast<std::string *>(opaque);
    auto avail_capacity = stats->capacity() - stats->size();
    auto len = strlen(str);
    if (len > avail_capacity) {
        len = avail_capacity;
    }

    stats->append(str, len);
}

void je_dump_malloc_stats(const char *opts, size_t buf_sz, std::string &stats)
{
    // Avoid malloc in callback.
    stats.reserve(buf_sz);

    ::malloc_stats_print(je_stats_cb, &stats, opts);
}

const char *je_stats_type_to_opts(je_stats_type type)
{
    static const char *opts_map[] = {
        "gmdablxe", "mdablxe", "gblxe", "",
    };

    RETURN_ARRAY_ELEM_BY_ENUM_TYPE(type, opts_map);
}

size_t je_stats_type_to_default_buf_sz(je_stats_type type)
{
    static const size_t buf_sz_map[] = {
        2 * 1024, 4 * 1024, 1024 * 1024, 2 * 1024 * 1024,
    };

    RETURN_ARRAY_ELEM_BY_ENUM_TYPE(type, buf_sz_map);
}

bool je_check_err(const char *action, const int err, std::string *msg = nullptr)
{
    std::string my_msg;
    if (dsn_likely(err == 0)) {
        my_msg = fmt::format("{} successfully", action);
    } else {
        my_msg = fmt::format(
            "failed to {}: errno={}, message={}", action, err, dsn::utils::safe_strerror(err));
    }

    LOG_INFO_F("<jemalloc> {}", my_msg);

    if (msg != nullptr) {
        *msg = std::move(my_msg);
    }

    return err == 0;
}

inline bool je_check_get_err(const char *name, int err, std::string *msg = nullptr)
{
    std::string action(fmt::format("get {}", name));
    return je_check_err(action.c_str(), err, msg);
}

template <typename T>
inline bool je_get_val(const char *name, T &val, std::string *msg = nullptr)
{
    size_t sz = sizeof(val);
    int je_ret = ::mallctl(name, &val, &sz, nullptr, 0);
    return je_check_get_err(name, je_ret, msg);
}

template <typename T>
inline bool
je_check_set_err(const char *name, const T &val, int err, std::string *err_msg = nullptr)
{
    std::string action(fmt::format("set {} to {}", name, val));
    return je_check_err(action.c_str(), err, err_msg);
}

template <typename T>
inline bool je_set_val(const char *name, const T &val, std::string *err_msg = nullptr)
{
    auto new_val = reinterpret_cast<void *>(const_cast<T *>(&val));
    int je_ret = ::mallctl(name, nullptr, nullptr, new_val, sizeof(val));
    return je_check_set_err(name, val, je_ret, err_msg);
}

inline bool je_get_str(const char *name, std::string &val, std::string *msg = nullptr)
{
    const char *p = nullptr;
    size_t sz = sizeof(p);
    int je_ret = ::mallctl(name, reinterpret_cast<void *>(&p), &sz, nullptr, 0);
    if (!je_check_get_err(name, je_ret, msg)) {
        return false;
    }

    val = p;
    return true;
}

inline bool je_set_str(const char *name, const char *val, std::string *err_msg = nullptr)
{
    void *p = nullptr;
    size_t sz = 0;
    if (val != nullptr) {
        p = reinterpret_cast<void *>(&val);
        sz = sizeof(val);
    }

    int je_ret = ::mallctl(name, nullptr, nullptr, p, sz);
    return je_check_set_err(name, val == nullptr ? "nullptr" : val, je_ret, err_msg);
}

inline bool je_is_prof_enabled(bool &enabled, std::string *err_msg)
{
    return je_get_val("opt.prof", enabled, err_msg);
}

inline bool je_is_prof_active(bool &active, std::string *err_msg)
{
    return je_get_val("opt.prof_active", active, err_msg);
}

inline bool je_is_writable_prof_active(bool &active, std::string *err_msg)
{
    return je_get_val("prof.active", active, err_msg);
}

#define CHECK_IF_PROF_ENABLED(err_msg)                                                             \
    do {                                                                                           \
        bool enabled = false;                                                                      \
        if (!je_is_prof_enabled(enabled, err_msg)) {                                               \
            return false;                                                                          \
        }                                                                                          \
        if (!enabled) {                                                                            \
            *err_msg = "<jemalloc> prof is disabled now, enable it by "                            \
                       "`export MALLOC_CONF=\"prof:true,prof_prefix:...\"`";                       \
            return false;                                                                          \
        }                                                                                          \
    } while (0)

#define CHECK_IF_PROF_ACTIVE(err_msg)                                                              \
    do {                                                                                           \
        bool active = false;                                                                       \
        if (!je_is_prof_active(active, err_msg)) {                                                 \
            return false;                                                                          \
        }                                                                                          \
        if (!active) {                                                                             \
            *err_msg = "<jemalloc> prof is not active now, activate it first";                     \
            return false;                                                                          \
        }                                                                                          \
    } while (0)

inline bool je_set_prof_active(bool active, std::string *err_msg)
{
    CHECK_IF_PROF_ENABLED(err_msg);
    return je_set_val("prof.active", active, err_msg);
}

inline bool je_get_prof_prefix(std::string &prefix, std::string *err_msg)
{
    return je_get_str("opt.prof_prefix", prefix, err_msg);
}

} // anonymous namespace

std::string get_all_je_stats_types_str()
{
    std::vector<std::string> names;
    for (size_t i = 0; i < static_cast<size_t>(je_stats_type::COUNT); ++i) {
        names.emplace_back(enum_to_string(static_cast<je_stats_type>(i)));
    }
    return boost::join(names, " | ");
}

void je_dump_stats(je_stats_type type, size_t buf_sz, std::string &stats)
{
    je_dump_malloc_stats(je_stats_type_to_opts(type), buf_sz, stats);
}

void je_dump_stats(je_stats_type type, std::string &stats)
{
    je_dump_stats(type, je_stats_type_to_default_buf_sz(type), stats);
}

bool je_get_prof_status(std::string &info)
{
    std::string err_msg;

    bool enabled;
    if (!je_is_prof_enabled(enabled, &err_msg)) {
        info = err_msg;
        return false;
    }
    info += "current status of jemalloc profile:";
    info += fmt::format("\nopt.prof: {}", enabled);

    if (!je_is_prof_active(enabled, &err_msg)) {
        info = err_msg;
        return false;
    }
    info += fmt::format("\nopt.prof_active: {}", enabled);

    if (!je_is_writable_prof_active(enabled, &err_msg)) {
        info = err_msg;
        return false;
    }
    info += fmt::format("\nprof.active: {}", enabled);

    std::string prefix;
    if (!je_get_prof_prefix(prefix, &err_msg)) {
        info = err_msg;
        return false;
    }
    info += fmt::format("\nopt.prof_prefix: {}", prefix);

    return true;
}

bool je_activate_prof(std::string *err_msg) { return je_set_prof_active(true, err_msg); }

bool je_deactivate_prof(std::string *err_msg) { return je_set_prof_active(false, err_msg); }

bool je_dump_prof(const char *path, std::string *err_msg)
{
    CHECK_IF_PROF_ENABLED(err_msg);
    CHECK_IF_PROF_ACTIVE(err_msg);
    return je_set_str("prof.dump", path, err_msg);
}

} // namespace dsn

#endif // DSN_USE_JEMALLOC
