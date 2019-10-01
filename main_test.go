package main

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestMain(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	t.Log(client.SetNX("key", "value", 10*time.Second).Result())
	t.Log(client.Get("key").Result())
}
