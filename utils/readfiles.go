package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Envs struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Task struct {
	Name    string   `yaml:"name"`
	Node    []string `yaml:"node"`
	Command string   `yaml:"command"`
}
type CDRunfile struct {
	Name      string `yaml:"cdrunfile"`
	GitUrl    string `yaml:"git_url"`
	LocalPath string `yaml:"localpath"`
	On        string `yaml:"on"`
	Envs      []Envs `yaml:"envs"`
	Tasks     []Task `yaml:"tasks"`
}

func CDReadRunfile(path string) (*CDRunfile, error) {

	var runfile CDRunfile

	f, err := os.ReadFile(path)
	if err != nil {
		return &runfile, fmt.Errorf("unable to read cdrunfile: %v", err)
	}

	err = yaml.Unmarshal(f, &runfile)
	if err != nil {
		return &runfile, fmt.Errorf("unable to parse cdrunfile, please verify file: %v", err)
	}

	return &runfile, nil

}
