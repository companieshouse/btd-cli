/*
Copyright © 2023 Crown Copyright (Companies House)

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
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse business transaction data into a human-readable format",
	Long: `Parse the content of a file or command-line argument string containing
business transaction data (BTD) into a human-readable output format. Use the
subcommands 'file' and 'string' to read the transaction data from a file or string
argument respectively.

When parsing the content of a file, each line within the file is assumed to
contain a complete business transaction data string.

String arguments must be quoted (single or double) when using the 'data'
subcommand.

Examples:
  btd-cli parse data '...'
  btd-cli parse file <path>`,
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
