package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	latest "github.com/UedaTakeyuki/getLatestReleaseName/latestReleaseName"
)

func main() {
	user := os.Args[1]
	repository := os.Args[2]
	if name, err := latest.GetLatestReleaseName(user, repository); err != nil {
		if errors.Is(err, latest.ERR_NORELEASE) {
			slog.Info("No release")
		} else {
			slog.Error(err.Error())
		}
		fmt.Println("")
	} else {
		slog.Info("no error")
		fmt.Println(name)
	}
}
