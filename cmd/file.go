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
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/companieshouse/btd-cli/pkg/btd"
	"github.com/companieshouse/btd-cli/pkg/btd/renderer/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file <path>",
	Short: "Parse business transaction data from an input file",
	Long: `Parse the content of a file containing business transaction data (BTD) into a
human-readable output format. Each line within the file is assumed to contain a
complete business transaction data string.
	
Examples:
  btd parse file <path>`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tagMap, err := btd.LoadTagMap(os.ExpandEnv(viper.GetString("tag-map")))
		if err != nil {
			return err
		}

		fmt.Println("Using tag map:", tagMap.LoadedFromFile())

		arg := args[0]

		if len(arg) <= 0 {
			return errors.New("filename cannot be empty")
		}

		var r btd.Renderer = table.New()

		file, err := os.Open(arg)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		line := 1

		for scanner.Scan() {
			if len(scanner.Text()) > 0 {
				data, err := tagMap.ParseTagData(scanner.Text())
				if err != nil {
					return err
				}

				fmt.Printf("%v:%d:\n", arg, line)
				fmt.Println(r.Render(data))
			}

			line++
		}

		file.Close()

		return nil
	},
}

func init() {
	parseCmd.AddCommand(fileCmd)
}
