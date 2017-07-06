package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/JulzDiverse/aviator/aviator"

	"github.com/urfave/cli"
)

func main() {
	cmd := setCli()

	cmd.Action = func(c *cli.Context) error {

		var yml aviator.Aviator

		aviatorFile := c.String("file")
		if !VerifyAviatorFile(aviatorFile) {
			fmt.Println("No Aviator file found. Does the file exist?\n ")
			fmt.Println("Please navigate to a Aviator directory or specify a AVIATOR YAML ('.vtr')  with [--file|-f] option  and run aviator again")
			fmt.Println("NOTE: specified AVIATOR files require the '.vtr' extension")
			os.Exit(1)
		} else {
			ymlBytes, err := ioutil.ReadFile(aviatorFile)
			if err != nil {
				panic(err)
			}

			yml = aviator.ReadYaml(aviator.ResolveEnvVars(ymlBytes))
			err = aviator.ProcessSprucePlan(yml.Spruce)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			if yml.Fly.Target != "" && yml.Fly.Name != "" && yml.Fly.Config != "" {
				fmt.Println("Target set to", yml.Fly.Target)
				aviator.FlyPipeline(yml.Fly)
			}

		}

		return nil
	}
	cmd.Run(os.Args)
}

func VerifyAviatorFile(file string) bool {
	if file == "aviator.yml" {
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			return true
		}
	} else {
		sl := strings.Split(file, ".")
		if sl[len(sl)-1] == "vtr" {
			if _, err := os.Stat(file); !os.IsNotExist(err) {
				return true
			}
		}
	}
	return false
}
