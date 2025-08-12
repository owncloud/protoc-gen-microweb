//go:build tools
// +build tools

package main

// This file pins CLI tool dependencies in go.mod for reproducible installs.
// Tools are imported as blanks and guarded by the "tools" build tag so they
// are not included in normal builds.

import (
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
