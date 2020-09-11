// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"log"
)

// ErrorExit exits on error.
func ErrorExit(err error) {
	log.Fatal(fmt.Errorf("%v", err))
}

// ErrorCheckExit exits if err != nil
func ErrorCheckExit(err error) {
	if err != nil {
		ErrorExit(err)
	}
}
