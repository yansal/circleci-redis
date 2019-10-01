package main

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestMain(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	t.Log(client.Set("key", "value", 0).Result())
	t.Log(client.Get("key").Result())
}
