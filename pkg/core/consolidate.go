package core

import (
	"encoding/json"
	"log/slog"
	"os"
)

type RevisionType string

const (
	Major RevisionType = "major"
	Minor RevisionType = "minor"
	Patch RevisionType = "patch"
)

type Revision struct {
	Type     RevisionType `json:"type"`
	Contents string       `json:"contents"`
}

func Consolidate() error {

	f, err := os.Open(".tinker")
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer f.Close()

	files, err := f.Readdirnames(0)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	for _, file := range files {
		slog.Info(file)
		b, err := os.ReadFile(".tinker/" + file)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		var revision Revision
		json.Unmarshal(b, &revision)
		slog.Info("Tinker", slog.String("type", string(revision.Type)), slog.String("contents", revision.Contents))
	}

	return nil
}
