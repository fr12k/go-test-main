# go-test-main

This package provides a simple way to test the `main` function of a Go program.

## Example

### Without Test Coverage

This is a simple example of how to test the `main` function of a Go program.

```go mdox-exec="cat main.go"
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

```go mdox-exec="cat main_test.go"
package main

import (
	"testing"

	. "github.com/fr12k/go-test-main/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testSuite = NewTestSuite("main_test")

func TestMain(m *testing.M) {
	testSuite.TestMain(m)
}

func TestCallMain(t *testing.T) {
	cmd, buf := testSuite.Command()
	err := cmd.Run()
	require.NoError(t, err)

	assert.Equal(t, "Hello, World!\n", buf.String())
}
```

### With Test Coverage

```go mdox-exec="cat example/coverage/main.go"
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

```go mdox-exec="cat example/coverage/main_test.go"
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
```
