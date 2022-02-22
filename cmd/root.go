package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spc",
	Short: "spc <profile> <sub-cmd> [flags}",
}

// Execute executes cmd
func Execute() error {

	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableNoDescFlag:   true,
		HiddenDefaultCmd:    true,
		DisableDescriptions: true,
	}

	//rootCmd.SetHelpFunc(Help)
	return rootCmd.Execute()

}

func Help(cmd *cobra.Command, s []string) {

	fmt.Printf("%s: warpper of aws cli for SPC\n\n", cmd.Use)

	fmt.Printf("Usage: %s <profile> <sub-cmd> [flags] \n", cmd.Use)
	fmt.Printf("  <profile>: profile name in aws config\n")
	fmt.Printf("  <sub-cmd>: \n")
	fmt.Printf("    - admin-vpc \n")
	fmt.Printf("    - ec2 \n")
}
