// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
)

type postConfig struct {
	PostTitle       string
	PostDescription string
	PostImage       string
}
type post struct {
	date    string
	tag     string
	config  postConfig
	content string
}

func postRead(file os.FileInfo) post {
	fileBytes, err := ioutil.ReadFile(
		path.Join("posts", file.Name()),
	)
	if err != nil {
		ErrorExit(err)
	}
	postMd := string(fileBytes)
	if len(postMd) == 0 {
		ErrorExit(fmt.Errorf("could not read post at %s", file.Name()))
	}
	fileName := strings.TrimSuffix(file.Name(), ".md")
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
	tagRegex := regexp.MustCompile(`\w{1,32}$`)
	date := dateRegex.FindString(fileName)
	tag := tagRegex.FindString(fileName)
	postConfig, postMdContent := parsePost(postMd)
	postHtmlContent := string(markdown.ToHTML([]byte(postMdContent), nil, nil))
	if len(date) == 0 {
		ErrorExit(errors.New("date must be in yyyy-mm-dd format"))
	}
	if len(tag) == 0 {
		ErrorExit(errors.New("tag must be a single word between 1 and 32 characters"))
	}
	return post{
		date:    date,
		tag:     tag,
		config:  postConfig,
		content: postHtmlContent,
	}
}

func postReadAll() []post {
	posts := []post{}
	dirInfo, err := ioutil.ReadDir("posts")
	if err != nil {
		ErrorExit(err)
	}
	for _, file := range dirInfo {
		posts = append([]post{
			postRead(file),
		}, posts...)
	}
	return posts
}
