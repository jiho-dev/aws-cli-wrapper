package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spc",
	Short: "spc <profile> <sub-cmd> [flags}",
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	Hidden:                true,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(CompletionCmd)
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

	return nil

}

func Help(cmd *cobra.Command, s []string) {

	fmt.Printf("%s: warpper of aws cli for SPC\n\n", cmd.Use)

	fmt.Printf("Usage: %s <profile> <sub-cmd> [flags] \n", cmd.Use)
	fmt.Printf("  <profile>: profile name in aws config\n")
	fmt.Printf("  <sub-cmd>: \n")
	fmt.Printf("    - admin-vpc \n")
	fmt.Printf("    - ec2 \n")
}
