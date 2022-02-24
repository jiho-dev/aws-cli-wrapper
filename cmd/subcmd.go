package cmd

import (
	"fmt"
	"os"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

func InitApiGroupCmd(apiGroup string, apis config.AcwConfigApis) *cobra.Command {
	cmd := &cobra.Command{
		Use: apiGroup,
		Run: apiGroupMain,
	}

	for apiName, opt := range apis {
		subCmd := &cobra.Command{
			Use: apiName,
			Run: cmd.Run,
		}

		for _, o := range opt.Required {
			subCmd.Flags().String(o, "", "")
			subCmd.MarkFlagRequired(o)
		}

		for _, o := range opt.Args {
			subCmd.Flags().String(o, "", "")
		}

		cmd.AddCommand(subCmd)
	}

	return cmd
}

func apiGroupMain(cobraCmd *cobra.Command, args []string) {
	var depth int
	var inCmds []string

	flags := cobraCmd.Flags()
	inCmds = append(inCmds, cobraCmd.Use)

	var c1 = cobraCmd
	for c1.HasParent() {
		c1 = c1.Parent()
		inCmds = append(inCmds, c1.Use)
		depth++
	}

	// acw <profile> <api-group> <cmd> [args]
	//                             ^
	// 0   1         2           3
	if depth < 3 {
		cobraCmd.Help()
		os.Exit(0)
	}

	parent := cobraCmd.Parent()
	apis, ok := AcwConfig.ApiGroup[parent.Use]

	var apiArgs []string
	opts, ok := apis[cobraCmd.Use]
	if ok {
		apiArgs = append(apiArgs, opts.Args...)
		apiArgs = append(apiArgs, opts.Required...)
	}

	isAdminVpc := parent.Use == TYPE_ADMIN_VPC

	output, err := RunCmd(inCmds, apiArgs, isAdminVpc, flags)
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

	output1 := ParseOutput(output, opts.OutputField)
	if output1 == "" {
		output1 = output
	}

	output2 := FormatJson(output1)
	if output2 == "" || output2 == "{}" {
		output2 = output1
	}

	//output = strings.Replace(output, "\\r\\n", "\r\n", -1)

	fmt.Printf("%s\n", output2)
}
