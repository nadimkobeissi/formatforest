# SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
# SPDX-License-Identifier: GPL-3.0-only

all:
	@make -s windows
	@make -s linux
	@make -s macos
	@make -s freebsd

windows:
	@/bin/echo -n "[FormatForest] Building FormatForest for Windows..."
	@make -s dep
	@GOOS="windows" go generate ./...
	@GOOS="windows" go build -gcflags="-e" -ldflags="-s -w" -o build/windows formatforest.com/...
	@/bin/echo " OK"

linux:
	@/bin/echo -n "[FormatForest] Building FormatForest for Linux..."
	@make -s dep
	@GOOS="linux" go build -gcflags="-e" -ldflags="-s -w" -o build/linux formatforest.com/...
	@/bin/echo "   OK"

macos:
	@/bin/echo -n "[FormatForest] Building FormatForest for macOS..."
	@make -s dep
	@GOOS="darwin" go build -gcflags="-e" -ldflags="-s -w" -o build/macos formatforest.com/...
	@/bin/echo "   OK"

freebsd:
	@/bin/echo -n "[FormatForest] Building FormatForest for FreeBSD..."
	@make -s dep
	@GOOS="freebsd" go build -gcflags="-e" -ldflags="-s -w" -o build/freebsd formatforest.com/...
	@/bin/echo " OK"

dep:
	go get -u ./...

lint:
	@/bin/echo "[FormatForest] Running golangci-lint..."
	@golangci-lint run

clean:
	@/bin/echo -n "[FormatForest] Cleaning up..."
	@$(RM) -f cmd/formatforest/resource.syso
	@$(RM) build/windows/formatforest.exe
	@$(RM) build/linux/formatforest
	@$(RM) build/macos/formatforest
	@$(RM) build/freebsd/formatforest
	@$(RM) -rf dist
	@$(RM) -f examples/formatforest
	@$(RM) -f formatforest
	@/bin/echo "                   OK"

.PHONY: dep all windows linux macos freebsd lint clean build cmd docs examples internal
