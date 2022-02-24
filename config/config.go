package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

/////////////////////

type AcwConfigApiOpt struct {
	OutputField string   `yaml:outputfield`
	Required    []string `yaml:required`
	Args        []string `yaml:args`
}

// Key: api name
type AcwConfigApis map[string]AcwConfigApiOpt

// key: api group: [ ec2 | admin-vpc ]
type AcwConfigApiGroup map[string]AcwConfigApis

type AcwConfig struct {
	Version  string
	ApiGroup AcwConfigApiGroup
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

func YamlTest() {
	c := AcwConfig{
		Version: "1",
		ApiGroup: AcwConfigApiGroup{
			"admin-vpc": AcwConfigApis{
				"list-network-interface": AcwConfigApiOpt{
					OutputField: "NetworkInterfaces",
					Args:        []string{"host-ip", "network-interface-id", "vpc-id"},
				},

				"list-address-associations": AcwConfigApiOpt{
					OutputField: "NetworkInterfaces",
				},
				"list-public-ips": AcwConfigApiOpt{
					OutputField: "NetworkInterfaces",
				},
			},

			"ec2": AcwConfigApis{
				"describe-instances": AcwConfigApiOpt{
					OutputField: "NetworkInterfaces",
					Args:        []string{"instance-ids"},
				},
				"describe-network-interfaces": AcwConfigApiOpt{
					OutputField: "NetworkInterfaces",
					Args:        []string{"network-interface-ids"},
				},

				"get-console-output": AcwConfigApiOpt{
					OutputField: "Output",
					Required:    []string{"instance-id"},
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(&c)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(" --- YAML ---")
	fmt.Printf("%s \n", string(yamlData))
}
