// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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

func postConfigParse(postMd string) (postConfig, string) {
	var postConfigJson postConfig
	postConfigText := strings.Join(strings.Split(postMd, "\n")[1:4], "\n")
	err := json.Unmarshal([]byte(postConfigText), &postConfigJson)
	if err != nil {
		errorExit(err)
	}
	postMdContent := strings.Join(strings.Split(postMd, "\n")[5:], "\n")
	return postConfigJson, postMdContent
}

func postRead(file os.FileInfo) post {
	fileBytes, err := ioutil.ReadFile(
		fmt.Sprintf("posts/%s", file.Name()),
	)
	if err != nil {
		errorExit(err)
	}
	postMd := string(fileBytes)
	if len(postMd) == 0 {
		errorExit(fmt.Errorf("could not read post at %s", file.Name()))
	}
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
	tagRegex := regexp.MustCompile(`\w{1,32}\.html`)
	date := dateRegex.FindString(file.Name())
	tag := tagRegex.FindString(file.Name())
	postConfig, postMdContent := postConfigParse(postMd)
	// TODO: validate date
	// TODO: validate tag
	// TODO: validate title
	// TODO: validate description
	// TODO: validate image
	return post{
		date:    date,
		tag:     tag[:len(tag)-5],
		config:  postConfig,
		content: postMdContent,
	}
}

func postReadAll() []post {
	posts := []post{}
	dirInfo, err := ioutil.ReadDir("posts")
	if err != nil {
		errorExit(err)
	}
	for _, file := range dirInfo {
		posts = append([]post{
			postRead(file),
		}, posts...)
	}
	return posts
}
