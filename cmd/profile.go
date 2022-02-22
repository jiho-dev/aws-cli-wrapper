package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vaughan0/go-ini"
)

var profileCmds []*cobra.Command

func init() {
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
