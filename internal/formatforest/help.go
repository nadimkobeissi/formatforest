/* SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
 * SPDX-License-Identifier: GPL-3.0-only */

package formatforest

import (
	"fmt"
	"os"
	"strings"
)

// Help displays Format Forest command-line usage instructions.
func Help() {
	fmt.Fprintf(os.Stdout, strings.Join([]string{
		"format:  TBD.",
		"publish: TBD.",
		"help:    Show this help text.",
	}, "\n"))
}
