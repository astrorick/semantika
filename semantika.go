// Semantika is a Go package that provides a simple but strict implementation of semantic versioning in Go, ased on the [Semantic Versioning 2.0.0](https://semver.org/) specification.
package semantika

import (
	"fmt"
	"strconv"
	"strings"
)

// Version stores version numbers, pre-release metadata and build metadata.
type Version struct {
	Major uint64
	Minor uint64
	Patch uint64
}

// New takes a version string in the format "Major.Minor.Patch" and populates the Version object with the parsed values.
// It returns an error if the input string is not in the correct format or if the numeric values cannot be parsed.
func New(versionString string) (*Version, error) {
	// split input string
	parts := strings.Split(versionString, ".")

	// check for consistency
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid version string: %s", versionString)
	}

	// parse version parts
	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", parts[0])
	}
	minor, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", parts[1])
	}
	patch, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid patch version: %s", parts[2])
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

// String return a string representation of the Version object in the format "Major.Minor.Patch".
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Compare matches the reference Version object against the argument Version object.
// It returns -1 if the reference version is older, 0 if they are equal, and 1 if the reference version is newer.
func (v1 *Version) Compare(v2 *Version) int8 {
	// compare major versions
	if v1.Major != v2.Major {
		if v1.Major < v2.Major {
			return -1
		}
		return 1
	}

	// compare minor versions
	if v1.Minor != v2.Minor {
		if v1.Minor < v2.Minor {
			return -1
		}
		return 1
	}

	// compare patch versions
	if v1.Patch != v2.Patch {
		if v1.Patch < v2.Patch {
			return -1
		}
		return 1
	}

	// versions match
	return 0
}

// OlderThan returns true only if the reference version is older than the argument version, and false otherwise.
func (v1 *Version) OlderThan(v2 *Version) bool {
	return v1.Compare(v2) == -1
}

// OlderThanOrEquals returns true if the reference version is older than or equal to the argument version, and false otherwise.
func (v1 *Version) OlderThanOrEquals(v2 *Version) bool {
	res := v1.Compare(v2)
	return res == -1 || res == 0
}

// Equals returns true only if the reference version is equal to the argument version, and false otherwise.
func (v1 *Version) Equals(v2 *Version) bool {
	return v1.Compare(v2) == 0
}

// NewerThanOrEquals returns true if the reference version is newer than or equal to the argument version, and false otherwise.
func (v1 *Version) NewerThanOrEquals(v2 *Version) bool {
	res := v1.Compare(v2)
	return res == 0 || res == 1
}

// NewerThan returns true only if the reference version is newer than the argument version, and false otherwise.
func (v1 *Version) NewerThan(v2 *Version) bool {
	return v1.Compare(v2) == 1
}
