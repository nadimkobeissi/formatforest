/* SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
 * SPDX-License-Identifier: GPL-3.0-only */

package formatforest

import (
	"fmt"
	"os"
	"strings"
)

// Help displays FormatForest command-line usage instructions.
func Help() {
	fmt.Fprintf(os.Stdout, strings.Join([]string{
		"format:  Generate public static folder.",
		"sync:    Sync with remote server via rsync. ",
		"help:    Show this help text.",
	}, "\n"))
}
