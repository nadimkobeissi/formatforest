#!/usr/bin/env bash
# SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
# SPDX-License-Identifier: GPL-3.0-only
set -euo pipefail

echo -n "[FormatForest] Enter version: "
read VERSION

if [[ "$OSTYPE" == "darwin"* ]]; then
	gsed -i -e "s/version = \"\([0-9]\|.\)\+\"/version = \"${VERSION}\"/g" cmd/formatforest/main.go
else
	sed -i -e "s/version = \"\([0-9]\|.\)\+\"/version = \"${VERSION}\"/g" cmd/formatforest/main.go
fi

git commit -am "FormatForest ${VERSION}" &> /dev/null
git push &> /dev/null
git tag -a "v${VERSION}" -m "FormatForest ${VERSION}" -m "${RELEASE_NOTES}" &> /dev/null
git push origin "v${VERSION}" &> /dev/null

echo "[FormatForest] FormatForest ${VERSION} tagged."
