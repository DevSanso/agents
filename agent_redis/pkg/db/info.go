package db

import (
	"agent_redis/pkg/protos"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseRedisMemoryInfo(data []string) (*protos.RedisMemoryInfo, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}

	var redisMemoryInfo = &protos.RedisMemoryInfo{}
	for _, line := range data {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			return nil, errors.New("invalid data format :" + line)
		}

		field, value := strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1])
		var err error

		switch field {
		case "used_memory":
			redisMemoryInfo.UsedMemory, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_human":
			redisMemoryInfo.UsedMemoryHuman = value
		case "used_memory_rss":
			redisMemoryInfo.UsedMemoryRss, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_rss_human":
			redisMemoryInfo.UsedMemoryRssHuman = value
		case "used_memory_peak":
			redisMemoryInfo.UsedMemoryPeak, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_peak_human":
			redisMemoryInfo.UsedMemoryPeakHuman = value
		case "used_memory_peak_perc":
			redisMemoryInfo.UsedMemoryPeakPerc = value
		case "used_memory_overhead":
			redisMemoryInfo.UsedMemoryOverhead, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_startup":
			redisMemoryInfo.UsedMemoryStartup, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_dataset":
			redisMemoryInfo.UsedMemoryDataset, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_dataset_perc":
			redisMemoryInfo.UsedMemoryDatasetPerc = value
		case "allocator_allocated":
			redisMemoryInfo.AllocatorAllocated, err = strconv.ParseInt(value, 10, 64)
		case "allocator_active":
			redisMemoryInfo.AllocatorActive, err = strconv.ParseInt(value, 10, 64)
		case "allocator_resident":
			redisMemoryInfo.AllocatorResident, err = strconv.ParseInt(value, 10, 64)
		case "total_system_memory":
			redisMemoryInfo.TotalSystemMemory, err = strconv.ParseInt(value, 10, 64)
		case "total_system_memory_human":
			redisMemoryInfo.TotalSystemMemoryHuman = value
		case "used_memory_lua":
			redisMemoryInfo.UsedMemoryLua, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_vm_eval":
			redisMemoryInfo.UsedMemoryVmEval, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_lua_human":
			redisMemoryInfo.UsedMemoryLuaHuman = value
		case "used_memory_scripts_eval":
			redisMemoryInfo.UsedMemoryScriptsEval, err = strconv.ParseInt(value, 10, 64)
		case "number_of_cached_scripts":
			redisMemoryInfo.NumberOfCachedScripts, err = strconv.ParseInt(value, 10, 64)
		case "number_of_functions":
			redisMemoryInfo.NumberOfFunctions, err = strconv.ParseInt(value, 10, 64)
		case "number_of_libraries":
			redisMemoryInfo.NumberOfLibraries, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_vm_functions":
			redisMemoryInfo.UsedMemoryVmFunctions, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_vm_total":
			redisMemoryInfo.UsedMemoryVmTotal, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_vm_total_human":
			redisMemoryInfo.UsedMemoryVmTotalHuman = value
		case "used_memory_functions":
			redisMemoryInfo.UsedMemoryFunctions, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_scripts":
			redisMemoryInfo.UsedMemoryScripts, err = strconv.ParseInt(value, 10, 64)
		case "used_memory_scripts_human":
			redisMemoryInfo.UsedMemoryScriptsHuman = value
		case "maxmemory":
			redisMemoryInfo.MaxMemory, err = strconv.ParseInt(value, 10, 64)
		case "maxmemory_human":
			redisMemoryInfo.MaxMemoryHuman = value
		case "maxmemory_policy":
			redisMemoryInfo.MaxMemoryPolicy = value
		case "allocator_frag_ratio":
			redisMemoryInfo.AllocatorFragRatio, err = strconv.ParseFloat(value, 64)
		case "allocator_frag_bytes":
			redisMemoryInfo.AllocatorFragBytes, err = strconv.ParseInt(value, 10, 64)
		case "allocator_rss_ratio":
			redisMemoryInfo.AllocatorRssRatio, err = strconv.ParseFloat(value, 64)
		case "allocator_rss_bytes":
			redisMemoryInfo.AllocatorRssBytes, err = strconv.ParseInt(value, 10, 64)
		case "rss_overhead_ratio":
			redisMemoryInfo.RssOverheadRatio, err = strconv.ParseFloat(value, 64)
		case "rss_overhead_bytes":
			redisMemoryInfo.RssOverheadBytes, err = strconv.ParseInt(value, 10, 64)
		case "mem_fragmentation_ratio":
			redisMemoryInfo.MemFragmentationRatio, err = strconv.ParseFloat(value, 64)
		case "mem_fragmentation_bytes":
			redisMemoryInfo.MemFragmentationBytes, err = strconv.ParseInt(value, 10, 64)
		case "mem_not_counted_for_evict":
			redisMemoryInfo.MemNotCountedForEvict, err = strconv.ParseInt(value, 10, 64)
		case "mem_replication_backlog":
			redisMemoryInfo.MemReplicationBacklog, err = strconv.ParseInt(value, 10, 64)
		case "mem_total_replication_buffers":
			redisMemoryInfo.MemTotalReplicationBuffers, err = strconv.ParseInt(value, 10, 64)
		case "mem_clients_slaves":
			redisMemoryInfo.MemClientsSlaves, err = strconv.ParseInt(value, 10, 64)
		case "mem_clients_normal":
			redisMemoryInfo.MemClientsNormal, err = strconv.ParseInt(value, 10, 64)
		case "mem_cluster_links":
			redisMemoryInfo.MemClusterLinks, err = strconv.ParseInt(value, 10, 64)
		case "mem_aof_buffer":
			redisMemoryInfo.MemAofBuffer, err = strconv.ParseInt(value, 10, 64)
		case "mem_allocator":
			redisMemoryInfo.MemAllocator = value
		case "active_defrag_running":
			redisMemoryInfo.ActiveDefragRunning, err = strconv.ParseInt(value, 10, 64)
		case "lazyfree_pending_objects":
			redisMemoryInfo.LazyFreePendingObjects, err = strconv.ParseInt(value, 10, 64)
		case "lazyfreed_objects":
			redisMemoryInfo.LazyFreedObjects, err = strconv.ParseInt(value, 10, 64)
		default:
			return nil, errors.New("unknown field: " + field)
		}

		if err != nil {
			return nil, err
		}
	}

	return redisMemoryInfo, nil
}

