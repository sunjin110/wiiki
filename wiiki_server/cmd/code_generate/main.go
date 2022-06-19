package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"wiiki_server/common/rootdir"
	"wiiki_server/common/testtool"
	"wiiki_server/common/utils/fileutil"
)

type RepoModelStruct struct {
	ModelName string
	FieldList []*Field
}

type Field struct {
	Name string
	Type string
	Tag  string
}

func main() {

	// get root dir
	rootDir := rootdir.Dir()
	repoModelDir := filepath.Join(rootDir, "domain", "model", "repomodel")

	repoModelFileList, err := fileutil.WalkDir(repoModelDir)
	if err != nil {
		panic(err)
	}
	testtool.Log("repo model files is ", repoModelFileList)

	for _, repoModelFilePath := range repoModelFileList {
		GetRepoModelStructList(repoModelFilePath)
	}

	// code parser

}

func GetRepoModelStructList(filePath string) ([]*RepoModelStruct, error) {

	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var lintList []string
	for scanner.Scan() {
		line := scanner.Text()
		lintList = append(lintList, line)
	}

	// 構造を作成
	// var repoModelStructList []*RepoModelStruct
	// var fieldList []*Field
	mode := "normal"
	for _, line := range lintList {

		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "import") {
			continue
		}

		var wardList []string
		for _, ward := range strings.Split(line, " ") {

			if ward == "" {
				continue
			}

			ward = strings.TrimSpace(ward)
			ward = strings.ReplaceAll(ward, "\t", "")
			wardList = append(wardList, ward)
		}

		if mode == "normal" {
			if wardList[0] == "type" {
				if strings.HasPrefix(wardList[0], "Update") {
					mode = "in_generated_struct"
					continue
				} else {
					mode = "in_struct"
				}
			}
		} else if mode == "in_struct" {

			if wardList[0] == "}" {
				mode = "normal"
				continue
			}

		} else if mode == "in_generated_struct" {
			if wardList[0] == "}" {
				mode = "normal"
			}
		}

		testtool.Log("wardList is ", wardList)
		// if strings.Split()

	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
