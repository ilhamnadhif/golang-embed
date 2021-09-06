package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"reflect"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed logo.jpg
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpeg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var file embed.FS

func TestMultipleFile(t *testing.T) {
	a, _ := file.ReadFile("files/a.txt")
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)
	fmt.Println(string(a))
	b, _ := file.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := file.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

func TestTestDoang(t *testing.T) {
	var arrName []uint8 = []uint8{65, 67, 68, 69, 70}
	fmt.Println(string(arrName))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
