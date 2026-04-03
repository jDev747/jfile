package jfile

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type JItem struct {
	Name  string
	Content []byte
	Files []JItem
}


func IsJdir(item JItem) bool {
    return len(item.Files) > 0
}
func FileToJfile(path string) *JItem {
	filename := filepath.Base(path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return &JItem{Name: filename, Content: bytes}
}
func JFileToFile(jfile JItem, exportPath string) {
	if IsJdir(jfile) {
		log.Fatal("<FROM JFILE> ERROR: passed jdir to a jfile function (jfiletofile)")
	}
	fullpath := filepath.Join(exportPath, jfile.Name)
	os.WriteFile(fullpath, jfile.Content, 0777)
}
func DirToJdir(path string) *JItem {
	dirpath := filepath.Base(path)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var jfiles []JItem
	for _, file := range files {
		fname := filepath.Join(path, file.Name())
		if file.IsDir() {
			jfiles = append(jfiles, *DirToJdir(fname))
		} else {
			jfiles = append(jfiles, *FileToJfile(fname))
		}
	}
	return &JItem{Name: dirpath, Files: jfiles}
}
func JdirTodir(jdir JItem, exportPath string) {
	if !IsJdir(jdir) {
		log.Fatal("<FROM JFILE> ERROR: passed jfile to a jdir function (jdirtodir)")
	}
	fullpath := filepath.Join(exportPath, jdir.Name)
	os.MkdirAll(fullpath, 0777)
	for _, item := range jdir.Files {
		if IsJdir(item){
			JdirTodir(item, fullpath)
		} else {
			JFileToFile(item, fullpath)
			}
		}
	}
func ReadJson(path string) JItem {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	obj := JItem{}
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		log.Fatal(err)
	}
	return obj
}
func WriteJson(path string, obj any) {
	bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(path, bytes, 0777)
}


// have fun with the module yall
