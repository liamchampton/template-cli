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

	"github.com/spf13/cobra"
)

var ErrNoSubCommand = errors.New("Sub-command expected - use `create -h` to see an example of the expected input")

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "A brief description of your command",
	Example: "create golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ErrNoSubCommand)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
