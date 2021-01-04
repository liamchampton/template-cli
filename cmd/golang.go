/*
Copyright © 2020 Liam Hampton liam.hampton@hotmail.co.uk

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/template-cli/pkg/golang"
)

var ErrGoNotInstalled = errors.New("Golang not installed")
var ErrGoPathDoesNotExist = errors.New("expected go path does not exist on the system")

// golangCmd represents the golang command
var Name string
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for go installation")

		// Get User
		user, err := user.Current()
		if err != nil {
			fmt.Println(err)
		}

		installed, err := golang.CheckInstallation()
		if err != nil {
			fmt.Println(ErrGoNotInstalled)
			os.Exit(1)
		}

		if installed == true {

			// check for typical file structure ~/go/src/github.com
			fmt.Println("Checking expected golang file structure ~/go/src/github.com exists on the system")

			if _, err := os.Stat(user.HomeDir + "/go/src/github.com/"); !os.IsNotExist(err) {
				fmt.Println("path exists\ncreating new project directory")

				// if _, err := os.Stat(user.HomeDir + "Desktop/testdir"); os.IsNotExist(err) {
				// 	//os.Mkdir("/Users/liam/Desktop/testdir", 0700)
				// 	fmt.Println("Making new directory foo bar (test comment)")
				// }

				os.Mkdir(user.HomeDir+"/go/src/github.com/"+args[1], 0700)

			} else {
				fmt.Println(ErrGoPathDoesNotExist)
				os.Exit(1)
			}

			// if file structure exists then create a new directory

			// inside directory clone the template code from github
		}
	},
}

func init() {
	createCmd.AddCommand(golangCmd)

	createCmd.Flags().StringVarP(&Name, "project-name", "n", "", "Name of new project")
	createCmd.MarkFlagRequired("project-name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// golangCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// golangCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
