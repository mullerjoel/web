package reader

import (
	"fmt"
	"gopkg.in/yaml.v3"
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

	var dataset []shared.Item
	walk(filepath.Join(home, ".config", "web"), &dataset)
	validateItems(dataset)
	return dataset
}

func walk(dir string, dataset *[]shared.Item) {
	entries, err := os.ReadDir(dir)
	shared.CheckError(err)

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		info, err := os.Stat(path)
		if err != nil {
			continue
		}
		if info.IsDir() {
			walk(path, dataset)
		} else if ext := filepath.Ext(path); ext == ".yaml" || ext == ".yml" {
			readFile(path, dataset)
		}
	}
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
