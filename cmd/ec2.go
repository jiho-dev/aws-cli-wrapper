package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var ec2Cmds = []string{
	"describe-instances",
	"describe-network-interfaces",
}

var ec2CmdOpts = map[string][]string{
	"describe-instances": []string{
		"instance-ids",
	},

	"describe-network-interfaces": []string{
		"network-interface-ids",
	},

	"describe-nat-gateways": []string{
		"nat-gateway-ids",
	},

	"": []string{
		"",
	},
}

func newEc2Cmd() *cobra.Command {
	ec2Cmd := &cobra.Command{
		Use: "ec2",
		Run: ec2Main,
	}

	for _, c := range ec2Cmds {
		// cmd
		c := c
		cmd := &cobra.Command{
			Use: c,
			Run: ec2Main,
		}

		opts, ok := ec2CmdOpts[c]
		if ok && len(opts) > 0 {
			for _, o := range opts {
				cmd.Flags().String(o, "", "")
			}
		}

		ec2Cmd.AddCommand(cmd)
	}

	return ec2Cmd
}

func ec2Main(cobraCmd *cobra.Command, args []string) {
	if cobraCmd.Use == "ec2" {
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

	RunCmd(inCmds, false, flags)
}
