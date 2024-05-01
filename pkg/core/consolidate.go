package core

import (
	"log/slog"
	"os"
)

type File struct {
	name string
}

func Consolidate() error {

	f, err := os.Open(".tinker")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	files, err := f.Readdirnames(0)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	for _, file := range files {
		slog.Info(file)
	}

	return nil
}
