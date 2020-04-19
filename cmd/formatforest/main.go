// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"errors"
	"fmt"
	"os"

	"formatforest.com/internal/formatforest"
)

var version = "0.0.0"

func main() {
	if len(os.Args) >= 2 {
		mainEntry(os.Args)
	} else {
		mainIntro()
		formatforest.Help()
	}
}

func mainIntro() {
	fmt.Fprintf(os.Stdout, "FormatForest %s - %s\n\n",
		version, "https://formatforest.com",
	)
}

func mainEntry(args []string) {
	switch args[1] {
	case "format":
		mainIntro()
		if len(args) != 2 {
			formatforest.ErrorExit(errors.New("format requires zero arguments"))
		}
		formatforest.Format()
	case "sync":
		mainIntro()
		if len(args) != 5 {
			formatforest.ErrorExit(errors.New("sync requires three arguments"))
		}
		formatforest.Sync(args[2], args[3], args[4])
	default:
		mainIntro()
		formatforest.Help()
	}
}
