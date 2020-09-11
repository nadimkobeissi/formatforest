// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"fmt"
	"strings"
	"time"
)

// Format formats the blog resources into publishable HTML files.
func Format() {
	fmt.Printf("[FormatForest] Parsing config.txt...")
	config := parseConfig()
	fmt.Println(" OK")
	fmt.Printf("[FormatForest] Reading posts...")
	posts := postReadAll()
	fmt.Println(" OK")
	fmt.Printf("[FormatForest] Writing public folders...")
	writePublicFolders()
	fmt.Println(" OK")
	fmt.Printf("[FormatForest] Writing home...")
	writeHome(posts, config)
	fmt.Println(" OK")
	fmt.Printf("[FormatForest] Writing posts...")
	writePosts(posts, config)
	fmt.Println(" OK")
	fmt.Printf("[FormatForest] Writing RSS...")
	writeRss(posts, config)
	fmt.Println(" OK")
	fmt.Printf("[FormatForest] Writing resources...")
	writeRes()
	fmt.Println(" OK")
}

func formatStandard(html string, config config, posts []post) string {
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
		html, "{{FF:WebsiteURI:FF}}", config.WebsiteURI,
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
	html = strings.ReplaceAll(
		html, "{{FF:PostList:FF}}", formatPostList(posts),
	)
	return html
}

func formatPost(html string, p post, config config) string {
	html = formatStandard(html, config, []post{})
	t, _ := time.Parse("2006-01-02", p.date)
	commentoHead := ""
	if config.CommentoIntegration {
		commentoHead = fmt.Sprintf(
			"<script defer src=\"%s\" data-page-id=\"%s-%s\"></script>",
			"https://cdn.commento.io/js/commento.js", p.date, p.tag,
		)
	}
	html = strings.ReplaceAll(
		html, "{{FF:PostDate:FF}}", p.date,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostRssDate:FF}}", t.Format(time.RFC1123Z),
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostTag:FF}}", p.tag,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostTitle:FF}}", p.config.PostTitle,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostDescription:FF}}", p.config.PostDescription,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostImage:FF}}", p.config.PostImage,
	)
	html = strings.ReplaceAll(
		html, "{{FF:PostContent:FF}}", p.content,
	)
	html = strings.ReplaceAll(
		html, "{{FF:CommentoHead:FF}}", commentoHead,
	)
	return html
}

func formatPostList(posts []post) string {
	postListHTML := []string{}
	for _, post := range posts {
		postListHTML = append(postListHTML, fmt.Sprintf(
			"<li><em>%s:</em> <a href=\"/posts/%s-%s.html\">%s</a></li>",
			post.date, post.date, post.tag, post.config.PostTitle,
		))
	}
	return fmt.Sprintf("<ul>%s</ul>", strings.Join(postListHTML, "\n"))
}

func formatRss(posts []post, config config) string {
	postsRSSXML := []string{}
	for _, p := range posts {
		postRSSXML := strings.Join([]string{
			"<item>",
			"<title>{{FF:PostTitle:FF}}</title>",
			"<link>{{FF:WebsiteURI:FF}}/posts/{{FF:PostDate:FF}}-{{FF:PostTag:FF}}.html</link>",
			"<dc:creator><![CDATA[{{FF:AuthorName:FF}}]]></dc:creator>",
			"<pubDate>{{FF:PostRssDate:FF}}</pubDate>",
			"<description><![CDATA[{{FF:PostDescription:FF}}]]></description>",
			"<content:encoded><![CDATA[{{FF:PostContent:FF}}]]></content:encoded>",
			"<media:thumbnail url=\"{{FF:WebsiteURI:FF}}/res/img/{{FF:PostImage:FF}}\" />",
			"</item>",
		}, "\n")
		postRSSXML = formatStandard(postRSSXML, config, posts)
		postRSSXML = formatPost(postRSSXML, p, config)
		postsRSSXML = append(postsRSSXML, postRSSXML)
	}
	return strings.Join(postsRSSXML, "\n")
}
