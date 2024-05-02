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

func parseTinkerFromBytes(b []byte) (Revision, error) {
	var revision Revision
	err := json.Unmarshal(b, &revision)
	if err != nil {
		return Revision{}, err
	}

	return revision, nil
}

func ReadTinkers() ([]Revision, error) {
	f, err := os.Open(".tinker")
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer f.Close()

	files, err := f.Readdirnames(0)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	tinkers := make([]Revision, 0)
	for _, file := range files {
		b, err := os.ReadFile(".tinker/" + file)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		revision, err := parseTinkerFromBytes(b)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		tinkers = append(tinkers, revision)
	}

	return tinkers, nil
}
