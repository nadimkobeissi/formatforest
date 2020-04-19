// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"strings"
	"time"
)

func Format() {
	posts := postReadAll()
	writeFolders()
	writeHome(posts)
	writePosts(posts)
	writeRss(posts)
}

func formatPost(html string, post post) string {
	t, _ := time.Parse("2006-01-02", post.date)
	html = strings.ReplaceAll(
		html, "{{FF:PostDate:FF}}", post.date,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostRssDate:FF}}", t.Format(time.RFC1123Z),
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostTag:FF}}", post.tag,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostTitle:FF}}", post.config.PostTitle,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostDescription:FF}}", post.config.PostDescription,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostImage:FF}}", post.config.PostImage,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostContent:FF}}", post.content,
	)
	return html
}

func formatPostList(posts []post) string {
	postListHtml := []string{}
	for _, post := range posts {
		postListHtml = append(postListHtml, fmt.Sprintf(
			"<li><em>%s:</em> <a href=\"posts/%s-%s.html\">%s</a></li>",
			post.date, post.date, post.tag, post.config.PostTitle,
		))
	}
	return strings.Join(postListHtml, "\n")
}

func formatRss(posts []post) string {
	postsRssXml := []string{}
	for _, post := range posts {
		postRssXml := strings.Join([]string{
			"<item>",
			"<title>{{FF:PostTitle:FF}}</title>",
			"<link>https://nadim.computer/posts/{{FF:PostDate:FF}}-{{FF:PostTag:FF}}.html</link>",
			"<dc:creator><![CDATA[Nadim Kobeissi]]></dc:creator>",
			"<pubDate>{{FF:PostRssDate:FF}}</pubDate>",
			"<description><![CDATA[{{FF:PostDescription:FF}}]]></description>",
			"<content:encoded><![CDATA[{{FF:PostContent:FF}}]]></content:encoded>",
			"<media:thumbnail url=\"{{FF:WebsiteUri:FF}}/res/img/{{FF:PostImage:FF}}\" />",
			"</item>",
		}, "\n")
		postRssXml = formatPost(postRssXml, post)
		postsRssXml = append(postsRssXml, postRssXml)
	}
	return strings.Join(postsRssXml, "\n")
}
