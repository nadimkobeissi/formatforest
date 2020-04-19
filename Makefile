# SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
# SPDX-License-Identifier: GPL-3.0-only

all:
	@make -s windows
	@make -s linux
	@make -s macos
	@make -s freebsd

windows:
	@/bin/echo -n "[FF] Building Format Forest for Windows..."
	@GOOS="windows" go generate ./...
	@GOOS="windows" go build -gcflags="-e" -ldflags="-s -w" -o build/windows formatforest.com/...
	@/bin/echo " OK"

linux:
	@/bin/echo -n "[FF] Building Format Forest for Linux..."
	@GOOS="linux" go build -gcflags="-e" -ldflags="-s -w" -o build/linux formatforest.com/...
	@/bin/echo "   OK"

macos:
	@/bin/echo -n "[FF] Building Format Forest for macOS..."
	@GOOS="darwin" go build -gcflags="-e" -ldflags="-s -w" -o build/macos formatforest.com/...
	@/bin/echo "   OK"

freebsd:
	@/bin/echo -n "[FF] Building Format Forest for FreeBSD..."
	@GOOS="freebsd" go build -gcflags="-e" -ldflags="-s -w" -o build/freebsd formatforest.com/...
	@/bin/echo " OK"

lint:
	@/bin/echo "[FF] Running golangci-lint..."
	@golangci-lint run

clean:
	@/bin/echo -n "[FF] Cleaning up..."
	@$(RM) -f cmd/formatforest/resource.syso
	@$(RM) build/windows/formatforest.exe
	@$(RM) build/linux/formatforest
	@$(RM) build/macos/formatforest
	@$(RM) build/freebsd/formatforest
	@$(RM) -rf dist
	@/bin/echo "                   OK"

.PHONY: all windows linux macos freebsd lint clean build cmd docs examples internal
