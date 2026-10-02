package reader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"web/internal/shared"
)

func validateItems(items []shared.Item) {
	for _, item := range items {
		if item.Name == "" || item.URL == "" {
			shared.CheckError(fmt.Errorf("invalid syntax in file: %+v", item))
		}
	}
}

func Read() []shared.Item {
	home, err := os.UserHomeDir()
	shared.CheckError(err)

	root := filepath.Join(home, ".config", "web")
	var dataset []shared.Item

	err = filepath.WalkDir(root, func(path string, _ fs.DirEntry, err error) error {
		shared.CheckError(err)
		ext := filepath.Ext(path)
		if ext == ".yaml" || ext == ".yml" {
			readFile(path, &dataset)
		}
		return nil
	})
	shared.CheckError(err)
	validateItems(dataset)
	return dataset
}

func readFile(file string, dataset *[]shared.Item) {
	data, err := os.ReadFile(file)
	shared.CheckError(err)

	var config map[string][]shared.Item
	err = yaml.Unmarshal(data, &config)
	shared.CheckError(err)

	for category, items := range config {
		for _, item := range items {
			item.Category = category
			*dataset = append(*dataset, item)
		}
	}
}
