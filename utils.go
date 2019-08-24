package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

// WorkPath generate workpath.
func WorkPath(file string) string {
	current, _ := os.Getwd()
	return current + "/" + file
}

// MD5 md5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// VolumnPath t
func VolumnPath(file string) string {
	s := os.Getenv("GIN_MODE")
	if s == "release" {
		return "/blog_backend/" + file
	}
	home := os.Getenv("HOME") + "/blog_backend"
	return home + "/" + file
}
