package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// `json:"xxx"`语法可以指定转JSON后的key
type Item struct {
	Date    string   `json:"time"`
	Tag     []string `json:"tags"`
	Title   string   `json:"title"`
	ID      string   `json:"id"`
	Content string   `json:"content"`
}

var ListInitStatus = true
var listJSON string
var Articles map[string]Item

//List 获取文章列表
func List(c *gin.Context) {
	if ListInitStatus {
        ListInitStatus = false
		list, m := GenerateList()
		Articles = m
		ob, error := json.Marshal(&list)
		if error != nil {
			fmt.Println(error)
        }
		listJSON = string(ob)
	}
	c.JSON(http.StatusOK, gin.H{
		"result":   listJSON,
		"errorno":  0,
		"errormsg": ""})
}

// WorkPath a
func WorkPath(file string) string{
	home := os.Getenv("HOME") + "/blog_backend"
    return home + "/" + file
}

// GenerateList ggg
func GenerateList() ([]Item, map[string]Item) {
	resourcePath := WorkPath("resource")
    articlesPath := WorkPath("Items")

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

	for _, content := range contents {
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

		if !foundDate {
			if strings.HasPrefix(istring, "[date]") {
				args.Date = istring
				foundDate = true
				continue
			}
		}
		if !foundTag {
			if strings.HasPrefix(istring, "[tag]") {
				arr := strings.Split(istring, " ")
				args.Tag = arr[1:]
				foundTag = true
				continue
			}
		}
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

// MD5 md5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
