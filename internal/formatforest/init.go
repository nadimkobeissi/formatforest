// SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
// SPDX-License-Identifier: GPL-3.0-only

package formatforest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Init(blogFolder string) {
	_, err := os.Stat(blogFolder)
	if err == nil {
		ErrorExit(errors.New("folder already exists"))
	}
	fmt.Printf("[FormatForest] Initializing folder %s...", blogFolder)
	writeInitFolders(blogFolder)
	fmt.Println(" OK")
	fmt.Print("[FormatForest] Writing default config.txt...")
	initDefaultConfig(blogFolder)
	fmt.Println(" OK")
	fmt.Print("[FormatForest] Downloading default assets...")
	downloadInitAssets(blogFolder)
	fmt.Println(" OK")
}

func initDefaultConfig(blogFolder string) {
	configJson := config{
		WebsiteName:        "My FormatForest Blog",
		WebsiteUri:         "https://myblog.com",
		WebsiteDescription: "Welcome to my personal blog.",
		WebsiteIcon:        "formatforest.png",
		WebsiteTwitter:     "forestformat",
		WebsiteLang:        "en",
		WebsiteLangDir:     "ltr",
		AuthorName:         "Format Gardener",
		AuthorEmail:        "your@email.com",
		AuthorTwitter:      "yourtwitter",
		AuthorLinkedIn:     "yourlinkedinid",
		AuthorFacebook:     "yourfacebookid",
		AuthorInstagram:    "yourinstagramid",
		SyncUser:           "user",
		SyncHost:           "myblog.com",
		SyncPath:           "/var/www/myblog.com",
	}
	configJsonBytes, err := json.MarshalIndent(configJson, "", "\t")
	ErrorCheckExit(err)
	err = ioutil.WriteFile(path.Join(blogFolder, "config.txt"), configJsonBytes, 0755)
	ErrorCheckExit(err)
}
