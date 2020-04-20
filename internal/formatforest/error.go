// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"log"
)

func ErrorExit(err error) {
	log.Fatal(fmt.Errorf("[FormatForest] Error: %v.\n", err))
}

func ErrorCheckExit(err error) {
	if err != nil {
		ErrorExit(err)
	}
}
