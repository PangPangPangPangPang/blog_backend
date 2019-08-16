package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/feeds"
	"gopkg.in/russross/blackfriday.v2"
	"os"
	"time"
)

// Rss rss generator.
func Rss() {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Max's Blog",
		Link:        &feeds.Link{Href: "https://mmmmmax.cn"},
		Description: "分享一下自己乱七八糟的生活",
		Author:      &feeds.Author{Name: "Max", Email: "im.yf.wang@gmail.com"},
		Created:     now,
	}
	list := ArticleList
	feed.Items = []*feeds.Item{}
	for i := 0; i < len(list); i++ {
		item := list[i]
		id := item.ID

		content := []byte(Articles[id].Content)
		des := string(blackfriday.Run(content, blackfriday.WithExtensions(blackfriday.FencedCode|blackfriday.Tables|blackfriday.SpaceHeadings|blackfriday.Strikethrough)))
		create, _ := time.Parse("2006-01-02 15:04:05 -0700", fmt.Sprintf(`%s +0800`, item.Date))
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("https://mmmmmax.cn/#list/%s", item.ID)},
			Description: string(des),
			Created:     create,
		})
	}
	rss, err := feed.ToRss()
	if err != nil {
		fmt.Printf("Generate RSS content error.")
	}

	path := WorkPath("static")

	ofile, err := os.OpenFile(path+"/"+"feed", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("Generate RSS file error.")
	}
	defer ofile.Close()
	owriter := bufio.NewWriter(ofile)
	owriter.WriteString(rss)
	owriter.Flush()
}
