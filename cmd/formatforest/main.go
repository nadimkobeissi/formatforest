// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"fmt"
	"os"

	"formatforest.com/internal/formatforest"
	"github.com/logrusorgru/aurora"
)

var version = "0.0.0"

func main() {
	switch len(os.Args) {
	case 3:
		mainEntry(os.Args)
	case 2:
		mainIntro()
		formatforest.Help()
	default:
		mainIntro()
		formatforest.Help()
	}
}

func mainIntro() {
	fmt.Fprintf(os.Stdout,
		aurora.Bold("Format Forest %s - %s\n").String(),
		version, "https://formatforest.com",
	)
}

func mainEntry(args []string) {
	switch args[1] {
	case "format":
		mainIntro()
		formatforest.Format()
	case "publish":
		mainIntro()
	default:
		mainIntro()
		formatforest.Help()
	}
}
