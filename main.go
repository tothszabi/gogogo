package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

var (
	version = ""
	hash    = ""
)

func init() {
	if version != "" && hash != "" {
		return
	}

	info, available := debug.ReadBuildInfo()
	if available && info.Main.Version != "" && info.Main.Sum != "" {
		version = info.Main.Version
		hash = fmt.Sprintf("mod sum: %q", info.Main.Sum)
	} else {
		version = "local"
		hash = "unknown"
	}
}

func doNothing2() {}

func main() {
	log.Printf("Version: %s, hash: %s", version, hash)
}
