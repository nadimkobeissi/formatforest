// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type post struct {
	date    string
	tag     string
	title   string
	descr   string
	image   string
	content string
}

func postHarvest(file os.FileInfo) (post, bool) {
	fileBytes, err := ioutil.ReadFile(
		fmt.Sprintf("posts/%s", file.Name()),
	)
	if err != nil {
		errorExit(err)
	}
	content := string(fileBytes)
	if len(content) == 0 {
		return post{}, false
	}
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
	tagRegex := regexp.MustCompile(`\w{1,32}\.html`)
	titleRegex := regexp.MustCompile(`\[\[TITLE\]\]:.+`)
	descrRegex := regexp.MustCompile(`\[\[DESCR\]\]:.+`)
	imageRegex := regexp.MustCompile(`\[\[IMAGE\]\]:.+`)
	date := dateRegex.FindString(file.Name())
	tag := tagRegex.FindString(file.Name())
	title := titleRegex.FindString(content)
	descr := descrRegex.FindString(content)
	image := imageRegex.FindString(content)
	switch 0 {
	case len(date), len(tag), len(title), len(descr), len(image):
		return post{}, false
	}
	if err != nil {
		errorExit(err)
	}
	return post{
		date:    date,
		tag:     tag[:len(tag)-5],
		title:   title[10:],
		descr:   descr[10:],
		image:   image[10:],
		content: content,
	}, true
}

func postsHarvest() []post {
	posts := []post{}
	dirInfo, err := ioutil.ReadDir("posts")
	if err != nil {
		errorExit(err)
	}
	for _, file := range dirInfo {
		p, v := postHarvest(file)
		if v {
			posts = append([]post{p}, posts...)
		}
	}
	return posts
}

func postsRewrite(html string, post post) string {
	t, _ := time.Parse("2006-01-02", post.date)
	html = strings.ReplaceAll(
		html, "{{DATE}}", post.date,
	)
	html = strings.ReplaceAll(
		html, "{{RSSDATE}}", t.Format(time.RFC1123Z),
	)
	html = strings.ReplaceAll(
		html, "{{TAG}}", post.tag,
	)
	html = strings.ReplaceAll(
		html, "{{TITLE}}", post.title,
	)
	html = strings.ReplaceAll(
		html, "{{DESCR}}", post.descr,
	)
	html = strings.ReplaceAll(
		html, "{{IMAGE}}", post.image,
	)
	html = strings.ReplaceAll(
		html, "{{CONTENT}}", post.content,
	)
	return html
}

func postsList(posts []post) string {
	postsListHtml := []string{}
	for _, post := range posts {
		postsListHtml = append(postsListHtml, fmt.Sprintf(
			"<li><em>%s:</em> <a href=\"posts/%s-%s.html\">%s</a></li>",
			post.date, post.date, post.tag, post.title,
		))
	}
	return strings.Join(postsListHtml, "\n")
}

func postsRss(posts []post) string {
	postsRssXml := []string{}
	for _, post := range posts {
		postRssXml := strings.Join([]string{
			"<item>",
			"<title>{{TITLE}}</title>",
			"<link>https://nadim.computer/posts/{{DATE}}-{{TAG}}.html</link>",
			"<dc:creator><![CDATA[Nadim Kobeissi]]></dc:creator>",
			"<pubDate>{{RSSDATE}}</pubDate>",
			"<description><![CDATA[{{DESCR}}]]></description>",
			"<content:encoded><![CDATA[{{CONTENT}}]]></content:encoded>",
			"<media:thumbnail url=\"https://nadim.computer/posts/res/img/{{IMAGE}}\" />",
			"</item>",
		}, "\n")
		postRssXml = postsRewrite(postRssXml, post)
		postsRssXml = append(postsRssXml, postRssXml)
	}
	return strings.Join(postsRssXml, "\n")
}
