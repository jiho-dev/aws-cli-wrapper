package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

/////////////////////

type AcwConfigApiOpt struct {
	OutputField string   `yaml:"OutputField"`
	Required    []string `yaml:"Required"`
	Args        []string `yaml:"Args"`
}

// Key: api name
type AcwConfigApis map[string]AcwConfigApiOpt

const API_TERMINATED = "api-terminated"

type AcwConfig struct {
	Version         string                         `yaml:"Version"`
	ApiPrefixFilter []string                       `yaml:"ApiPrefixFilter"`
	ApiPrefix       map[string]map[string][]string `yaml:"ApiPrefix"`
	ApiOptions      map[string]AcwConfigApiOpt     `yaml:"ApiOptions"` // key: api, value: api-option
}

func ParseConfig(fileName string) (*AcwConfig, error) {
	fileName, _ = filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var cc AcwConfig
	err = yaml.Unmarshal(yamlFile, &cc)
	if err != nil {
		return nil, err
	}

	return &cc, nil
}

func WriteConfig(conf *AcwConfig, fileName string) error {
	yamlData, err := yaml.Marshal(conf)

	fileName, _ = filepath.Abs(fileName)
	err = ioutil.WriteFile(fileName, yamlData, 644)
	if err != nil {
		return err
	}

	return nil
}

func YamlTest() {
	c := AcwConfig{
		Version: "1",

		ApiPrefixFilter: []string{
			"describe-network",
			"describe-no",
		},

		ApiPrefix: map[string]map[string][]string{
			"describe": {
				"network": []string{"acls", "interfaces"},
				"volumes": []string{API_TERMINATED, "modifications"},
			},
			"import": {
				"image":    []string{API_TERMINATED},
				"key":      []string{"pair"},
				"snapshot": []string{API_TERMINATED},
			},
			"modify": {
				"hosts":    []string{API_TERMINATED},
				"instance": []string{"attribute", "capacity-reservation-attributes", "event-window"},
				"snapshot": []string{API_TERMINATED},
			},
			"wait": {
				API_TERMINATED: []string{},
			},
		},

		ApiOptions: map[string]AcwConfigApiOpt{
			"describe-network-interfaces": {
				OutputField: "NetworkInterfaces",
				Args:        []string{"host-ip", "network-interface-id", "vpc-id"},
			},

			"describe-network-acls": {
				OutputField: "NetworkInterfaces",
				Args:        []string{"host-ip", "nacl-id", "vpc-id"},
			},
		},
	}

	sort.Strings(c.ApiPrefixFilter)

	yamlData, err := yaml.Marshal(&c)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(" --- YAML ---")
	fmt.Printf("%s \n", string(yamlData))
}
