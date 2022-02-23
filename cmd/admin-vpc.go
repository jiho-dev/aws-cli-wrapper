package cmd

import (
	"os"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

var adminVpcCmds map[string][]string

func newAdminVcpCmd(conf config.Config) *cobra.Command {
	adminVpcRootCmd := &cobra.Command{
		Use: TYPE_ADMIN_VPC,
		//Short: "admin-vpc",
		Run: adminVpcMain,
	}

	adminVpcCmds, _ = conf[TYPE_ADMIN_VPC]
	for c, opts := range adminVpcCmds {
		// cmd
		c := c
		cmd := &cobra.Command{
			Use: c,
			//Short: fmt.Sprintf("admin-vpc %s", c),
			Run: adminVpcMain,
		}

		if len(opts) > 0 {
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
	if cobraCmd.Use == TYPE_ADMIN_VPC {
		cobraCmd.Help()
		os.Exit(0)
	}

	var inCmds []string
	flags := cobraCmd.Flags()

	inCmds = append(inCmds, cobraCmd.Use)

	var c1 = cobraCmd
	for c1.HasParent() {
		c1 = c1.Parent()
		inCmds = append(inCmds, c1.Use)
	}

	opts, _ := adminVpcCmds[cobraCmd.Use]
	RunCmd(inCmds, opts, true, flags)
}
