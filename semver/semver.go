// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

// Package semver implements a Version implementation.
package semver

import "fmt"

// Version holds the semantic version.
type Version struct {
	major      int
	minor      int
	patch      int
	preRelease string
	build      string
}

// New returns a new Version with no pre-release or build
func New(major, minor, patch int) Version {
	return Version{major: major, minor: minor, patch: patch}
}

// NewBuild returns a new Version with build tag
func NewBuild(major, minor, patch int, build string) Version {
	return Version{major: major, minor: minor, patch: patch, build: build}
}

// NewPreRelease returns a new Version with pre-release tag
func NewPreRelease(major, minor, patch int, preRelease string) Version {
	return Version{major: major, minor: minor, patch: patch, preRelease: preRelease}
}

// NewPreReleaseAndBuild returns a new Version with pre-release and build tags
func NewPreReleaseAndBuild(major, minor, patch int, preRelease, build string) Version {
	return Version{major: major, minor: minor, patch: patch, preRelease: preRelease, build: build}
}

// String implements the Stringer interface
func (s Version) String() string {
	if s.preRelease != "" && s.build != "" {
		return fmt.Sprintf("%d.%d.%d-%s+%s", s.major, s.minor, s.patch, s.preRelease, s.build)
	} else if s.build != "" {
		return fmt.Sprintf("%d.%d.%d+%s", s.major, s.minor, s.patch, s.build)
	} else if s.preRelease != "" {
		return fmt.Sprintf("%d.%d.%d-%s", s.major, s.minor, s.patch, s.preRelease)
	}
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}
