package database_test

import (
	"testing"

	"github.com/dnilosek/kubing/apps/visitor-count/lib/database"
	"github.com/dnilosek/kubing/apps/visitor-count/lib/test"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	redisUrl := "redis://" + mr.Addr()
	db, err := database.Open(redisUrl)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	// Bad path
	db, err = database.Open("")
	assert.Nil(t, db)
	assert.NotNil(t, err)

	// Bad parameter
	redisUrl += "?boogyman=exists"
	db, err = database.Open(redisUrl)
	assert.Nil(t, db)
	assert.NotNil(t, err)
}

func TestClose(t *testing.T) {
	db := database.DB{
		Client: test.MockRedis(),
	}
	err := db.Close()
	assert.Nil(t, err)
}

func TestSetGet(t *testing.T) {
	db := database.DB{
		Client: test.MockRedis(),
	}

	err := db.Set("foo", "bar")
	assert.Nil(t, err)

	val, err := db.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, val, "bar")
}
