package test_test

import (
	"testing"

	. "github.com/dnilosek/kubing/app/lib/test"
	"github.com/stretchr/testify/assert"
)

func TestMockDB(t *testing.T) {
	db := MockDB()
	assert.NotNil(t, db)
}
