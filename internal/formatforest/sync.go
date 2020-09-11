// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"os/exec"
)

// Sync synchronizes the FormatForest public files with a remote server.
func Sync() {
	fmt.Printf("[FormatForest] Parsing config.txt...")
	configJSON := parseConfig()
	fmt.Println(" OK")
	syncExec(configJSON)
}

func syncExec(config config) {
	fmt.Printf("[FormatForest] Syncing...")
	arg := fmt.Sprintf("%s@%s:%s", config.SyncUser, config.SyncHost, config.SyncPath)
	rsync := exec.Command("rsync", "-av", "--delete", ".", arg)
	rsync.Dir = "public"
	err := rsync.Run()
	ErrorCheckExit(err)
	fmt.Println(" OK")
}