func (rcc *redisClientCommander) InfoMemory(ctx context.Context) (*protos.RedisMemoryInfo, error) {
	cmd := rcc.client.Do(ctx, "info", "memory")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	s := cmd.String()
	split := strings.Split(s, "\n")

	return parseRedisMemoryInfo(split[1:])
}

func parseRedisStatsInfo(data map[string]string) (*protos.RedisStatsInfo, error) {
	var redisStatsInfo = &protos.RedisStatsInfo{}

	for field, value := range data {
		switch field {
		case "total_connections_received":
			redisStatsInfo.TotalConnectionsReceived, _ = strconv.ParseInt(value, 10, 64)
		case "total_commands_processed":
			redisStatsInfo.TotalCommandsProcessed, _ = strconv.ParseInt(value, 10, 64)
		case "instantaneous_ops_per_sec":
			redisStatsInfo.InstantaneousOpsPerSec, _ = strconv.ParseInt(value, 10, 64)
		case "total_net_input_bytes":
			redisStatsInfo.TotalNetInputBytes, _ = strconv.ParseInt(value, 10, 64)
		case "total_net_output_bytes":
			redisStatsInfo.TotalNetOutputBytes, _ = strconv.ParseInt(value, 10, 64)
		case "total_net_repl_input_bytes":
			redisStatsInfo.TotalNetReplInputBytes, _ = strconv.ParseInt(value, 10, 64)
		case "total_net_repl_output_bytes":
			redisStatsInfo.TotalNetReplOutputBytes, _ = strconv.ParseInt(value, 10, 64)
		case "instantaneous_input_kbps":
			redisStatsInfo.InstantaneousInputKbps, _ = strconv.ParseFloat(value, 64)
		case "instantaneous_output_kbps":
			redisStatsInfo.InstantaneousOutputKbps, _ = strconv.ParseFloat(value, 64)
		case "instantaneous_input_repl_kbps":
			redisStatsInfo.InstantaneousInputReplKbps, _ = strconv.ParseFloat(value, 64)
		case "instantaneous_output_repl_kbps":
			redisStatsInfo.InstantaneousOutputReplKbps, _ = strconv.ParseFloat(value, 64)
		case "rejected_connections":
			redisStatsInfo.RejectedConnections, _ = strconv.ParseInt(value, 10, 64)
		case "sync_full":
			redisStatsInfo.SyncFull, _ = strconv.ParseInt(value, 10, 64)
		case "sync_partial_ok":
			redisStatsInfo.SyncPartialOk, _ = strconv.ParseInt(value, 10, 64)
		case "sync_partial_err":
			redisStatsInfo.SyncPartialErr, _ = strconv.ParseInt(value, 10, 64)
		case "expired_keys":
			redisStatsInfo.ExpiredKeys, _ = strconv.ParseInt(value, 10, 64)
		case "expired_stale_perc":
			redisStatsInfo.ExpiredStalePerc, _ = strconv.ParseFloat(value, 64)
		case "expired_time_cap_reached_count":
			redisStatsInfo.ExpiredTimeCapReachedCount, _ = strconv.ParseInt(value, 10, 64)
		case "expire_cycle_cpu_milliseconds":
			redisStatsInfo.ExpireCycleCpuMilliseconds, _ = strconv.ParseInt(value, 10, 64)
		case "evicted_keys":
			redisStatsInfo.EvictedKeys, _ = strconv.ParseInt(value, 10, 64)
		case "evicted_clients":
			redisStatsInfo.EvictedClients, _ = strconv.ParseInt(value, 10, 64)
		case "total_eviction_exceeded_time":
			redisStatsInfo.TotalEvictionExceededTime, _ = strconv.ParseInt(value, 10, 64)
		case "current_eviction_exceeded_time":
			redisStatsInfo.CurrentEvictionExceededTime, _ = strconv.ParseInt(value, 10, 64)
		case "keyspace_hits":
			redisStatsInfo.KeyspaceHits, _ = strconv.ParseInt(value, 10, 64)
		case "keyspace_misses":
			redisStatsInfo.KeyspaceMisses, _ = strconv.ParseInt(value, 10, 64)
		case "pubsub_channels":
			redisStatsInfo.PubsubChannels, _ = strconv.ParseInt(value, 10, 64)
		case "pubsub_patterns":
			redisStatsInfo.PubsubPatterns, _ = strconv.ParseInt(value, 10, 64)
		case "pubsubshard_channels":
			redisStatsInfo.PubsubShardChannels, _ = strconv.ParseInt(value, 10, 64)
		case "latest_fork_usec":
			redisStatsInfo.LatestForkUsec, _ = strconv.ParseInt(value, 10, 64)
		case "total_forks":
			redisStatsInfo.TotalForks, _ = strconv.ParseInt(value, 10, 64)
		case "migrate_cached_sockets":
			redisStatsInfo.MigrateCachedSockets, _ = strconv.ParseInt(value, 10, 64)
		case "slave_expires_tracked_keys":
			redisStatsInfo.SlaveExpiresTrackedKeys, _ = strconv.ParseInt(value, 10, 64)
		case "active_defrag_hits":
			redisStatsInfo.ActiveDefragHits, _ = strconv.ParseInt(value, 10, 64)
		case "active_defrag_misses":
			redisStatsInfo.ActiveDefragMisses, _ = strconv.ParseInt(value, 10, 64)
		case "active_defrag_key_hits":
			redisStatsInfo.ActiveDefragKeyHits, _ = strconv.ParseInt(value, 10, 64)
		case "active_defrag_key_misses":
			redisStatsInfo.ActiveDefragKeyMisses, _ = strconv.ParseInt(value, 10, 64)
		case "total_active_defrag_time":
			redisStatsInfo.TotalActiveDefragTime, _ = strconv.ParseInt(value, 10, 64)
		case "current_active_defrag_time":
			redisStatsInfo.CurrentActiveDefragTime, _ = strconv.ParseInt(value, 10, 64)
		case "tracking_total_keys":
			redisStatsInfo.TrackingTotalKeys, _ = strconv.ParseInt(value, 10, 64)
		case "tracking_total_items":
			redisStatsInfo.TrackingTotalItems, _ = strconv.ParseInt(value, 10, 64)
		case "tracking_total_prefixes":
			redisStatsInfo.TrackingTotalPrefixes, _ = strconv.ParseInt(value, 10, 64)
		case "unexpected_error_replies":
			redisStatsInfo.UnexpectedErrorReplies, _ = strconv.ParseInt(value, 10, 64)
		case "total_error_replies":
			redisStatsInfo.TotalErrorReplies, _ = strconv.ParseInt(value, 10, 64)
		case "dump_payload_sanitizations":
			redisStatsInfo.DumpPayloadSanitizations, _ = strconv.ParseInt(value, 10, 64)
		case "total_reads_processed":
			redisStatsInfo.TotalReadsProcessed, _ = strconv.ParseInt(value, 10, 64)
		case "total_writes_processed":
			redisStatsInfo.TotalWritesProcessed, _ = strconv.ParseInt(value, 10, 64)
		case "io_threaded_reads_processed":
			redisStatsInfo.IoThreadedReadsProcessed, _ = strconv.ParseInt(value, 10, 64)
		case "io_threaded_writes_processed":
			redisStatsInfo.IoThreadedWritesProcessed, _ = strconv.ParseInt(value, 10, 64)
		case "reply_buffer_shrinks":
			redisStatsInfo.ReplyBufferShrinks, _ = strconv.ParseInt(value, 10, 64)
		case "reply_buffer_expands":
			redisStatsInfo.ReplyBufferExpands, _ = strconv.ParseInt(value, 10, 64)
		case "eventloop_cycles":
			redisStatsInfo.EventloopCycles, _ = strconv.ParseInt(value, 10, 64)
		case "eventloop_duration_sum":
			redisStatsInfo.EventloopDurationSum, _ = strconv.ParseInt(value, 10, 64)
		case "eventloop_duration_cmd_sum":
			redisStatsInfo.EventloopDurationCmdSum, _ = strconv.ParseInt(value, 10, 64)
		case "instantaneous_eventloop_cycles_per_sec":
			redisStatsInfo.InstantaneousEventloopCyclesPerSec, _ = strconv.ParseInt(value, 10, 64)
		case "instantaneous_eventloop_duration_usec":
			redisStatsInfo.InstantaneousEventloopDurationUsec, _ = strconv.ParseInt(value, 10, 64)
		case "acl_access_denied_auth":
			redisStatsInfo.AclAccessDeniedAuth, _ = strconv.ParseInt(value, 10, 64)
		case "acl_access_denied_cmd":
			redisStatsInfo.AclAccessDeniedCmd, _ = strconv.ParseInt(value, 10, 64)
		case "acl_access_denied_key":
			redisStatsInfo.AclAccessDeniedKey, _ = strconv.ParseInt(value, 10, 64)
		case "acl_access_denied_channel":
			redisStatsInfo.AclAccessDeniedChannel, _ = strconv.ParseInt(value, 10, 64)
		default:
			// 알 수 없는 필드는 무시
		}
	}

	return redisStatsInfo, nil
}
func (rcc *redisClientCommander) InfoStat(ctx context.Context) (*protos.RedisStatsInfo, error) {
	cmd := rcc.client.InfoMap(ctx, "stats")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return parseRedisStatsInfo(cmd.Val()["stats"])
}

// 문자열을 CPUInfo 구조체로 변환하는 함수
func parseCPUInfo(data []string) (*protos.RedisCpuInfo, error) {
	var cpuInfo = &protos.RedisCpuInfo{}

	for _, line := range data {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid data format: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value for key %s: %v", key, err)
		}

		switch key {
		case "used_cpu_sys":
			cpuInfo.UsedCPUSys = value
		case "used_cpu_user":
			cpuInfo.UsedCPUUser = value
		case "used_cpu_sys_children":
			cpuInfo.UsedCPUSysChildren = value
		case "used_cpu_user_children":
			cpuInfo.UsedCPUUserChildren = value
		case "used_cpu_sys_main_thread":
			cpuInfo.UsedCPUSysMainThread = value
		case "used_cpu_user_main_thread":
			cpuInfo.UsedCPUUserMainThread = value
		default:
			return nil, fmt.Errorf("unknown key: %s", key)
		}
	}

	return cpuInfo, nil
}

func (rcc *redisClientCommander) InfoCpu(ctx context.Context) (*protos.RedisCpuInfo, error) {
	cmd := rcc.client.Do(ctx, "info", "cpu")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	s := cmd.String()
	split := strings.Split(s, "\n")

	return parseCPUInfo(split[1:])
}
