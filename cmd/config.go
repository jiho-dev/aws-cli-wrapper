package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config map[string]map[string][]string

/*
func getConfigItem(key string) map[string]string {
	values := map[string]string{}

	vals := viper.Get(key)
	val, ok := vals.([]interface{})
	if ok {
		for _, val2 := range val {
			val3, ok3 := val2.(map[interface{}]interface{})
			if ok3 {
				for key4, val4 := range val3 {
					k := key4.(string)
					v := val4.(string)
					values[k] = v
				}
			}
		}
	}

	return values
}
*/

var yamlConfig = `
"admin-vpc":
   - "list-network-interface": ["network-interface-id" , "host-ip"]
   - "list-blackpearl":

"ec2":
   - "describe-instances": ["instance-ids"]
   - "describe-network-interfaces": [ "network-interface-ids"]
   - "describe-nat-gateways": [ "nat-gateway-ids"]

`

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

		values := make(map[string][]string)

		_v2, _ := v1.([]interface{})
		for _, v2 := range _v2 {
			_v3, _ := v2.(map[interface{}]interface{})
			for k3, v3 := range _v3 {
				k4, _ := k3.(string)
				if v3 == nil {
					values[k4] = nil
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

				values[k4] = val
			}
		}

		conf[k1] = values
	}

	return conf, nil
}
