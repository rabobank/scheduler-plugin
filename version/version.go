package version

import (
	"fmt"
	"strconv"
	"strings"
)

var VERSION = "1.0.9"
var COMMIT = "dev"

type VersionCommand struct {
}

// Execute - returns the version
func (c *VersionCommand) Execute([]string) error {
	fmt.Println(GetFormattedVersion())
	return nil
}

func GetVersion() string {
	return strings.Replace(VERSION, "v", "", 1)
}

func GetFormattedVersion() string {
	return fmt.Sprintf("Version: [%s], Commit: [%s]", GetVersion(), COMMIT)
}

func GetMajorVersion() int {
	if len(GetVersion()) == 0 {
		return 1
	}
	version, _ := strconv.Atoi(strings.Split(GetVersion(), ".")[0])
	return version
}

func GetMinorVersion() int {
	if len(GetVersion()) == 0 {
		return 0
	}
	version, _ := strconv.Atoi(strings.Split(GetVersion(), ".")[1])
	return version
}

func GetPatchVersion() int {
	if len(GetVersion()) == 0 {
		return 0
	}
	version, _ := strconv.Atoi(strings.Split(GetVersion(), ".")[2])
	return version
}
