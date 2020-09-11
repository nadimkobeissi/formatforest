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
	WebsiteName         string
	WebsiteURI          string
	WebsiteDescription  string
	WebsiteIcon         string
	WebsiteTwitter      string
	WebsiteLang         string
	WebsiteLangDir      string
	AuthorName          string
	AuthorEmail         string
	AuthorTwitter       string
	AuthorLinkedIn      string
	AuthorFacebook      string
	AuthorInstagram     string
	CommentoIntegration bool
	SyncUser            string
	SyncHost            string
	SyncPath            string
}

func parseConfig() config {
	configJSON := config{}
	configBytes, err := ioutil.ReadFile("config.txt")
	ErrorCheckExit(err)
	configText := string(configBytes)
	err = json.Unmarshal([]byte(configText), &configJSON)
	ErrorCheckExit(err)
	if len(configJSON.WebsiteName) == 0 {
		ErrorExit(errors.New("invalid WebsiteName in config.txt"))
	}
	if len(configJSON.WebsiteURI) == 0 {
		ErrorExit(errors.New("invalid WebsiteURI in config.txt"))
	}
	if len(configJSON.WebsiteIcon) == 0 {
		configJSON.WebsiteIcon = "formatforest.png"
	}
	if len(configJSON.WebsiteLang) == 0 {
		configJSON.WebsiteLang = "en"
	}
	if len(configJSON.WebsiteLangDir) == 0 {
		configJSON.WebsiteLangDir = "ltr"
	}
	if len(configJSON.AuthorName) == 0 {
		ErrorExit(errors.New("invalid AuthorName in config.txt"))
	}
	return configJSON
}

func parsePost(postMd string) (postConfig, string) {
	var postconfig postConfig
	postConfigText := strings.Join(strings.Split(postMd, "\n")[0:5], "\n")
	err := json.Unmarshal([]byte(postConfigText), &postconfig)
	ErrorCheckExit(err)
	postMdContent := strings.Join(strings.Split(postMd, "\n")[5:], "\n")
	return postconfig, postMdContent
}
