// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"os/exec"
)

func Sync(user string, host string, targetDir string) {
	fmt.Println("[FormatForest] In order for sync to succeed,")
	fmt.Println("               rsync must be installed.")
	fmt.Println("")
	syncExec(user, host, targetDir)
}

func syncExec(user string, host string, targetDir string) {
	fmt.Println("[FormatForest] Syncing...")
	arg := fmt.Sprintf("%s@%s:%s", user, host, targetDir)
	rsync := exec.Command("rsync", "-av", "--delete", ".", arg)
	rsync.Dir = "public"
	err := rsync.Run()
	if err != nil {
		ErrorExit(err)
	}
}
