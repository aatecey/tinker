package version

import (
	"bytes"
	"errors"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Version struct {
	Major uint64
	Minor uint64
	Patch uint64
}

func NewVersion(major uint64, minor uint64, patch uint64) Version {
	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func parseVersionFromString(ver string) (Version, error) {
	versions := strings.Split(ver, ".")
	if len(versions) != 3 {
		return Version{}, errors.New("Invalid version format")
	}

	major, err := strconv.ParseUint(versions[0], 10, 64)
	if err != nil {
		return Version{}, err
	}
	minor, err := strconv.ParseUint(versions[1], 10, 64)
	if err != nil {
		return Version{}, err
	}
	patch, err := strconv.ParseUint(versions[2], 10, 64)
	if err != nil {
		return Version{}, err
	}

	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func ReadVersion() (Version, error) {
	b, err := os.ReadFile("version.txt")
	if err != nil {
		slog.Error(err.Error())
		return Version{}, err
	}

	return parseVersionFromString(string(bytes.TrimSpace(b)))
}
