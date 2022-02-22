package cmd

import (
	"github.com/spf13/cobra"
)

func newEc2Cmd() *cobra.Command {
	return &cobra.Command{
		Use: "ec2",
		//Short: "ec2",
		Run: adminVpcMain,
	}

}

func init() {

}
