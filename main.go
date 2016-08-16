package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"

	"mig.ninja/mig/modules"
	"mig.ninja/mig/modules/audit"
)

func main() {
	var (
		// provide serveraddress as an address of a server which accepts POST requests
		// audit.rules.json contains the rules as specified by libaudit-go which are to be set in the kernel (modify them as required)
		paramargs = `{"class":"parameters", "parameters":{"rulefilepath": "audit.rules.json", "serveraddress": "http://127.0.0.1:4443"}}`
		out       string
		pretty    = true
	)

	infd := bufio.NewReader(bytes.NewBuffer([]byte(paramargs)))
	// instantiate and call module
	// run := modules.Available["audit"].NewRun()
	moduler := audit.GetInstance()
	run := moduler.NewRun()
	out = run.Run(infd)
	if pretty {
		var modres modules.Result
		err := json.Unmarshal([]byte(out), &modres)
		if err != nil {
			panic(err)
		}
		out = ""
		if _, ok := run.(modules.HasResultsPrinter); ok {
			outRes, err := run.(modules.HasResultsPrinter).PrintResults(modres, false)
			if err != nil {
				panic(err)
			}
			for _, resLine := range outRes {
				out += fmt.Sprintf("%s\n", resLine)
			}
		} else {
			out = fmt.Sprintf("[error] no printer available for module '%s'\n", "audit")
		}
	}
	fmt.Println(out)
	return
}
