package cmd

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

var adminVpcCmds = []string{
	"allocate-random-ip-pool",
	"blackpearl-health",
	"create-public-ipv4-pool",
	"delete-public-ipv4-pool",
	"deregister-public-ipv4-pool",
	"disable-public-ipv4-pool",
	"disassociate-public-ip",
	"enable-public-ipv4-pool",
	"list-address-associations",
	"list-blackpearl",
	"list-network-acl",
	"list-network-interface",
	"list-public-ips",
	"list-public-ipv4-pool",
	"list-route-table",
	"list-security-group",
	"list-vrouters",
	"register-public-ipv4-pool",
	"release-ip-pool",
	"release-public-ip",
	"request-ip-pool",
	"show-dataversion",
	"show-flowlog",
	"show-network-interface",
	"show-papyrus-flowlog",
	"show-papyrus-summary",
	"show-revision",
	"show-snat",
	"show-summary",
	"show-vrevision",
	"show-vrouter-flowlog",
	"show-vrouter-flow",
	"show-vrouter-network-acl",
	"show-vrouter-network-interface",
	"show-vrouter-port",
	"show-vrouter-route",
	"show-vrouter-security-group",
	"show-vrouter-subnet",
	"show-vrouter-summary",
	"show-vrouter-table",
	"update-network-interface",
}

var genAdminVpcCmd = &cobra.Command{
	Use:   "gen-admin-vpc",
	Short: "Generate adminvpc commands",
	//Hidden:                true,
	//DisableFlagsInUseLine: true,
	//ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	//Args: cobra.ExactValidArgs(1),

	Run: genAdminVpcCmdMain,
}

func InitGenAdminVpcCmd() *cobra.Command {
	addProfileCmd(genAdminVpcCmd)

	return genAdminVpcCmd
}

func genAdminVpcCmdMain(cobraCmd *cobra.Command, args []string) {

	flags := cobraCmd.Flags()
	flags.Bool(SHOW_HELP, true, "")
	isAdminVpc := true
	adminVpcGroup, _ := AcwConfig.ApiGroup[ADMIN_VPC]

	for _, api := range adminVpcCmds {
		inCmds := []string{api}

		output, err := RunCmd(inCmds, nil, isAdminVpc, flags)
		if err != nil {
			if output != "" {
				fmt.Printf("Output: %s \n", output)
			}

			fmt.Printf("ERR: %s \n", err)
			continue
		}

		if output == "" {
			fmt.Printf("No Output\n")
			continue
		}

		output1 := ParseOutput(output, "Result")
		if output1 == "" {
			output1 = output
		}

		output2 := FormatJson(output1)
		if output2 == "" || output2 == "{}" {
			output2 = output1
		}

		fmt.Printf("%s\n", output2)

		newOpts := config.AcwConfigApiOpt{}

		oldOpts, ok := adminVpcGroup[api]
		if ok {
			newOpts.OutputField = oldOpts.OutputField
			newOpts.Required = oldOpts.Required
		} else {
			newOpts.OutputField = "Result"
		}

		args := strings.Split(output2, "\n")
		var seeParams bool
		for _, arg := range args {
			if strings.HasPrefix(arg, "Parameters:") {
				seeParams = true
				continue
			}

			if seeParams {
				required := strings.Contains(arg, "(required)")

				arg = strings.TrimSpace(arg)
				tmp := strings.Split(arg, " ")
				key := tmp[0]
				key = strings.TrimSpace(key)

				if required {
					if !contains(newOpts.Required, key) {
						newOpts.Required = append(newOpts.Required, key)
					}
				} else if !contains(newOpts.Required, key) {
					newOpts.Args = append(newOpts.Args, key)
				}
			}
		}

		adminVpcGroup[api] = newOpts
	}

	confFile := path.Join(awsDir, acwConf)
	config.WriteConfig(AcwConfig, confFile)
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
