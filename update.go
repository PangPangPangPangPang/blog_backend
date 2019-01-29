package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Item `json:"xxx"`语法可以指定转JSON后的key
type Item struct {
	Date    string   `json:"time"`
	Tag     []string `json:"tags"`
	Title   string   `json:"title"`
	ID      string   `json:"id"`
	Content string   `json:"content"`
}

// ListInitStatus 检查文章的初始化状态,如果为false,则生成文章列表等并缓存在内存中,如果为true,则从内存取
var ListInitStatus = true
// ListJSON 生成文章列表后的JSON数据
var ListJSON string
// Articles 文章map
var Articles map[string]Item

// CheckUpdate 检查文章更新状态并且更新
func CheckUpdate() {
	if ListInitStatus {
        ListInitStatus = false
		list, m := GenerateList()
		Articles = m
		ob, error := json.Marshal(&list)
		if error != nil {
        }
		ListJSON = string(ob)
	}
}

// GenerateList generate article list
func GenerateList() ([]Item, map[string]Item) {
	resourcePath := WorkPath("resource")
    articlesPath := WorkPath("articles")

	// create artitles dir.
	if err := os.MkdirAll(articlesPath, 0777); err != nil {
		return nil, nil
	}

	contents, err := ioutil.ReadDir(resourcePath)
	if err != nil {
		return nil, nil
	}

	var list []Item
    var m = make(map[string]Item)

	for _, content := range contents { // 格式化md文件并存在'./'articles文件夹中
		name := content.Name()
		rpath := resourcePath + "/" + name

		_, args, _ := convert(rpath, articlesPath)
		list = append(list, args)

		id := args.ID
        m[id] = args
	}
	return list, m
}

func convert(path, dir string) (*os.File, Item, error) {
	var args Item
	ifile, ierror := os.Open(path)
	if ierror != nil {
		return nil, args, ierror
	}
	defer ifile.Close()

	ireader := bufio.NewReader(ifile)

	parms := strings.Split(path, "/")
	title := parms[len(parms)-1]
	ofile, oerror := os.OpenFile(dir+"/"+MD5(title)+".md", os.O_CREATE|os.O_RDWR, 0777)
	if oerror != nil {
		return nil, args, oerror
	}
	defer ofile.Close()

	owriter := bufio.NewWriter(ofile)

	foundTitle := false
	foundDate := false
	foundTag := false

	args.ID = MD5(title)

	for {
		istring, error := ireader.ReadString('\n')

        // 日期
		if !foundDate {
			if strings.HasPrefix(istring, "[date]") {
				args.Date = istring
				foundDate = true
				continue
			}
		}
        // tags
		if !foundTag {
			if strings.HasPrefix(istring, "[tag]") {
				arr := strings.Split(istring, " ")
				args.Tag = arr[1:]
				foundTag = true
				continue
			}
		}
        // 标题
		if !foundTitle {
			if strings.HasPrefix(istring, "# ") {
				args.Title = istring
				foundTitle = true
			}
		}

		owriter.WriteString(istring)

		if error == io.EOF {
			owriter.Flush()
			return ofile, args, nil
		}
	}
	return ofile, args, nil
}
