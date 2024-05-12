package main

import (
	"hackbar-report/internal/infrastructure/cmd/report"
	"log/slog"
	"os"
)

func main() {
	err := report.Run(os.Stdout, os.Stdin)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
