package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vaughan0/go-ini"
	"gopkg.in/yaml.v2"
)

var profileCmds []*cobra.Command

type CmdConfig struct {
}

var yamlConfig = `
"admin-vpc":
   - "list-network-interface": ["network-interface-id" , "host-ip"]
   - "list-blackpearl":
`

type ConfItem struct {
	Test []string `yaml:"array.test,flow"`
}

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

func init() {

	m := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(yamlConfig), &m)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		return
	}

	fmt.Printf("--- m:\n%v\n\n", m)

	values := make(map[string][]string)
	for k, v := range m {
		fmt.Printf("k: %s, v:%T, %+v\n", k, v, v)

		if v == nil {
			continue
		}

		val, _ := v.([]interface{})
		for _, v2 := range val {
			fmt.Printf("  v2: %+v \n", v2)
			vv, _ := v2.(map[interface{}]interface{})
			for k4, v4 := range vv {
				fmt.Printf("    v4: %s=%T,%+v \n", k4, v4, v4)

				k5, _ := k4.(string)
				if v4 == nil {
					values[k5] = nil
					continue
				}

				//v5, _ := v4.([]string)
				var vvv []string
				for _, v5 := range v4.([]interface{}) {
					fmt.Printf("      v5: %T,%+v \n", v5, v5)
					vvv = append(vvv, v5.(string))
				}

				values[k5] = vvv
			}
		}
	}

	fmt.Printf("### values: %+v \n", values)

	os.Exit(1)

	////////////////////

	var profile []string
	profile = listProfiles()

	for _, p := range profile {
		cmd := &cobra.Command{
			Use: p,
			//Short: fmt.Sprintf("profile idx: %d", i),
			Run: profileMain,
		}

		// admin-vpc
		c := newAdminVcpCmd()
		cmd.AddCommand(c)

		// ec2
		c = newEc2Cmd()
		cmd.AddCommand(c)

		rootCmd.AddCommand(cmd)

		profileCmds = append(profileCmds, cmd)
	}
}

func profileMain(cmd *cobra.Command, args []string) {
	cmd.Help()
	os.Exit(0)
}

func listProfiles() []string {
	// Make sure the config file exists
	config := os.Getenv("HOME") + "/.aws/config"

	if _, err := os.Stat(config); os.IsNotExist(err) {
		fmt.Println("No credentials file found at: %s", config)
		os.Exit(1)
	}

	file, _ := ini.LoadFile(config)
	profiles := make([]string, 0)

	for key, _ := range file {
		if key == "default" {
			profiles = append(profiles, key)
		} else if strings.HasPrefix(key, "profile") {
			k := strings.Split(key, " ")
			if len(k) >= 2 {
				profiles = append(profiles, k[1])
			}
		}
	}

	sort.Strings(profiles)

	return profiles
}
