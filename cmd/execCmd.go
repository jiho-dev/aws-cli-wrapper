package cmd

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

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

	return ParseOutput(o), nil
}

func ParseOutput(output string) string {
	// Result
	//value := gjson.Get(output, "Result")
	//value := gjson.Parse(output)

	value := gjson.Get(output, "*")

	return FormatJson(value.String())

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
