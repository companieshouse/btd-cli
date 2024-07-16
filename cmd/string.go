/*
Copyright © 2023 Companies House (Crown Copyright)

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
	"errors"
	"fmt"
	"os"

	"github.com/companieshouse/btd-cli/pkg/btd"
	"github.com/companieshouse/btd-cli/pkg/btd/renderer/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stringCmd represents the string command
var stringCmd = &cobra.Command{
	Use:   "string <btd>",
	Short: "Parse business transaction data from a string",
	Long: `Parse a command-line argument string containing business transaction data (BTD)
into a human-readable output format. The string arguments must be quoted
(single or double) when using this subcommand.

Examples:
  btd-cli parse string '...'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tagMap, err := btd.LoadTagMap(os.ExpandEnv(viper.GetString("tag-map")))
		if err != nil {
			return err
		}

		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println("Using tag map:", tagMap.LoadedFromFile())

		path := args[0]

		if len(path) <= 0 {
			return errors.New("business transaction data string cannot be empty")
		}

		data, err := tagMap.ParseTagData(path)
		if err != nil {
			return err
		}

		fmt.Println(table.New().Render(data))

		return nil
	},
}

func init() {
	parseCmd.AddCommand(stringCmd)
}
