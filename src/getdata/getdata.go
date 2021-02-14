package getdata

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ninetypercentlanguage/misc/files"
	"github.com/ninetypercentlanguage/word-utils/combined"
)

type WordData struct {
	Word    string
	Content combined.Content
}

func GetWordsData(dirPath string) []WordData {
	var paths []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	var data []WordData
	for _, path := range paths {
		splitBySlash := strings.Split(path, "/")
		fileName := splitBySlash[len(splitBySlash)-1]

		data = append(
			data,
			WordData{
				Word:    strings.Split(fileName, ".")[0],
				Content: getSingleWordData(path),
			},
		)
	}
	return data
}

func getSingleWordData(filePath string) combined.Content {
	content := combined.Content{}
	_, err := files.GetJSONWhenFileMayNotExist(filePath, &content)
	if err != nil {
		panic(fmt.Sprintf("could not get json for path %v: %v", filePath, err))
	}
	return content
}
