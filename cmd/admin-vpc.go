package cmd

import (
	"fmt"
	"os"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

//var adminVpcCmds map[string][]string
var adminVpcCmds *config.ConfigApi

func newAdminVcpCmd(conf *config.ConfigApiList) *cobra.Command {
	adminVpcRootCmd := &cobra.Command{
		Use: TYPE_ADMIN_VPC,
		//Short: "admin-vpc",
		Run: adminVpcMain,
	}

	//adminVpcCmds, _ = conf[TYPE_ADMIN_VPC]
	adminVpcCmds = conf.GetApiList(TYPE_ADMIN_VPC)
	for _, api := range adminVpcCmds.Apis {
		// cmd
		cmd := &cobra.Command{
			Use: api.Api,
			//Short: fmt.Sprintf("admin-vpc %s", c),
			Run: adminVpcMain,
		}

		if len(api.Opts) > 0 {
			for _, o := range api.Opts {
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

	//opts, _ := adminVpcCmds[cobraCmd.Use]
	var opts []string
	api := adminVpcCmds.GetOpts(cobraCmd.Use)
	if api != nil {
		opts = api.Opts
	}

	output, err := RunCmd(inCmds, opts, true, flags)
	if err != nil {
		if output != "" {
			fmt.Printf("Output: %s \n", output)
		}

		fmt.Printf("ERR: %s \n", err)
		return
	}

	if output == "" {
		fmt.Printf("No Output")
		return
	}

	output1 := ParseOutput(output, api.OutputField)
	if output1 == "" {
		output1 = output
	}

	output2 := FormatJson(output1)
	if output2 == "" {
		output2 = output1
	}

	//output = strings.Replace(output, "\\r\\n", "\r\n", -1)

	fmt.Printf("%s\n", output2)
}
