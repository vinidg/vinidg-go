package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase() (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

func (db *Database) SetValue(key, value string) string {
	rdb := db.Client

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Error("Get failed ->", err)
	}

	log.Info(key, val)

	var list []string

	if val != "" {
		list = strings.Split(val, ",")
	}

	list = append(list, value)
	listString := strings.Join(list, ",")

	err = rdb.Set(ctx, key, listString, 0).Err()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return listString
}

func (db *Database) GetValue(key string) string {
	rdb := db.Client

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		val = fmt.Sprintf("%s does not exist", key)
		log.Error(val)
	}
	if err != nil {
		log.Error("Get failed ->", err)
	}

	return val
}
