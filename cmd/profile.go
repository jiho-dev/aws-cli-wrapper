package cmd

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
	"github.com/vaughan0/go-ini"
)

var awsDir = os.Getenv("HOME") + "/.aws/"
var AcwConfig *config.AcwConfig

var CompOpt = cobra.CompletionOptions{
	DisableDefaultCmd:   true,
	DisableNoDescFlag:   true,
	HiddenDefaultCmd:    true,
	DisableDescriptions: true,
}

func init() {
	var profile []string
	profile = listProfiles()

	confFile := path.Join(awsDir, "acw.yaml")
	conf, err := config.ParseConfig(confFile)
	if err != nil {
		//fmt.Printf("ERR: %s\n", err)
		return
	}

	AcwConfig = conf

	for _, p := range profile {
		cmd := &cobra.Command{
			Use:               p,
			Run:               profileMain,
			CompletionOptions: CompOpt,
		}

		for apiGroup, apis := range conf.ApiGroup {
			subCmd := InitApiGroupCmd(apiGroup, apis)
			cmd.AddCommand(subCmd)
		}

		rootCmd.AddCommand(cmd)
	}
}

func profileMain(cmd *cobra.Command, args []string) {
	cmd.Help()
	os.Exit(0)
}

func listProfiles() []string {
	// Make sure the config file exists
	config := path.Join(awsDir, "config")

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
