package config_test


import (
	"testing"
	"agent_redis/pkg/config"
)

const agentRedisConfigPath = "../../assets/config/agent_redis.toml"

func TestReadConfigFromFile(t *testing.T) {
	cfg, err := config.ReadConfigFromFile(agentRedisConfigPath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Config read successfully")
	t.Logf("%+v", cfg)
}