package profile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const ext = ".yaml"

type Profile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

func path(name string) string {
	return name + ext
}

func Create(name, user, project string) error {
	if user == "" {
		return errors.New("user cannot be empty")
	}
	if project == "" {
		return errors.New("project cannot be empty")
	}
	fpath := path(name)

	f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("profile %s already exists", name)
		}
		return err
	}
	defer f.Close()

	p := Profile{User: user, Project: project}
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

func Get(name string) (*Profile, error) {
	data, err := os.ReadFile(path(name))
	if err != nil {
		return nil, fmt.Errorf("profile %s not found", name)
	}
	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func Delete(name string) error {
	err := os.Remove(path(name))
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("profile %s not found", name)
	}
	return err
}

func List() ([]string, error) {
	files, err := filepath.Glob("*" + ext)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(files))
	for _, f := range files {
		names = append(names, strings.TrimSuffix(f, ext))
	}
	return names, nil
}
