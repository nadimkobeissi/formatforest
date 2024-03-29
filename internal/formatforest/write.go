// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

func writeInitFolders(blogFolder string) {
	err := os.Mkdir(blogFolder, 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "posts"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "public"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "templates"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "res"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "res", "img"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "res", "css"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "res", "js"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "res", "pdf"), 0755)
	ErrorCheckExit(err)
	err = os.Mkdir(filepath.Join(blogFolder, "res", "zip"), 0755)
	ErrorCheckExit(err)
}

func writePublicFolders() {
	publicFolderInfo, err := os.Stat("public")
	if err != nil || !publicFolderInfo.IsDir() {
		err = os.Mkdir("public", 0755)
		ErrorCheckExit(err)
	}
	postsFolderInfo, err := os.Stat(filepath.Join("public", "posts"))
	if err != nil || !postsFolderInfo.IsDir() {
		err = os.Mkdir(filepath.Join("public", "posts"), 0755)
		ErrorCheckExit(err)
	}
}

func writeHome(posts []post, config config) {
	homeHTMLBytes, err := ioutil.ReadFile(
		filepath.Join("templates", "home.html"),
	)
	ErrorCheckExit(err)
	homeHTML := formatStandard(string(homeHTMLBytes), config, posts)
	err = os.WriteFile(
		filepath.Join("public", "index.html"),
		[]byte(homeHTML), 0755,
	)
	ErrorCheckExit(err)
}

func writePosts(posts []post, config config) {
	postHTMLBytes, err := ioutil.ReadFile(
		filepath.Join("templates", "post.html"),
	)
	ErrorCheckExit(err)
	for _, post := range posts {
		postHTML := formatPost(string(postHTMLBytes), post, config)
		err = ioutil.WriteFile(
			filepath.Join("public", "posts",
				fmt.Sprintf("%s-%s.html",
					post.date, post.tag,
				)), []byte(postHTML), 0755)
		ErrorCheckExit(err)
	}
}

func writeRss(posts []post, config config) {
	postsRSSXMLBytes, err := ioutil.ReadFile(
		"templates/rss.xml",
	)
	ErrorCheckExit(err)
	postsRSSXML := formatStandard(string(postsRSSXMLBytes), config, posts)
	postsRSSXML = strings.ReplaceAll(
		postsRSSXML,
		"{{FF:PostRss:FF}}", formatRss(posts, config),
	)
	err = ioutil.WriteFile(
		filepath.Join("public", "rss.xml"),
		[]byte(postsRSSXML), 0755,
	)
	ErrorCheckExit(err)
}

func writeRes() {
	err := copy.Copy("res", filepath.Join("public", "res"))
	ErrorCheckExit(err)
}
