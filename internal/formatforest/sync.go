// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"os/exec"
)

func Sync() {
	fmt.Printf("[FormatForest] Parsing config.txt...")
	configJson := parseConfig()
	fmt.Println(" OK")
	syncExec(configJson)
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
