package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type JSONNode struct {
	Children []JSONNode `json:children`
	Type     string     `json:type`
	Name     string     `json:name`
}

var (
	stRootDir02   string
	stSeparator02 string
	stRootNode    JSONNode
)

const stJSONFileName = "dir.json"

func generateJSNode() {
	stWorkDir, _ := os.Getwd()
	stSeparator02 = string(filepath.Separator)
	stRootDir02 = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator02)]

	jsFileDataBytes, _ := os.ReadFile(stWorkDir + stSeparator02 + stJSONFileName)

	err := json.Unmarshal(jsFileDataBytes, &stRootNode)

	if err != nil {
		panic("Parse json file error: " + err.Error())
	}
}

func createDirorFileByNode(node JSONNode, parentPath string) {
	path := parentPath + node.Name
	fmt.Println(path)
	if node.Name != "" {
		switch node.Type {
		case "file":
			_, err := os.Create(path)
			if err != nil {
				panic("Create file error:" + err.Error())
			}
		case "dir":
			err := os.Mkdir(path, fs.ModePerm)
			if err != nil {
				panic("Create dir error: " + err.Error())
			}
		}
	}

	for _, child := range node.Children {
		createDirorFileByNode(child, path+stSeparator02)
	}
}

func TestGenerate02(t *testing.T) {
	generateJSNode()
	createDirorFileByNode(stRootNode, stRootDir02)
}
