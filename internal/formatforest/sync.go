// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"os/exec"
)

func Sync(config config) {
	fmt.Println("[FormatForest] In order for sync to succeed,")
	fmt.Println("               rsync must be installed.")
	fmt.Println("")
	syncExec(config config)
}

func syncExec(config config) {
	fmt.Printf("[FormatForest] Syncing...")
	arg := fmt.Sprintf("%s@%s:%s", config.SyncUser, config.SyncHost, config.SyncPath)
	rsync := exec.Command("rsync", "-av", "--delete", ".", arg)
	rsync.Dir = "public"
	err := rsync.Run()
	if err != nil {
		ErrorExit(err)
	}
}
