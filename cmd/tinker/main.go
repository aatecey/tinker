package main

import (
	"log/slog"

	cn "github.com/aatecey/tinker/pkg/core"
)

func main() {
	slog.Info("Starting the CLI")

	err := cn.Consolidate()
	if err != nil {
		slog.Error(err.Error())
	}
}
