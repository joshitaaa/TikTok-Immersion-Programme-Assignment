package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type RedisDatabase struct {
	cli *redis.Client
}

func (ptr *RedisClient) StartDatabase (cont context.Context, address string, password string) error {
    client := redis.NewClient(&redis.Options{
        Addr:	  address,
        Password: password, // no password set
        DB:		  0,  // use default DB
    })

	if err := client.Ping(cont); err != nil {
		return err
	}
	ptr.cli = client
	return nil 
}

func (ptr *RedisClient) StoreMessage (cont context.Context, message *Message) error {
	text, err := json.Marshal(message)
	if err != nil {
		return err
	}
}