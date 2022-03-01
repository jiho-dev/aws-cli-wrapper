package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

var awsDir = os.Getenv("HOME") + "/.aws/"
var acwConf = "acw.yaml"
var AcwConfig *config.AcwConfig

var apiGroups = []string{CMD_ADMIN_VPC, CMD_EC2}

var CompOpt = cobra.CompletionOptions{
	DisableDefaultCmd:   true,
	DisableNoDescFlag:   true,
	HiddenDefaultCmd:    true,
	DisableDescriptions: true,
}

/////////////////////////////////

var rootCmd = &cobra.Command{
	Use:               "acw",
	Short:             "acw <api-group> <sub-cmd> [flags]",
	Long:              "aws-cli-wrapper to support shell completion for some command",
	CompletionOptions: CompOpt,
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
			cmd.Root().GenBashCompletionV2(os.Stdout, false)
		case "zsh":
			cmd.Root().GenZshCompletionNoDesc(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, false)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	confFile := path.Join(awsDir, acwConf)
	conf, err := config.ParseConfig(confFile)
	if err != nil {
		return
	}

	AcwConfig = conf

	for apiGroup, apis := range conf.ApiGroup {
		subCmd := InitApiGroupCmd(apiGroup, apis)

		rootCmd.AddCommand(subCmd)
	}

	rootCmd.AddCommand(CompletionCmd)
}

// Execute executes cmd
func Execute() error {

	//rootCmd.SetHelpFunc(Help)
	return rootCmd.Execute()

	return nil

}

func Help(cmd *cobra.Command, s []string) {

	fmt.Printf("%s: warpper of aws cli for SPC\n\n", cmd.Use)

	fmt.Printf("Usage: %s <api-group> <sub-cmd> [flags] \n", cmd.Use)
	fmt.Printf("  <api-group>: group of api, admin-vpc or ec2 \n")
	fmt.Printf("  <sub-cmd>: \n")
}
