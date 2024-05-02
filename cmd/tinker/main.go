package main

import (
	"log/slog"

	cn "github.com/aatecey/tinker/pkg/core"
	v "github.com/aatecey/tinker/pkg/core/version"
)

func main() {
	version, err := v.ReadVersion()
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("current version", slog.Uint64("major", version.Major), slog.Uint64("minor", version.Minor), slog.Uint64("patch", version.Patch))

	tinkers, err := cn.ReadTinkers()
	if err != nil {
		slog.Error(err.Error())
	}
	for _, tinker := range tinkers {
		slog.Info("tinker", slog.String("type", string(tinker.Type)), slog.String("contents", tinker.Contents))
	}
}
