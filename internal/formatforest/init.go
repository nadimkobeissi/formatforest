// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func Init(blogFolder string) {
	fmt.Printf("[FormatForest] Initializing folder %s...", blogFolder)
	_, err := os.Stat(blogFolder)
	if err == nil {
		ErrorExit(errors.New("folder already exists"))
	}
	err = os.Mkdir(blogFolder, 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = os.Mkdir(path.Join(blogFolder, "posts"), 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = os.Mkdir(path.Join(blogFolder, "public"), 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = os.Mkdir(path.Join(blogFolder, "templates"), 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = os.Mkdir(path.Join(blogFolder, "res"), 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = os.Mkdir(path.Join(blogFolder, "res", "img"), 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = os.Mkdir(path.Join(blogFolder, "res", "css"), 0755)
	if err != nil {
		ErrorExit(err)
	}
	configJson := config{
		WebsiteName:        "My FormatForest Blog",
		WebsiteUri:         "https://myblog.com",
		WebsiteDescription: "Welcome to my personal blog.",
		WebsiteIcon:        "formatforest.png",
		WebsiteTwitter:     "forestformat",
		WebsiteLang:        "en",
		WebsiteLangDir:     "ltr",
		AuthorName:         "Format Gardener",
		AuthorEmail:        "your@email.com",
		AuthorTwitter:      "yourtwitter",
		AuthorLinkedIn:     "yourlinkedinid",
		AuthorFacebook:     "yourfacebookid",
		AuthorInstagram:    "yourinstagramid",
		SyncUser:           "user",
		SyncHost:           "myblog.com",
		SyncPath:           "/var/www/myblog.com",
	}
	configJsonBytes, err := json.MarshalIndent(configJson, "", "\t")
	if err != nil {
		ErrorExit(err)
	}
	err = ioutil.WriteFile(path.Join(blogFolder, "config.txt"), configJsonBytes, 0755)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "res", "img", "formatforest.png"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/res/img/formatforest.png",
	)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "res", "img", "hello.png"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/res/img/hello.png",
	)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "res", "css", "style.css"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/res/css/style.css",
	)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "templates", "home.html"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/templates/home.html",
	)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "templates", "post.html"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/templates/post.html",
	)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "templates", "rss.xml"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/templates/rss.xml",
	)
	if err != nil {
		ErrorExit(err)
	}
	err = initDownload(
		path.Join(blogFolder, "posts", "2020-04-19-hello.md"),
		"https://gitlab.com/nadimk/formatforest/-/raw/master/examples/posts/2020-04-19-hello.md",
	)
	if err != nil {
		ErrorExit(err)
	}
}

func initDownload(filepath string, url string) error {
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
