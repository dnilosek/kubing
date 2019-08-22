package test

import (
	"github.com/alicebob/miniredis"
	"github.com/dnilosek/kubing/app/lib/database"
	"github.com/go-redis/redis"
)

func MockDB() *database.DB {
	mr, _ := miniredis.Run()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	db := database.DB{
		Client: client,
	}
	return &db
}
