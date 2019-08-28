package work_test

import (
	"testing"
	"time"

	"github.com/dnilosek/kubing/apps/fib-overkill/lib/database"
	"github.com/dnilosek/kubing/apps/fib-overkill/lib/test"
	. "github.com/dnilosek/kubing/apps/fib-overkill/worker/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestListen(t *testing.T) {

	db := database.DB{
		Client: test.MockRedis(),
	}

	listener := NewListener(&db)

	channel := "testC"
	message := "testM"
	msgChan := listener.Listen(channel)

	go func() {
		// Delay for 1s and send a message
		time.Sleep(1 * time.Second)
		db.Client.Publish(channel, message)
	}()

	v, ok := <-msgChan
	assert.True(t, ok)
	assert.Equal(t, v, message)

	// Kill connection and force error
	db.Client.Close()
	v, ok = <-msgChan
	assert.False(t, ok)
	assert.Zero(t, v)
}
