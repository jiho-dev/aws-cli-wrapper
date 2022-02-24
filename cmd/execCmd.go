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

func RunCmd(inCmds []string, apiArgs []string, adminVpc bool, flags *flag.FlagSet) (string, error) {
	profile := inCmds[2]
	cmd := inCmds[0]

	var cmdOpt []string

	cmdOpt = append(cmdOpt, "ec2")
	cmdOpt = append(cmdOpt, "--profile")
	cmdOpt = append(cmdOpt, profile)

	if adminVpc {
		cmdOpt = append(cmdOpt, "admin-vpc")
		cmdOpt = append(cmdOpt, "--admin-action")
	}

	cmdOpt = append(cmdOpt, cmd)

	var optCnt int
	for _, o := range apiArgs {
		if v, err := flags.GetString(o); v != "" && err == nil {
			if adminVpc {
				if optCnt == 0 {
					cmdOpt = append(cmdOpt, "--parameters")
				}

				cmdOpt = append(cmdOpt, fmt.Sprintf("Name=%s,Values=%v", o, v))
				optCnt++
			} else {
				cmdOpt = append(cmdOpt, fmt.Sprintf("--%s", o))
				cmdOpt = append(cmdOpt, v)
				optCnt += 2
			}
		}
	}

	return ExecuteAwsCli("aws", cmdOpt...)

	/*
		fmt.Printf(">> %s\n", output)
		query, err := gojq.Parse(".foo | ..")
	*/

	/*

		if adminVpc {
			output = ParseOutput(output)
		}

		output = FormatJson(output)
		output = strings.Replace(output, "\\r\\n", "\r\n", -1)

		fmt.Printf(">>> %s\n", output)
	*/
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

func ParseOutput(output string, outField string) string {
	// Result
	//value := gjson.Get(output, "Result")
	//value := gjson.Parse(output)

	value := gjson.Get(output, outField)
	//value := gjson.Get(output, "*")

	output = value.String()

	return output

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
	f.RawStrings = true

	// Marshall the Colorized JSON
	b, _ := f.Marshal(obj)

	return string(b)
}
