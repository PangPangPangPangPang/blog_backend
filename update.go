package main

import (
	"bufio"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// Item `json:"xxx"`语法可以指定转JSON后的key
type Item struct {
	Date        string   `json:"time"`
	Tag         []string `json:"tags"`
	Title       string   `json:"title"`
	ID          string   `json:"id"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
}

// Items 实现按照时间排序Item
type Items []Item

func (l Items) Less(i, j int) bool {
	di := l[i].Date
	dj := l[j].Date
	format := "2006-01-02 15:04:05"
	ri, _ := time.Parse(format, di)
	rj, _ := time.Parse(format, dj)
	return ri.After(rj)
}

func (l Items) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Items) Len() int {
	return len(l)
}

// ListInitStatus 检查文章的初始化状态,如果为false,则生成文章列表等并缓存在内存中,如果为true,则从内存取
var ListInitStatus = true

// ListJSON 生成文章列表后的JSON数据
var ListJSON string

// Articles 文章map
var Articles map[string]Item

// ArticleList 获取未JSON化的文章数组
var ArticleList []Item

// Update force update article
func Update(c *gin.Context) {
	secretkey := c.Param("secretkey")
	envkey := os.Getenv("BBE_SECRET_KEY")
	if secretkey == envkey {
		// update
		stdout, err := updateStaticFile()
		ListInitStatus = true
		CheckUpdate()
		if nil != err {
			c.JSON(http.StatusOK, gin.H{
				"result":    "",
				"errorcode": 1,
				"errormsg":  "Update failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result":    stdout,
			"value":     ListJSON,
			"errorcode": 0,
			"errormsg":  ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":    "",
		"errorcode": 1,
		"errormsg":  "Error secret key"})
}

func updateStaticFile() (string, error) {
	path := WorkPath("scripts/update_bundle.sh")
	cmd := exec.Command("/bin/bash", path)
	stdout, err := cmd.Output()
	if nil != err {
		return "", err
	}
	return string(stdout), nil

}

// CheckUpdate 检查文章更新状态并且更新
func CheckUpdate() {

	if ListInitStatus {
		ListInitStatus = false
		list, m := GenerateList()
		Articles = m
		ArticleList = list
		ob, error := json.Marshal(&list)
		if error != nil {
		}
		ListJSON = string(ob)
		readArticlesIntoMemary()
		Rss()
	}
}

func readArticlesIntoMemary() {
	for k, v := range Articles {
		id := v.ID
		if v.Content == "" {
			articlesPath := WorkPath("articles")
			path := articlesPath + "/" + id + ".md"
			content, _ := ReadFile(path)
			v.Content = content
			Articles[k] = v
		}
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

	var list Items
	var m = make(map[string]Item)

	for _, content := range contents { // 格式化md文件并存在'./'articles文件夹中
		name := content.Name()
		rpath := resourcePath + "/" + name

		_, args, _ := convert(rpath, articlesPath)
		list = append(list, args)

		id := args.ID
		m[id] = args
	}
	sort.Sort(list)
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
	foundDes := false

	args.ID = MD5(title)

	for {
		istring, error := ireader.ReadString('\n')

		// 日期
		if !foundDate {
			trimString := strings.TrimSuffix(istring, "\n")
			if strings.HasPrefix(trimString, "[date]") {
				args.Date = strings.TrimPrefix(trimString, "[date]")
				args.Date = strings.TrimPrefix(args.Date, " ")
				foundDate = true
				continue
			}
		}
		// tags
		if !foundTag {
			trimString := strings.TrimSuffix(istring, "\n")
			if strings.HasPrefix(trimString, "[tag]") {
				arr := strings.Split(strings.TrimPrefix(trimString, "[tag]"), " ")
				args.Tag = arr[1:]
				foundTag = true
				continue
			}
		}
		// 简介
		if !foundDes {
			trimString := strings.TrimSuffix(istring, "\n")
			if strings.HasPrefix(trimString, "[description]") {
				args.Description = strings.TrimPrefix(trimString, "[description]")
				args.Description = strings.TrimPrefix(args.Description, " ")
				foundDes = true
				continue
			}
		}
		// 标题
		if !foundTitle {
			trimString := strings.TrimSuffix(istring, "\n")
			if strings.HasPrefix(trimString, "# ") {
				args.Title = strings.TrimPrefix(trimString, "# ")
				foundTitle = true
			}
		}

		owriter.WriteString(istring)

		if error == io.EOF {
			owriter.Flush()
			return ofile, args, nil
		}
	}
}
