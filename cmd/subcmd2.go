/*
Copyright © 2024 hassaku63 <hassaku63@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	token    string
	intParam int64
)

// subcmd2Cmd represents the subcmd2 command
var subcmd2Cmd = &cobra.Command{
	Use:   "subcmd2",
	Short: "A brief description of your command",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("subcmd2.PersistentPreRun called")

		if token == "" {
			token = os.Getenv("API_TOKEN")
			if token == "" {
				return fmt.Errorf("--token is required")
			}
		}

		if intParam < 1 || intParam > 100 {
			return fmt.Errorf("intParam should be in range [1, 100]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subcmd2.Run called")
		fmt.Printf("[subcmd2] debug: %v\n", debugMode)
		fmt.Printf("[subcmd2] token: %v\n", token)
		fmt.Printf("[subcmd2] intParam: %v\n", intParam)
	},
}

func init() {
	rootCmd.AddCommand(subcmd2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	subcmd2Cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "API token for subcmd2, If not set, use the value from API_TOKEN environment variable")
	subcmd2Cmd.PersistentFlags().Int64VarP(&intParam, "int", "i", 10, "common parameter for subcmd2, (1 <= int <= 100)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subcmd2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
