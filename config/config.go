package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigItem map[string][]string
type Config map[string]ConfigItem

type ConfigApiItem struct {
	Api         string   `yaml:api`
	OutputField string   `yaml:outputfield`
	Opts        []string `yaml:opts`
}

type ConfigApis map[string]ConfigItem

type ConfigApi struct {
	//ApiType string          `yaml:apitype`
	Apis []ConfigApiItem `yaml:apis`
}

type ConfigApiList struct {
	//ApiList []ConfigApi `yaml:apilist`
	ApiList []ConfigApi
}

func (api *ConfigApi) GetOpts(apiName string) *ConfigApiItem {

	/*
		for _, api := range api.Apis {
			if api.Api == apiName {
				return &api
			}
		}
	*/

	return nil
}

func (c *ConfigApiList) GetApiList(t string) *ConfigApi {
	/*
		for _, a := range c.ApiList {
			if a.ApiType == t {
				return &a
			}
		}
	*/

	return nil
}

func ParseConfigApi(fileName string) (*ConfigApiList, error) {
	fileName, _ = filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var cc ConfigApiList
	err = yaml.Unmarshal(yamlFile, &cc)
	if err != nil {
		return nil, err
	}

	return &cc, nil
}

func YamlTest() {
	//c := ConfigApiList{
	//ApiList: []ConfigApi{
	// c := []ConfigApi{
	//c := map[string]ConfigApi{
	c := map[string][]ConfigApiItem{

		//"admin-vpc": //ConfigApi{
		//ApiType: "admin-vpc",
		//Apis: []ConfigApiItem{
		"admin-vpc": []ConfigApiItem{
			ConfigApiItem{
				Api:         "list-network-interface",
				OutputField: "NetworkInterfaces",
				Opts:        []string{"host-ip", "network-interface-id", "vpc-id"},
			},
			ConfigApiItem{
				Api:         "list-address-associations",
				OutputField: "NetworkInterfaces",
			},
			ConfigApiItem{
				Api:         "list-public-ips",
				OutputField: "NetworkInterfaces",
			},
		},
		//},

		//"ec2": ConfigApi{
		//ApiType: "ec2",
		//Apis: []ConfigApiItem{
		"ec2": []ConfigApiItem{
			ConfigApiItem{
				Api:         "describe-instances",
				OutputField: "NetworkInterfaces",
				Opts:        []string{"instance-ids"},
			},
			ConfigApiItem{
				Api:         "describe-network-interfaces",
				OutputField: "NetworkInterfaces",
				Opts:        []string{"network-interface-ids"},
			},

			ConfigApiItem{
				Api:         "get-console-output",
				OutputField: "Output",
				Opts:        []string{"instance-id"},
			},
		},
		//},
		//},
	}

	yamlData, err := yaml.Marshal(&c)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(" --- YAML ---")
	fmt.Printf("%s \n", string(yamlData))
}

func ParseConfig(fileName string) (Config, error) {
	fileName, _ = filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	conf := make(Config)
	for k1, v1 := range m {
		if v1 == nil {
			continue
		}

		//values := make(map[string][]string)
		item := make(map[string][]string)

		_v2, _ := v1.([]interface{})
		for _, v2 := range _v2 {
			_v3, _ := v2.(map[interface{}]interface{})
			for k3, v3 := range _v3 {
				k4, _ := k3.(string)
				if v3 == nil {
					item[k4] = nil
					continue
				}

				_v4, _ := v3.([]interface{})
				var val []string
				for _, v4 := range _v4 {
					v5 := v4.(string)
					if v5 != "" {
						val = append(val, v5)
					}
				}

				item[k4] = val
			}
		}

		conf[k1] = item
	}

	return conf, nil
}
