//go:build testrunmain

package main

import (
	"testing"

	. "github.com/fr12k/go-test-main/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testSuite = NewTestSuite("main_test", WithTestCoverage(), WithTestCoverageFile("../../coverage.txt"))

func TestMain(m *testing.M) {
	testSuite.TestMain(m)
}

func TestCallMainWithCoverage(t *testing.T) {
	cmd, buf := testSuite.Command()
	err := cmd.Run()
	require.NoError(t, err)

	assert.Equal(t, "Hello, World!\n", buf.String())
}
