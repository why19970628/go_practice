package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	rds *Redis
)

func initNewRedis() error {
	var err error
	rds, err = NewRedis(
		RedisConfig{
			Host:      "127.0.0.2:999",
			Port:      0,
			Password:  "",
			KeyPrefix: "test",
		}, nil, nil, nil)
	return err
}

func TestNewRedis(t *testing.T) {
	assert.Equal(t, nil, initNewRedis())
}
