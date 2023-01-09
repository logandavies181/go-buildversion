package buildversion

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// Some notes on debug.ReadBuildInfo copied from: https://github.com/logandavies181/tfd/blob/main/main.go
// TODO: update
//
// buildVersion checks if version has been set at build time, otherwise uses debug.ReadBuildInfo to infer a tag.
// debug.ReadBuildInfo behaves differently given the following scenarios so buildVersion does the following:
// Locally compiled - Main.Version = "devel"; return "0.0.0+devel"
// go get tfd@<some git hash> - Main.Version = "0.0.0+v0.0.0-<timestamp>-<short_hash>"; return 0.0.0-<timestamp>-<short_hash>
// go get tfd@<tag> - Main.Version = "0.0.0-v<tag>"; return "<tag>"

// TODO
// BuildVersionLong
// BuildVersion
// BuildVersionShort

func BuildVersionShortE(version string) (string, error) {
	if version == "" {
		debugInfo, ok := debug.ReadBuildInfo()
		if !ok {
			return "", fmt.Errorf("Could not determine build info. This binary might be stripped")
		}

		version = debugInfo.Main.Version
		switch version {
		case "(devel)":
			return "0.0.0+devel", nil
		case "":
			fmt.Println(debugInfo.Settings)
			return "0.0.0+go-run", nil
		}

		// get rid of prefix included during `go install`
		version = strings.TrimPrefix(version, "0.0.0-v")
	}
	return version, nil
}
