package cmd

import (
	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

//var ec2Cmds map[string][]string
var ec2Cmds config.AcwConfigApis

func newEc2Cmd(conf *config.AcwConfig) *cobra.Command {
	/*
		ec2Cmd := &cobra.Command{
			Use: TYPE_EC2,
			Run: ec2Main,
		}

		//ec2Cmds, _ = conf[TYPE_EC2]
		ec2Cmds = conf.GetApiList(TYPE_EC2)
		for _, api := range ec2Cmds.Apis {
			// cmd
			cmd := &cobra.Command{
				Use: api.Api,
				Run: ec2Main,
			}

			if len(api.Opts) > 0 {
				for _, o := range api.Opts {
					cmd.Flags().String(o, "", "")
				}
			}

			ec2Cmd.AddCommand(cmd)
		}

		return ec2Cmd
	*/
	return nil
}

func ec2Main(cobraCmd *cobra.Command, args []string) {
	/*
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

		//opts, _ := ec2Cmds[cobraCmd.Use]
		var opts []string
		api := ec2Cmds.GetOpts(cobraCmd.Use)
		if api != nil {
			opts = api.Opts
		}

		output, err := RunCmd(inCmds, opts, false, flags)
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
		if output2 == "" || output2 == "{}" {
			output2 = output1
		}

		//output = strings.Replace(output, "\\r\\n", "\r\n", -1)

		fmt.Printf("%s\n", output2)
	*/
}
