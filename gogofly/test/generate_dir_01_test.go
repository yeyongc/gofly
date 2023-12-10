package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type JSONField struct {
	name     string
	ftype    string
	children []JSONField
}

var (
	stRootDir   string
	stSeparator string
	jsonData    map[string]interface{}
)

const filename = "dir.json"

func loadJSON() {
	stWorkDir, _ := os.Getwd()                                        // 当前工作路径
	stSeparator = string(filepath.Separator)                          // 路径分隔符
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)] // 获取项目根目录

	// 加载json文件
	gnJSONBytes, _ := os.ReadFile(stWorkDir + stSeparator + filename)
	err := json.Unmarshal(gnJSONBytes, &jsonData)

	if err != nil {
		panic("Laod JSON Data Error: " + err.Error())
	}
}

func parseMap() JSONField {
	if jsonData == nil {
		return JSONField{}
	}
	jsonFields := JSONField{}
	for k, v := range jsonData {
		switch k {
		case "children":
			t, _ := v.([]interface{})
			res := []JSONField{}
			for _, child := range t {
				res = append(res, parseItem(child))
			}
			jsonFields.children = res
		case "name":
			jsonFields.name, _ = v.(string)
		case "ftype":
			jsonFields.ftype, _ = v.(string)
		}
	}
	return jsonFields
}

func createDirsOrFiles(jsf JSONField, parentPath string) {
	if jsf.name != "" {
		switch jsf.ftype {
		case "file":
			_, err := os.Create(parentPath + stSeparator + jsf.name)
			if err != nil {
				panic("Create file error" + err.Error())
			}
		case "dir":
			err := os.Mkdir(parentPath+stSeparator+jsf.name, fs.ModePerm)
			if err != nil {
				panic("Create file error" + err.Error())
			}
		}
	}

	for _, child := range jsf.children {
		createDirsOrFiles(child, parentPath+stSeparator+jsf.name)
	}
}

func parseItem(field interface{}) JSONField {
	tmap, _ := field.(map[string]interface{})
	jsf := JSONField{}
	for k, v := range tmap {
		switch k {
		case "name":
			jsf.name, _ = v.(string)
		case "type":
			jsf.ftype, _ = v.(string)
		case "children":
			t, _ := v.([]interface{})
			res := []JSONField{}
			for _, e := range t {
				res = append(res, parseItem(e))
			}
			jsf.children = res
		}
	}
	return jsf
}

func TestGenerateDir01(t *testing.T) {
	loadJSON()
	jsf := parseMap()
	createDirsOrFiles(jsf, stRootDir)
}
