package cmd

import (
	"os"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

var ec2Cmds map[string][]string

func newEc2Cmd(conf config.Config) *cobra.Command {
	ec2Cmd := &cobra.Command{
		Use: TYPE_EC2,
		Run: ec2Main,
	}

	ec2Cmds, _ = conf[TYPE_EC2]
	for c, opts := range ec2Cmds {
		// cmd
		c := c
		cmd := &cobra.Command{
			Use: c,
			Run: ec2Main,
		}

		if len(opts) > 0 {
			for _, o := range opts {
				cmd.Flags().String(o, "", "")
			}
		}

		ec2Cmd.AddCommand(cmd)
	}

	return ec2Cmd
}

func ec2Main(cobraCmd *cobra.Command, args []string) {
	if cobraCmd.Use == TYPE_EC2 {
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

	opts, _ := ec2Cmds[cobraCmd.Use]
	RunCmd(inCmds, opts, false, flags)
}
