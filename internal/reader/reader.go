package reader

import (
	"os"
	"path/filepath"
	"web/internal/check"

	"gopkg.in/yaml.v3"
)

type Item struct {
	Category string
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Desc     string `yaml:"desc,omitempty"`
	Git      string `yaml:"git,omitempty"`
}

func Read() []Item {
	home, err := os.UserHomeDir()
	check.CheckError(err)

	pattern := filepath.Join(home, ".config", "web", "*.yaml")
	files, err := filepath.Glob(pattern)
	check.CheckError(err)

	var dataset []Item

	for _, file := range files {
		data, err := os.ReadFile(file)
		check.CheckError(err)

		var config map[string][]Item
		err = yaml.Unmarshal(data, &config)
		check.CheckError(err)

		for category, items := range config {
			for _, item := range items {
				item.Category = category
				dataset = append(dataset, item)
			}
		}
	}

	return dataset
}
