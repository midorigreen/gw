package main

import (
	"log"
	"os"
)

var mid = []cmdMiddleware{}

func run(args []string) error {
	var cmd Cmder = cmdImpl{}
	cmd = Chain(
		WrapTime(),
		WrapFirstEcho("=== START ==="),
		WrapSlack(os.Getenv("SLACK_TOKEN"), "send test", os.Getenv("SLACK_CHANNEL")),
		WrapEndEcho("=== END ==="),
	)(cmd)
	return cmd.Run(args, os.Stdout, os.Stderr)
}

func main() {
	args := os.Args
	if err := run(args[1:]); err != nil {
		log.Fatalf("failed wrap: %v", err)
	}
}
