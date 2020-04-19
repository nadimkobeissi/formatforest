// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type config struct {
	WebsiteName        string
	WebsiteUri         string
	WebsiteDescription string
	WebsiteIcon        string
	WebsiteTwitter     string
	WebsiteLang        string
	WebsiteLangDir     string
	AuthorName         string
	AuthorEmail        string
	AuthorTwitter      string
	AuthorLinkedIn     string
	AuthorFacebook     string
	AuthorInstagram    string
}

func parseConfig() config {
	configJson := config{}
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		ErrorExit(err)
	}
	configText := string(configBytes)
	err = json.Unmarshal([]byte(configText), &configJson)
	if err != nil {
		ErrorExit(err)
	}
	if len(configJson.WebsiteName) == 0 {
		ErrorExit(errors.New("invalid WebsiteName in config.json"))
	}
	if len(configJson.WebsiteUri) == 0 {
		ErrorExit(errors.New("invalid WebsiteUri in config.json"))
	}
	if len(configJson.WebsiteIcon) == 0 {
		configJson.WebsiteIcon = "formatforest.png"
	}
	if len(configJson.WebsiteLang) == 0 {
		configJson.WebsiteLang = "en"
	}
	if len(configJson.WebsiteLangDir) == 0 {
		configJson.WebsiteLangDir = "ltr"
	}
	if len(configJson.AuthorName) == 0 {
		ErrorExit(errors.New("invalid AuthorName in config.json"))
	}
	return configJson
}

func parsePost(postMd string) (postConfig, string) {
	var postconfig postConfig
	postConfigText := strings.Join(strings.Split(postMd, "\n")[0:5], "\n")
	err := json.Unmarshal([]byte(postConfigText), &postconfig)
	if err != nil {
		ErrorExit(err)
	}
	postMdContent := strings.Join(strings.Split(postMd, "\n")[5:], "\n")
	return postconfig, postMdContent
}
