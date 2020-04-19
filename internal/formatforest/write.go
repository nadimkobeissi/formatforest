// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func writeFolders() {
	publicFolderInfo, err := os.Stat("public")
	if err != nil || !publicFolderInfo.IsDir() {
		err = os.Mkdir("public", 0755)
		if err != nil {
			errorExit(err)
		}
	}
	postsFolderInfo, err := os.Stat("posts")
	if err != nil || !postsFolderInfo.IsDir() {
		err = os.Mkdir(path.Join("public", "posts"), 0755)
		if err != nil {
			errorExit(err)
		}
	}
}

func writeHome(posts []post) {
	homeHtmlBytes, err := ioutil.ReadFile(
		"templates/home.html",
	)
	if err != nil {
		errorExit(err)
	}
	homeHtml := strings.ReplaceAll(
		string(homeHtmlBytes), "{{FF:PostList:FF}}", formatPostList(posts),
	)
	err = ioutil.WriteFile(
		"public/index.html",
		[]byte(homeHtml), 0755,
	)
	if err != nil {
		errorExit(err)
	}
}

func writePosts(posts []post) {
	postHtmlBytes, err := ioutil.ReadFile(
		"templates/post.html",
	)
	if err != nil {
		errorExit(err)
	}
	for _, post := range posts {
		postHtml := formatPost(string(postHtmlBytes), post)
		err = ioutil.WriteFile(fmt.Sprintf(
			"public/posts/%s-%s.html",
			post.date, post.tag,
		), []byte(postHtml), 0755)
		if err != nil {
			errorExit(err)
		}
	}
}

func writeRss(posts []post) {
	postsRssXmlBytes, err := ioutil.ReadFile(
		"templates/rss.xml",
	)
	if err != nil {
		errorExit(err)
	}
	postsRssXml := strings.ReplaceAll(
		string(postsRssXmlBytes), "{{FF:PostRss:FF}}", formatRss(posts),
	)
	err = ioutil.WriteFile(
		"public/rss.xml",
		[]byte(postsRssXml), 0755,
	)
	if err != nil {
		errorExit(err)
	}
}
