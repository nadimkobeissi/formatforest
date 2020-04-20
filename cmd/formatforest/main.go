// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"errors"
	"fmt"
	"os"

	"formatforest.com/internal/formatforest"
)

var version = "0.0.1"

func main() {
	mainIntro()
	if len(os.Args) >= 2 {
		mainEntry(os.Args)
	} else {
		formatforest.Help()
	}
}

func mainIntro() {
	fmt.Fprintf(os.Stdout, "FormatForest %s - %s\n",
		version, "https://formatforest.com",
	)
}

func mainEntry(args []string) {
	switch args[1] {
	case "init":
		if len(args) != 3 {
			formatforest.ErrorExit(
				errors.New("init requires one argument"),
			)
		}
		formatforest.Init(args[2])
	case "format":
		if len(args) != 2 {
			formatforest.ErrorExit(
				errors.New("format requires zero arguments"),
			)
		}
		formatforest.Format()
	case "sync":
		if len(args) != 2 {
			formatforest.ErrorExit(
				errors.New("sync requires zero arguments"),
			)
		}
		formatforest.Sync()
	default:
		formatforest.Help()
	}
}
