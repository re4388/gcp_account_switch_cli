package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bitfield/script"
	"github.com/urfave/cli/v2"
)

var correctConfigPath = "/Users/re4388/.config/gcloud/configurations/config_default"

/*
Here, I already setup two path in the dir xxxx-personal and xxx-wemo
So, I simple copy to a tmp file and mv to the correct file
*/
func copyAndRename(configPath string) {
	// copy config_default_personal to config_default_personal-tmp
	copy_script := fmt.Sprintf("cp %s %s-tmp", configPath, configPath)
	// fmt.Println(cp_script) // for debug
	script.Exec(copy_script).Wait()

	// rename to config_default_personal-tmo to config_default
	rename_script := fmt.Sprintf("mv %s-tmp %s", configPath, correctConfigPath)
	// fmt.Println(cp_script)
	script.Exec(rename_script).Wait()

}

func main() {

	app := &cli.App{
		Name:  "GCP Account Switcher",
		Usage: "./gcpSwticher wemo  or ./gcpSwticher personal",
		Action: func(cCtx *cli.Context) error {
			env := cCtx.Args().Get(0)

			switch env {
			case "wemo":
				fmt.Printf("switch to %s...\n", env)
				copyAndRename("/Users/re4388/.config/gcloud/configurations/config_default_wemo")
			case "personal":
				fmt.Printf("switch to %s...\n", env)
				copyAndRename("/Users/re4388/.config/gcloud/configurations/config_default_personal")

			default:
				fmt.Printf("no this env, choose `wemo` or `personal`")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")
}
