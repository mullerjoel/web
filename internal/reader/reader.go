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

	pattern := filepath.Join(home, ".config", "web", "*.yaml")
	files, err := filepath.Glob(pattern)
	shared.CheckError(err)

	var dataset []shared.Item

	for _, file := range files {
		readFile(file, &dataset)
	}
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
