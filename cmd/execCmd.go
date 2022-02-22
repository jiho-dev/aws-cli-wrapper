package cmd

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
	"github.com/tidwall/gjson"
)

func RunCmd(inCmds []string, adminVpc bool, flags *flag.FlagSet) {
	profile := inCmds[2]
	cmd := inCmds[0]
	opts, _ := adminVpcCmdOpts[cmd]

	var cmdOpt []string

	cmdOpt = append(cmdOpt, "ec2")
	cmdOpt = append(cmdOpt, "--profile")
	cmdOpt = append(cmdOpt, profile)
	if adminVpc {
		cmdOpt = append(cmdOpt, "admin-vpc")
		cmdOpt = append(cmdOpt, "--admin-action")
	}

	cmdOpt = append(cmdOpt, cmd)

	for i, o := range opts {
		if v, err := flags.GetString(o); v != "" && err == nil {
			if i == 0 {
				cmdOpt = append(cmdOpt, "--parameters")
			}

			cmdOpt = append(cmdOpt, fmt.Sprintf("Name=%s,Values=%v", o, v))
		}
	}

	output, err := ExecuteAwsCli("aws", cmdOpt...)

	if err != nil {
		fmt.Printf("ERR: %s \n", err)
		fmt.Printf("%s\n", output)
		return
	}

	if adminVpc {
		output = ParseOutput(output)
	}

	output = FormatJson(output)

	fmt.Printf("%s\n", output)
}

func ExecuteAwsCli(name string, args ...string) (string, error) {
	s := name
	if len(args) > 0 {
		for _, a := range args {
			s += " " + a
		}
	}

	fmt.Println(">", s)
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()

	o := string(out)
	if err != nil {
		fmt.Println("> error:", err.Error())
		fmt.Printf("> output: %s\n", o)
		return "", err
	}

	return o, nil
}

func ParseOutput(output string) string {
	// Result
	//value := gjson.Get(output, "Result")
	//value := gjson.Parse(output)

	value := gjson.Get(output, "*")

	return value.String()

	/*
		value = gjson.Get(value.Str, "NetworkInterfaces")

		if value.Str != "" {
			FormatJson(value.String())
		} else if value.IsArray() {
			value.ForEach(
				func(k, v gjson.Result) bool {
					ss := v.String()
					FormatJson(ss)
					return true
				})
		}
	*/
}

func FormatJson(output string) string {
	var obj map[string]interface{}
	json.Unmarshal([]byte(output), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.KeyColor = color.New(color.FgBlue)
	f.Indent = 3

	// Marshall the Colorized JSON
	b, _ := f.Marshal(obj)

	return string(b)
	//fmt.Println(string(b))
}
