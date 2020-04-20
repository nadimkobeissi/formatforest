// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"strings"
	"time"
)

func Format() {
	fmt.Printf("[FormatForest] Formatting...")
	config := parseConfig()
	posts := postReadAll()
	writeFolders()
	writeHome(posts, config)
	writePosts(posts, config)
	writeRss(posts, config)
	writeRes()
	fmt.Println(" OK")
}

func formatStandard(html string, config config) string {
	y, m, d := time.Now().Date()
	year := fmt.Sprintf("%d", y)
	month := fmt.Sprintf("%d", m)
	day := fmt.Sprintf("%d", d)
	if m < 10 {
		month = fmt.Sprintf("0%d", m)
	}
	if d < 10 {
		day = fmt.Sprintf("0%d", d)
	}
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteName:FF}}", config.WebsiteName,
	)
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteUri:FF}}", config.WebsiteUri,
	)
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteDescription:FF}}", config.WebsiteDescription,
	)
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteIcon:FF}}", config.WebsiteIcon,
	)
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteTwitter:FF}}", config.WebsiteTwitter,
	)
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteLang:FF}}", config.WebsiteLang,
	)
	html = strings.ReplaceAll(
		html, "{{FF:WebsiteLangDir:FF}}", config.WebsiteLangDir,
	)
	html = strings.ReplaceAll(
		html, "{{FF:AuthorName:FF}}", config.AuthorName,
	)
	html = strings.ReplaceAll(
		html, "{{FF:AuthorEmail:FF}}", config.AuthorEmail,
	)
	html = strings.ReplaceAll(
		html, "{{FF:AuthorTwitter:FF}}", config.AuthorTwitter,
	)
	html = strings.ReplaceAll(
		html, "{{FF:AuthorLinkedIn:FF}}", config.AuthorLinkedIn,
	)
	html = strings.ReplaceAll(
		html, "{{FF:AuthorFacebook:FF}}", config.AuthorFacebook,
	)
	html = strings.ReplaceAll(
		html, "{{FF:AuthorInstagram:FF}}", config.AuthorInstagram,
	)
	html = strings.ReplaceAll(
		html, "{{FF:CurrentYear:FF}}", year,
	)
	html = strings.ReplaceAll(
		html, "{{FF:CurrentMonth:FF}}", month,
	)
	html = strings.ReplaceAll(
		html, "{{FF:CurrentDay:FF}}", day,
	)
	return html
}

func formatPost(html string, post post, config config) string {
	html = formatStandard(html, config)
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
			"<li><em>%s:</em> <a href=\"/posts/%s-%s.html\">%s</a></li>",
			post.date, post.date, post.tag, post.config.PostTitle,
		))
	}
	return fmt.Sprintf("<ul>%s</ul>", strings.Join(postListHtml, "\n"))
}

func formatRss(posts []post, config config) string {
	postsRssXml := []string{}
	for _, post := range posts {
		postRssXml := strings.Join([]string{
			"<item>",
			"<title>{{FF:PostTitle:FF}}</title>",
			"<link>{{FF:WebsiteUri:FF}}/posts/{{FF:PostDate:FF}}-{{FF:PostTag:FF}}.html</link>",
			"<dc:creator><![CDATA[{{FF:WebsiteAuthor:FF}}]]></dc:creator>",
			"<pubDate>{{FF:PostRssDate:FF}}</pubDate>",
			"<description><![CDATA[{{FF:PostDescription:FF}}]]></description>",
			"<content:encoded><![CDATA[{{FF:PostContent:FF}}]]></content:encoded>",
			"<media:thumbnail url=\"{{FF:WebsiteUri:FF}}/res/img/{{FF:PostImage:FF}}\" />",
			"</item>",
		}, "\n")
		postRssXml = formatPost(postRssXml, post, config)
		postsRssXml = append(postsRssXml, postRssXml)
	}
	return strings.Join(postsRssXml, "\n")
}
