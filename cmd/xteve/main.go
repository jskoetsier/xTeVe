package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var version = "3.0.0"

func main() {
	var (
		port      = flag.Int("port", 34400, "HTTP port")
		configDir = flag.String("config", defaultConfigDir(), "Config directory")
		debug     = flag.Int("debug", 0, "Debug level 0-3")
	)
	flag.Parse()

	fmt.Printf("xTeVe %s\n", version)
	_ = port
	_ = configDir
	_ = debug

	log.Println("TODO: wire up server")
	os.Exit(0)
}

func defaultConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".xteve"
	}
	return home + "/.xteve"
}
