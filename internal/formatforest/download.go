// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func downloadInitAssets(blogFolder string) {
	dlPath := "https://gitlab.com/nadimk/formatforest/-/raw/master/examples/"
	err := download(
		filepath.Join(blogFolder, "res", "img", "formatforest.png"),
		fmt.Sprintf("%s%s", dlPath, "res/img/formatforest.png"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "res", "img", "hello.png"),
		fmt.Sprintf("%s%s", dlPath, "res/img/hello.png"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "res", "img", "rss.svg"),
		fmt.Sprintf("%s%s", dlPath, "res/img/rss.svg"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "res", "css", "style.css"),
		fmt.Sprintf("%s%s", dlPath, "res/css/style.css"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "res", "css", "highlight.css"),
		fmt.Sprintf("%s%s", dlPath, "res/css/highlight.css"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "res", "js", "highlight.js"),
		fmt.Sprintf("%s%s", dlPath, "res/js/highlight.js"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "templates", "home.html"),
		fmt.Sprintf("%s%s", dlPath, "templates/home.html"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "templates", "post.html"),
		fmt.Sprintf("%s%s", dlPath, "templates/post.html"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "templates", "rss.xml"),
		fmt.Sprintf("%s%s", dlPath, "templates/rss.xml"),
	)
	ErrorCheckExit(err)
	err = download(
		filepath.Join(blogFolder, "posts", "2020-04-19-hello.md"),
		fmt.Sprintf("%s%s", dlPath, "posts/2020-04-19-hello.md"),
	)
	ErrorCheckExit(err)
}
