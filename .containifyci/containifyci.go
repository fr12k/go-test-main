//go:generate bash -c "if [ ! -f go.mod ]; then echo 'Initializing go.mod...'; go mod init .containifyci; else echo 'go.mod already exists. Skipping initialization.'; fi"
//go:generate go get github.com/containifyci/engine-ci/protos2
//go:generate go get github.com/containifyci/engine-ci/client
//go:generate go mod tidy

package main

import (
	"os"

	"github.com/containifyci/engine-ci/client/pkg/build"
	"github.com/containifyci/engine-ci/protos2"
)

func main() {
	os.Chdir("../")
	opts := build.NewGoServiceBuild("go-test-main")
	opts.Verbose = false
	opts.Image = ""
	opts.File = "main.go"

	opts2 := build.NewGoServiceBuild("go-test-main-main")
	opts2.Verbose = false
	opts2.Image = ""
	opts2.File = "main.go"
	opts2.Properties = map[string]*build.ListValue{
		"tags":            build.NewList("testrunmain"),
		"nocoverage":      build.NewList("true"),
	}
	opts3 := build.NewGoServiceBuild("go-test-main-main")
	opts3.Verbose = false
	opts3.Image = ""
	opts3.File = "example/coverage/main.go"
	opts3.Properties = map[string]*build.ListValue{
		"tags":            build.NewList("testrunmain"),
		"nocoverage":      build.NewList("true"),
	}
	// build.Serve(opts, opts2, opts3)

	build.BuildGroups(
		&protos2.BuildArgsGroup{
			Args: []*protos2.BuildArgs{opts},
		},
		&protos2.BuildArgsGroup{
			Args: []*protos2.BuildArgs{opts2, opts3},
		},
	)
}
