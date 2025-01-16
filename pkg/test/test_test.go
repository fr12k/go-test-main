package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestSuite(t *testing.T) {
	testNewTestSuite := NewTestSuite("test_test")

	assert.Equal(t, "test_test", testNewTestSuite.AppName)
}
