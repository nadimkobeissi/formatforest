// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/otiai10/copy"
)

func WriteInit(blogFolder string) {
	fmt.Printf("[FormatForest] Initializing folder %s...\n", blogFolder)
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
	}
	configJsonBytes, err := json.MarshalIndent(configJson, "", "\t")
	if err != nil {
		ErrorExit(err)
	}
	err = ioutil.WriteFile(path.Join(blogFolder, "config.json"), configJsonBytes, 0755)
	if err != nil {
		ErrorExit(err)
	}
}

func writeFolders() {
	publicFolderInfo, err := os.Stat("public")
	if err != nil || !publicFolderInfo.IsDir() {
		err = os.Mkdir("public", 0755)
		if err != nil {
			ErrorExit(err)
		}
	}
	postsFolderInfo, err := os.Stat(path.Join("public", "posts"))
	if err != nil || !postsFolderInfo.IsDir() {
		err = os.Mkdir(path.Join("public", "posts"), 0755)
		if err != nil {
			ErrorExit(err)
		}
	}
}

func writeHome(posts []post, config config) {
	homeHtmlBytes, err := ioutil.ReadFile(
		path.Join("templates", "home.html"),
	)
	if err != nil {
		ErrorExit(err)
	}
	homeHtml := formatStandard(string(homeHtmlBytes), config)
	homeHtml = strings.ReplaceAll(
		homeHtml, "{{FF:PostList:FF}}", formatPostList(posts),
	)
	err = ioutil.WriteFile(
		path.Join("public", "index.html"),
		[]byte(homeHtml), 0755,
	)
	if err != nil {
		ErrorExit(err)
	}
}

func writePosts(posts []post, config config) {
	postHtmlBytes, err := ioutil.ReadFile(
		path.Join("templates", "post.html"),
	)
	if err != nil {
		ErrorExit(err)
	}
	for _, post := range posts {
		postHtml := formatPost(string(postHtmlBytes), post, config)
		err = ioutil.WriteFile(
			path.Join("public", "posts",
				fmt.Sprintf("%s-%s.html",
					post.date, post.tag,
				)), []byte(postHtml), 0755)
		if err != nil {
			ErrorExit(err)
		}
	}
}

func writeRss(posts []post, config config) {
	postsRssXmlBytes, err := ioutil.ReadFile(
		"templates/rss.xml",
	)
	if err != nil {
		ErrorExit(err)
	}
	postsRssXml := strings.ReplaceAll(
		string(postsRssXmlBytes),
		"{{FF:PostRss:FF}}", formatRss(posts, config),
	)
	err = ioutil.WriteFile(
		path.Join("public", "rss.xml"),
		[]byte(postsRssXml), 0755,
	)
	if err != nil {
		ErrorExit(err)
	}
}

func writeRes() {
	err := copy.Copy("res", path.Join("public", "res"))
	if err != nil {
		ErrorExit(err)
	}
}
