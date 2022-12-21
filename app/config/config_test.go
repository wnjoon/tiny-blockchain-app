package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	// Load Config with viper
	// conf, err := New(".", "config", "yaml")
	conf, err := New(".", "config", "yaml")
	if err != nil {
		t.Log(err)
		return
	}

	blockchainConfig := conf.BlockChain()
	assert.Equal(t, "http://127.0.0.1:22001", blockchainConfig.EndPoint)
	assert.Equal(t, "ws://127.0.0.1:32001", blockchainConfig.WebSocket)
}
