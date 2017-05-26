package goconf

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func Parse(v interface{}, options ...Option) error {
	var err error
	o := &opts{}
	for _, option := range options {
		option(o)
	}

	if o.yaml {
		if err = parseYaml(v, o); err != nil {
			return err
		}
	}

	if o.env {
		if err = envconfig.Process(o.envPrefix, v); err != nil {
			return err
		}
	}

	return nil
}

func ReloadYaml(path string, v interface{}) error {
	yamlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlBytes, v)
}

func ReloadYamlFromBytes(bytes []byte, v interface{}) error {
	return yaml.Unmarshal(bytes, v)
}

func parseYaml(v interface{}, o *opts) error {
	var err error
	if o.yamlBytes != nil {
		err = yaml.Unmarshal(o.yamlBytes, v)
	}

	if err == nil && o.yamlPath != "" {
		yamlBytes, err := ioutil.ReadFile(o.yamlPath)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(yamlBytes, v)
	}

	return err
}
