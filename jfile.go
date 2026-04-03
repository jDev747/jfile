package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type jItem struct {
	Name  string
	Content string
	Files []jItem
}


func isJdir(item jItem) bool {
    return len(item.Files) > 0
}
func fileToJfile(path string) *jItem {
	filename := filepath.Base(path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(bytes)
	return &jItem{Name: filename, Content: content}
}
func jFileToFile(jfile jItem, exportPath string) {
	if isJdir(jfile) {
		log.Fatal("<FROM JFILE> ERROR: passed jdir to a jfile function (jfiletofile)")
	}
	fullpath := filepath.Join(exportPath, jfile.Name)
	os.WriteFile(fullpath, []byte(jfile.Content), 0777)
}
func dirToJdir(path string) *jItem {
	dirpath := filepath.Base(path)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var jfiles []jItem
	for _, file := range files {
		fname := filepath.Join(path, file.Name())
		if file.IsDir() {
			jfiles = append(jfiles, *dirToJdir(fname))
		} else {
			jfiles = append(jfiles, *fileToJfile(fname))
		}
	}
	return &jItem{Name: dirpath, Files: jfiles}
}
func JdirTodir(jdir jItem, exportPath string) {
	if !isJdir(jdir) {
		log.Fatal("<FROM JFILE> ERROR: passed jfile to a jdir function (jdirtodir)")
	}
	fullpath := filepath.Join(exportPath, jdir.Name)
	fmt.Println(jdir)
	os.Mkdir(exportPath, 0777)
	os.Mkdir(fullpath, 0777)
	for _, item := range jdir.Files {
		fmt.Println(item)
		if isJdir(item){
			JdirTodir(item, fullpath)
		} else {
			jFileToFile(item, fullpath)
			}
		}
	}
func readJson(path string) jItem {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	obj := jItem{}
	json.Unmarshal(bytes, &obj)
	fmt.Println(obj)
	return obj
}
func WriteJson(path string, obj any) {
	bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(path, bytes, 0777)
}
func main() {
	obj := readJson("main.json")
	fmt.Println(obj)
	JdirTodir(obj, "EXPORT")
}


// i need to completely rewrite this!!!!!!!!
// i just did ;)