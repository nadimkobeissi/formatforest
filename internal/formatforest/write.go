// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func writePosts(posts []post) {
	postHtmlBytes, err := ioutil.ReadFile(
		"templates/post.html",
	)
	if err != nil {
		errorExit(err)
	}
	for _, post := range posts {
		postHtml := postsRewrite(string(postHtmlBytes), post)
		err = ioutil.WriteFile(fmt.Sprintf(
			"publish/posts/%s-%s.html",
			post.date, post.tag,
		), []byte(postHtml), 0755)
		if err != nil {
			errorExit(err)
		}
	}
}

func writePostsRss(posts []post) {
	postsRssXmlBytes, err := ioutil.ReadFile(
		"templates/rss.xml",
	)
	if err != nil {
		errorExit(err)
	}
	postsRssXml := strings.ReplaceAll(
		string(postsRssXmlBytes), "{{RSSPOSTS}}", postsRss(posts),
	)
	err = ioutil.WriteFile(
		"publish/posts/rss.xml",
		[]byte(postsRssXml), 0755,
	)
	if err != nil {
		errorExit(err)
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
		string(homeHtmlBytes), "{{POSTS}}", postsList(posts),
	)
	err = ioutil.WriteFile(
		"publish/index.html",
		[]byte(homeHtml), 0755,
	)
	if err != nil {
		errorExit(err)
	}
}
