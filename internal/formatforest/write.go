// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/otiai10/copy"
)

func writeFolders() {
	publicFolderInfo, err := os.Stat("public")
	if err != nil || !publicFolderInfo.IsDir() {
		err = os.Mkdir("public", 0755)
		ErrorCheckExit(err)
	}
	postsFolderInfo, err := os.Stat(path.Join("public", "posts"))
	if err != nil || !postsFolderInfo.IsDir() {
		err = os.Mkdir(path.Join("public", "posts"), 0755)
		ErrorCheckExit(err)
	}
}

func writeHome(posts []post, config config) {
	homeHtmlBytes, err := ioutil.ReadFile(
		path.Join("templates", "home.html"),
	)
	ErrorCheckExit(err)
	homeHtml := formatStandard(string(homeHtmlBytes), config)
	homeHtml = strings.ReplaceAll(
		homeHtml, "{{FF:PostList:FF}}", formatPostList(posts),
	)
	err = ioutil.WriteFile(
		path.Join("public", "index.html"),
		[]byte(homeHtml), 0755,
	)
	ErrorCheckExit(err)
}

func writePosts(posts []post, config config) {
	postHtmlBytes, err := ioutil.ReadFile(
		path.Join("templates", "post.html"),
	)
	ErrorCheckExit(err)
	for _, post := range posts {
		postHtml := formatPost(string(postHtmlBytes), post, config)
		err = ioutil.WriteFile(
			path.Join("public", "posts",
				fmt.Sprintf("%s-%s.html",
					post.date, post.tag,
				)), []byte(postHtml), 0755)
		ErrorCheckExit(err)
	}
}

func writeRss(posts []post, config config) {
	postsRssXmlBytes, err := ioutil.ReadFile(
		"templates/rss.xml",
	)
	ErrorCheckExit(err)
	postsRssXml := strings.ReplaceAll(
		string(postsRssXmlBytes),
		"{{FF:PostRss:FF}}", formatRss(posts, config),
	)
	err = ioutil.WriteFile(
		path.Join("public", "rss.xml"),
		[]byte(postsRssXml), 0755,
	)
	ErrorCheckExit(err)
}

func writeRes() {
	err := copy.Copy("res", path.Join("public", "res"))
	ErrorCheckExit(err)
}
