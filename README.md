<!---
# SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@nadim.computer>
# SPDX-License-Identifier: CC-BY-SA-4.0
-->

# [FormatForest](https://formatforest.com)

FormatForest is a simple and elegant blogging engine written in Go. It was originally developed as a personal project targeting Nadim Kobeissi's [personal blog](https://nadim.computer) but has now been rendered open source.

FormatForest is only suited for personal blogs operated by a single individual. Its advantages:

- **Simplicity and elegance** in the FormatForest design and functionality allow you to focus on growing your writing and personal web presence.
- **Complete Markdown support** allows you to draft posts quickly and without fuss.
- **Unparallelled security** compared to WordPress and other dynamic blog engines.
- **Secure sync over rsync** allows you to publish your website quickly from your local machine.

## Getting Started
Let's set up a blog using FormatForest.

### Step 1: Initializing a Blog
1. Install FormatForest as a system command: `go get -u formatforest.com/cmd/formatforest`
2. Type `formatforest init myBlog` to create a new FormatForest blog in a folder called `myBlog`.

The folder `myBlog` will contain:

- `config.txt`, your master configuration file for your FormatForest blog.
- `posts` folder, where you will write and store your blog posts in Markdown format.
- `public` folder, where FormatForest will generate static HTML to be served via your web server.
- `res` folder, which contains your static resources (includes `img`, `css` and `js` subfolders).
- `templates` folder, which includes the templates that FormatForest will use to render your formatted static blog.

Please note that `init` requires an Internet connection in order to download some very small assets for your blog. If you do not have an Internet connection, you can simply create the above files and folders manually and obtain the same effect, minus the boilerplate HTML templates, CSS and images.

### Step 2: Configuring Your Blog
Simply open `config.txt` using a text editor and modify the values. They are self-explanatory:

- `WebsiteName`: Your website's name.
- `WebsiteUri`: Your website's URL.
- `WebsiteDescription`: A small description for your website, one paragraph.
- `WebsiteIcon`: An icon for your website. Store the icon file in `res/img`. For example, if your icon is located in `res/img/myIcon.png`, then simply set this parameter to `"myIcon.png"`. Do not type the full path.
- `WebsiteTwitter`: Your Twitter handle. Do not include the `@`.
- `WebsiteLang`: Your website's language.
- `WebsiteLangDir`: Set to `ltr` for left-to-right languages (such as English or French), `rtl` for right-to-left languages (such as Arabic or Hebrew.)
- `AuthorName`: Your name.
- `AuthorEmail`: Your email.
- `AuthorTwitter`: Your Twitter handle, otherwise leave empty.
- `AuthorLinkedIn`: Your LinkedIn handle, otherwise leave empty.
- `AuthorFacebook`: Your Facebook handle, otherwise leave empty.
- `AuthorInstagram`: Your Instagram handle, otherwise leave empty.
- `SyncUser`: SSH username for FormatForest's sync feature.
- `SyncHost`: Remote host for FormatForest's sync feature.
- `SyncPath`: Remote directory path for FormatForest's sync feature.

### Step 3: Write Your First Blog Post
Blog posts are stored in the `posts` folder and **must imperatively** have a filename with the following format:
```
2020-04-19-mypost.md
```

In the above, we have the date of publication (year-month-day) and then a *post tag* which indicates a single-word unique identifier for the post which must be between 1 and 32 characters. `.md` specifies that this is a Markdown file. If your post filename does not meet the above format, FormatForest will not be able to process your blog.

Here is an example post:

```
{
	PostTitle: "My First Post",
	PostDescription: "A short description of my post for social media embeds.",
	PostImage: "fileNameOfMyImageInResImgFolderForSocialMediaEmbeds.jpg"
}

## Subtitle
Hello, and welcome to my post! Thanks for tuning in!

## Conclusion
Final Fantasy VII Remake is a fantastic game!
```

In FormatForest, the first five lines of a post specify a configuration for that post itself. This must be specified in order for the post to be processed correctly. The rest of the file can be standard Markdown.

### Step 4: Publish Your Blog
In order to generate the static HTML for your blog, return to your `myBlog` folder and type: `formatforest format`. Your website will then be ready inside the `public` folder.

You may then either:

- Serve the `public` folder locally from your computer as you see fit, or manually upload it somewhere, or,
- Type `formatforest sync` in order to automatically synchronize your `public` folder via rsync and SSH to a remote server. rsync must be installed locally for this to work and correct values for `SyncUser`, `SyncHost` and `SyncPath` must be specified in `config.txt`.

Easy!

## Customizing Your Blog with Templates
The `templates` folder contains three templates:

- `home.html`: used for your homepage and post listing.
- `post.html`: used for actual posts.
- `rss.xml`: used to generate your blog's RSS feed.

These templates can be fully customized to your liking. The CSS file can be found and modified inside `res/css/style.css`.

### Template Tags
FormatForest will find the following **template tags** inside HTML templates and replace them automatically with the appropriate content. Each tag is written in the HTML file in the format `{{FF:TagName:FF}}`. For example, the `PostDate` tag would be written as `{{FF:PostDate:FF}}`.

#### Post Tags
- `PostDate`: The date of publication of the post (yyyy-mm-dd), derived from its filename.
- `PostTag`: The tag of the post, derived from its filename.
- `PostContent`: The contents of the post, converted from Markdown to HTML format.
- `PostList`: An HTML list of all the posts in the blog, with links.

#### RSS Tags
- `PostRss`: Your posts in RSS format.
- `PostRssDate`: The publication date for your post in RSS-compatible format.

#### Miscellaneous Tags
- `CurrentYear`: The current year.
- `CurrentMonth`: The current month in numeric format.
- `CurrentDay`: The current day in numeric format.

## Discussion
Please check out the [FormatForest source code repository](https://gitlab.com/nadimk/formatforest) to discuss and contribute.

## License
FormatForest is authored by Nadim Kobeissi. It is provided as free and open source software, licensed under the [GNU General Public License, version 3](https://www.gnu.org/licenses/gpl-3.0.en.html).

© Copyright 2020-2021 Nadim Kobeissi. All Rights Reserved.