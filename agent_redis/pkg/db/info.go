package db

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type RedisCpuInfo struct {
	UsedCPUSys            float64
	UsedCPUUser           float64
	UsedCPUSysChildren    float64
	UsedCPUUserChildren   float64
	UsedCPUSysMainThread  float64
	UsedCPUUserMainThread float64
}

type RedisStatsInfo struct {
	TotalConnectionsReceived           int
	TotalCommandsProcessed             int
	InstantaneousOpsPerSec             int
	TotalNetInputBytes                 int
	TotalNetOutputBytes                int
	TotalNetReplInputBytes             int
	TotalNetReplOutputBytes            int
	InstantaneousInputKbps             float64
	InstantaneousOutputKbps            float64
	InstantaneousInputReplKbps         float64
	InstantaneousOutputReplKbps        float64
	RejectedConnections                int
	SyncFull                           int
	SyncPartialOK                      int
	SyncPartialErr                     int
	ExpiredKeys                        int
	ExpiredStalePerc                   float64
	ExpiredTimeCapReachedCount         int
	ExpireCycleCPUMilliseconds         int
	EvictedKeys                        int
	EvictedClients                     int
	TotalEvictionExceededTime          int
	CurrentEvictionExceededTime        int
	KeyspaceHits                       int
	KeyspaceMisses                     int
	PubsubChannels                     int
	PubsubPatterns                     int
	PubsubShardChannels                int
	LatestForkUsec                     int
	TotalForks                         int
	MigrateCachedSockets               int
	SlaveExpiresTrackedKeys            int
	ActiveDefragHits                   int
	ActiveDefragMisses                 int
	ActiveDefragKeyHits                int
	ActiveDefragKeyMisses              int
	TotalActiveDefragTime              int
	CurrentActiveDefragTime            int
	TrackingTotalKeys                  int
	TrackingTotalItems                 int
	TrackingTotalPrefixes              int
	UnexpectedErrorReplies             int
	TotalErrorReplies                  int
	DumpPayloadSanitizations           int
	TotalReadsProcessed                int
	TotalWritesProcessed               int
	IOThreadedReadsProcessed           int
	IOThreadedWritesProcessed          int
	ReplyBufferShrinks                 int
	ReplyBufferExpands                 int
	EventloopCycles                    int
	EventloopDurationSum               int
	EventloopDurationCmdSum            int
	InstantaneousEventloopCyclesPerSec int
	InstantaneousEventloopDurationUsec int
	ACLAccessDeniedAuth                int
	ACLAccessDeniedCmd                 int
	ACLAccessDeniedKey                 int
	ACLAccessDeniedChannel             int
}

