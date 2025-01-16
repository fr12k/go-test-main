package test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type (
	Option func(*TestSuite)

	TestSuite struct {
		AppName string

		TestCoverageFile string
		TestCoverage     bool
	}
)

func WithTestCoverage() Option {
	return func(tm *TestSuite) {
		tm.TestCoverage = true
	}
}

func WithTestCoverageFile(file string) Option {
	return func(tm *TestSuite) {
		tm.TestCoverageFile = file
	}
}

func NewTestSuite(appName string, opts ...Option) *TestSuite {
	tm := &TestSuite{
		AppName:          appName,
		TestCoverageFile: "coverage.txt",
		TestCoverage:     false,
	}
	for _, opt := range opts {
		opt(tm)
	}
	return tm
}

// TestMain is the entry point for the test coverage run.
// It builds the binary, runs the tests, and then merges the coverage data from
// the unit tests and the main integration test.
func (tm *TestSuite) TestMain(m *testing.M) {
	fmt.Println("-> Building...")

	build := exec.Command("go", "build", "-o", tm.AppName)
	if tm.TestCoverage {
		build = exec.Command("go", "build", "-cover", "-o", tm.AppName)
	}

	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building %s: %s", tm.AppName, err)
		os.Exit(1)
	}
	fmt.Println("-> Running...")
	if tm.TestCoverage {
		err := os.MkdirAll(".coverdata", 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s", err)
			os.Exit(1)
		}
	}
	result := m.Run()
	fmt.Println("-> Getting done...")

	if tm.TestCoverage {
		cmd := exec.Command("go", "tool", "covdata", "textfmt", "-i=.coverdata/", "-o", "system.out")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running command: %s", err)
			os.Exit(1)
		}

		var buf bytes.Buffer
		cmd = exec.Command("gocovmerge", "system.out", tm.TestCoverageFile)
		cmd.Stdout = &buf
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running command: %s", err)
			os.Exit(1)
		}

		err = os.WriteFile(tm.TestCoverageFile, buf.Bytes(), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s", err)
			os.Exit(1)
		}
	}
	os.Remove(tm.AppName)
	if tm.TestCoverage {
		os.Remove("system.out")
		os.RemoveAll(".coverdata")
	}
	os.Exit(result)
}

func (tm *TestSuite) Command() (*exec.Cmd, *bytes.Buffer) {
	var buf bytes.Buffer
	cmd := exec.Command("./" + tm.AppName)
	if tm.TestCoverage {
		cmd.Env = append(os.Environ(), "GOCOVERDIR=.coverdata")
	}
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd, &buf
}
