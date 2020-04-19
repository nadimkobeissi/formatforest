# SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
# SPDX-License-Identifier: GPL-3.0-only

lint:
	@golangci-lint run

.PHONY: lint cmd docs internal