func parseRedisStatsInfo(data map[string]string) (*RedisStatsInfo, error) {
	var redisStatsInfo = &RedisStatsInfo{}

	for field, value := range data {
		switch field {
		case "total_connections_received":
			redisStatsInfo.TotalConnectionsReceived, _ = strconv.Atoi(value)
		case "total_commands_processed":
			redisStatsInfo.TotalCommandsProcessed, _ = strconv.Atoi(value)
		case "instantaneous_ops_per_sec":
			redisStatsInfo.InstantaneousOpsPerSec, _ = strconv.Atoi(value)
		case "total_net_input_bytes":
			redisStatsInfo.TotalNetInputBytes, _ = strconv.Atoi(value)
		case "total_net_output_bytes":
			redisStatsInfo.TotalNetOutputBytes, _ = strconv.Atoi(value)
		case "total_net_repl_input_bytes":
			redisStatsInfo.TotalNetReplInputBytes, _ = strconv.Atoi(value)
		case "total_net_repl_output_bytes":
			redisStatsInfo.TotalNetReplOutputBytes, _ = strconv.Atoi(value)
		case "instantaneous_input_kbps":
			redisStatsInfo.InstantaneousInputKbps, _ = strconv.ParseFloat(value, 64)
		case "instantaneous_output_kbps":
			redisStatsInfo.InstantaneousOutputKbps, _ = strconv.ParseFloat(value, 64)
		case "instantaneous_input_repl_kbps":
			redisStatsInfo.InstantaneousInputReplKbps, _ = strconv.ParseFloat(value, 64)
		case "instantaneous_output_repl_kbps":
			redisStatsInfo.InstantaneousOutputReplKbps, _ = strconv.ParseFloat(value, 64)
		case "rejected_connections":
			redisStatsInfo.RejectedConnections, _ = strconv.Atoi(value)
		case "sync_full":
			redisStatsInfo.SyncFull, _ = strconv.Atoi(value)
		case "sync_partial_ok":
			redisStatsInfo.SyncPartialOK, _ = strconv.Atoi(value)
		case "sync_partial_err":
			redisStatsInfo.SyncPartialErr, _ = strconv.Atoi(value)
		case "expired_keys":
			redisStatsInfo.ExpiredKeys, _ = strconv.Atoi(value)
		case "expired_stale_perc":
			redisStatsInfo.ExpiredStalePerc, _ = strconv.ParseFloat(value, 64)
		case "expired_time_cap_reached_count":
			redisStatsInfo.ExpiredTimeCapReachedCount, _ = strconv.Atoi(value)
		case "expire_cycle_cpu_milliseconds":
			redisStatsInfo.ExpireCycleCPUMilliseconds, _ = strconv.Atoi(value)
		case "evicted_keys":
			redisStatsInfo.EvictedKeys, _ = strconv.Atoi(value)
		case "evicted_clients":
			redisStatsInfo.EvictedClients, _ = strconv.Atoi(value)
		case "total_eviction_exceeded_time":
			redisStatsInfo.TotalEvictionExceededTime, _ = strconv.Atoi(value)
		case "current_eviction_exceeded_time":
			redisStatsInfo.CurrentEvictionExceededTime, _ = strconv.Atoi(value)
		case "keyspace_hits":
			redisStatsInfo.KeyspaceHits, _ = strconv.Atoi(value)
		case "keyspace_misses":
			redisStatsInfo.KeyspaceMisses, _ = strconv.Atoi(value)
		case "pubsub_channels":
			redisStatsInfo.PubsubChannels, _ = strconv.Atoi(value)
		case "pubsub_patterns":
			redisStatsInfo.PubsubPatterns, _ = strconv.Atoi(value)
		case "pubsubshard_channels":
			redisStatsInfo.PubsubShardChannels, _ = strconv.Atoi(value)
		case "latest_fork_usec":
			redisStatsInfo.LatestForkUsec, _ = strconv.Atoi(value)
		case "total_forks":
			redisStatsInfo.TotalForks, _ = strconv.Atoi(value)
		case "migrate_cached_sockets":
			redisStatsInfo.MigrateCachedSockets, _ = strconv.Atoi(value)
		case "slave_expires_tracked_keys":
			redisStatsInfo.SlaveExpiresTrackedKeys, _ = strconv.Atoi(value)
		case "active_defrag_hits":
			redisStatsInfo.ActiveDefragHits, _ = strconv.Atoi(value)
		case "active_defrag_misses":
			redisStatsInfo.ActiveDefragMisses, _ = strconv.Atoi(value)
		case "active_defrag_key_hits":
			redisStatsInfo.ActiveDefragKeyHits, _ = strconv.Atoi(value)
		case "active_defrag_key_misses":
			redisStatsInfo.ActiveDefragKeyMisses, _ = strconv.Atoi(value)
		case "total_active_defrag_time":
			redisStatsInfo.TotalActiveDefragTime, _ = strconv.Atoi(value)
		case "current_active_defrag_time":
			redisStatsInfo.CurrentActiveDefragTime, _ = strconv.Atoi(value)
		case "tracking_total_keys":
			redisStatsInfo.TrackingTotalKeys, _ = strconv.Atoi(value)
		case "tracking_total_items":
			redisStatsInfo.TrackingTotalItems, _ = strconv.Atoi(value)
		case "tracking_total_prefixes":
			redisStatsInfo.TrackingTotalPrefixes, _ = strconv.Atoi(value)
		case "unexpected_error_replies":
			redisStatsInfo.UnexpectedErrorReplies, _ = strconv.Atoi(value)
		case "total_error_replies":
			redisStatsInfo.TotalErrorReplies, _ = strconv.Atoi(value)
		case "dump_payload_sanitizations":
			redisStatsInfo.DumpPayloadSanitizations, _ = strconv.Atoi(value)
		case "total_reads_processed":
			redisStatsInfo.TotalReadsProcessed, _ = strconv.Atoi(value)
		case "total_writes_processed":
			redisStatsInfo.TotalWritesProcessed, _ = strconv.Atoi(value)
		case "io_threaded_reads_processed":
			redisStatsInfo.IOThreadedReadsProcessed, _ = strconv.Atoi(value)
		case "io_threaded_writes_processed":
			redisStatsInfo.IOThreadedWritesProcessed, _ = strconv.Atoi(value)
		case "reply_buffer_shrinks":
			redisStatsInfo.ReplyBufferShrinks, _ = strconv.Atoi(value)
		case "reply_buffer_expands":
			redisStatsInfo.ReplyBufferExpands, _ = strconv.Atoi(value)
		case "eventloop_cycles":
			redisStatsInfo.EventloopCycles, _ = strconv.Atoi(value)
		case "eventloop_duration_sum":
			redisStatsInfo.EventloopDurationSum, _ = strconv.Atoi(value)
		case "eventloop_duration_cmd_sum":
			redisStatsInfo.EventloopDurationCmdSum, _ = strconv.Atoi(value)
		case "instantaneous_eventloop_cycles_per_sec":
			redisStatsInfo.InstantaneousEventloopCyclesPerSec, _ = strconv.Atoi(value)
		case "instantaneous_eventloop_duration_usec":
			redisStatsInfo.InstantaneousEventloopDurationUsec, _ = strconv.Atoi(value)
		case "acl_access_denied_auth":
			redisStatsInfo.ACLAccessDeniedAuth, _ = strconv.Atoi(value)
		case "acl_access_denied_cmd":
			redisStatsInfo.ACLAccessDeniedCmd, _ = strconv.Atoi(value)
		case "acl_access_denied_key":
			redisStatsInfo.ACLAccessDeniedKey, _ = strconv.Atoi(value)
		case "acl_access_denied_channel":
			redisStatsInfo.ACLAccessDeniedChannel, _ = strconv.Atoi(value)
		default:
			// 알 수 없는 필드는 무시
		}
	}

	return redisStatsInfo, nil
}
func (rcc *redisClientCommander) InfoStat(ctx context.Context) (*RedisStatsInfo, error) {
	cmd := rcc.client.InfoMap(ctx, "stats")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return parseRedisStatsInfo(cmd.Val()["stats"])
}

// 문자열을 CPUInfo 구조체로 변환하는 함수
func parseCPUInfo(data []string) (*RedisCpuInfo, error) {
	var cpuInfo = &RedisCpuInfo{}

	for _, line := range data {
		if line == "" { continue }
		
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

func (rcc *redisClientCommander) InfoCpu(ctx context.Context) (*RedisCpuInfo, error) {
	cmd := rcc.client.Do(ctx, "info", "cpu")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	s :=  cmd.String()
	split := strings.Split(s,"\n")

	return parseCPUInfo(split[1:])
}
