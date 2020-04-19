// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"log"
)

func errorExit(err error) {
	log.Fatal(fmt.Errorf("[FF] Error: %v.\n", err))
}
