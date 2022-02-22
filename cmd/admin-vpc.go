package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var adminVpcCmds = []string{
	"list-network-interface",
	"list-blackpearl",
	"list-security-group",
	"list-network-acl",
	"list-public-ips",
	"list-public-ipv4-pool",
	"list-address-associations",
	"list-route-table",
	"list-vrouters",
}

var adminVpcCmdOpts = map[string][]string{
	"list-network-interface": []string{
		"host-ip",
		"network-interface-id",
		"vpc-id",
		"subnet-id",
		"private-ip-address",
		"owner-id",
		"mac-address",
		"nat-ip",
		"instance-id",
	},

	"list-blackpearl": []string{
		"instance-id",
		"blackpearl-ip",
		"host-ip",
		"network-interface-id",
	},

	"list-security-group": []string{
		"vpc-id", "group-id",
	},

	"list-network-acl": []string{
		"network-acl-id", "vpc-id",
	},

	"list-route-table": []string{
		"route-table-id", "vpc-id",
	},
}

func newAdminVcpCmd() *cobra.Command {
	adminVpcRootCmd := &cobra.Command{
		Use: "admin-vpc",
		//Short: "admin-vpc",
		Run: adminVpcMain,
	}

	for _, c := range adminVpcCmds {
		// cmd
		c := c
		cmd := &cobra.Command{
			Use: c,
			//Short: fmt.Sprintf("admin-vpc %s", c),
			Run: adminVpcMain,
		}

		opts, ok := adminVpcCmdOpts[c]
		if ok && len(opts) > 0 {
			for _, o := range opts {
				cmd.Flags().String(o, "", "")

				//if o == "network-interface-id" {
				//	cmd.MarkFlagRequired(o)
				//}
			}
		}

		adminVpcRootCmd.AddCommand(cmd)
	}

	return adminVpcRootCmd
}

func adminVpcMain(cobraCmd *cobra.Command, args []string) {
	if cobraCmd.Use == "admin-vpc" {
		cobraCmd.Help()
		os.Exit(0)
	}

	var inCmds []string

	inCmds = append(inCmds, cobraCmd.Use)

	var c1 = cobraCmd
	for c1.HasParent() {
		c1 = c1.Parent()
		inCmds = append(inCmds, c1.Use)
	}

	profile := inCmds[2]
	cmd := inCmds[0]
	opts, _ := adminVpcCmdOpts[cmd]

	var cmdOpt []string

	cmdOpt = append(cmdOpt, "ec2")
	cmdOpt = append(cmdOpt, "--profile")
	cmdOpt = append(cmdOpt, profile)
	cmdOpt = append(cmdOpt, "admin-vpc")
	cmdOpt = append(cmdOpt, "--admin-action")
	cmdOpt = append(cmdOpt, cmd)

	flags := cobraCmd.Flags()
	for i, o := range opts {
		if v, err := flags.GetString(o); v != "" && err == nil {
			if i == 0 {
				cmdOpt = append(cmdOpt, "--parameters")
			}

			cmdOpt = append(cmdOpt, fmt.Sprintf("Name=%s,Values=%v", o, v))
		}
	}

	output, err := ExecuteAwsCli("aws", cmdOpt...)

	if err != nil {
		fmt.Printf("ERR: %s \n", err)
	} else {
		fmt.Printf("%s\n", output)
	}
}
