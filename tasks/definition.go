package tasks

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Definition struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"`
	AbortOnFail bool              `yaml:"abortOnFail"`
	Skip        bool              `yaml:"skip"`
	Args        map[string]string `yaml:"args"`
}

func (d Definition) Description() string {
	if len(d.Name) > 0 {
		return d.Name
	}
	return d.Type
}

func LoadFromData(rawData []byte, out *[]Definition) error {
	return yaml.Unmarshal(rawData, out)
}

func LoadFromFile(filePath string, out *[]Definition) error {
	rawTaskDefinitions, err := os.ReadFile(filePath)

	if err != nil {
		return err
	}
	return LoadFromData(rawTaskDefinitions, out)
}
