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
package table

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/companieshouse/btd-cli/pkg/btd"
	"golang.org/x/term"
)

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
)

type ColumnID int

const (
	idColumn ColumnID = iota
	xmlTagColumn
	lengthColumn
	dataColumn
)

type Table struct{}

func New() *Table {
	return &Table{}
}

func (t *Table) Render(data btd.TagData) string {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		IDColumnWidth     = 6
		XMLTagColumnWidth = 20
		LengthColumnWidth = 8
		DataColumnWidth   = 10 // miniumum width; dyanmically scaled below
		ColumnPadding     = 5
	)

	if max_data_length := data.GetMaxDataLength(); max_data_length > DataColumnWidth {
		tty_width, _, err := term.GetSize(0)
		if err != nil {
			fmt.Fprint(os.Stderr, "Warning: unable to determine terminal width")
		}

		max_data_column_width := tty_width - IDColumnWidth - XMLTagColumnWidth - LengthColumnWidth - ColumnPadding

		if max_data_length > max_data_column_width {
			DataColumnWidth = max_data_column_width
		} else {
			DataColumnWidth = max_data_length
		}
	}

	var (
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)

		IDColumnStyle     = re.NewStyle().Width(IDColumnWidth).Align(lipgloss.Center)
		XMLTagColumnStyle = re.NewStyle().Width(XMLTagColumnWidth).Align(lipgloss.Center)
		LengthColumnStyle = re.NewStyle().Width(LengthColumnWidth).Align(lipgloss.Center)
		DataColumnStyle   = re.NewStyle().Width(DataColumnWidth)

		OddRowStyle = re.NewStyle().Foreground(gray)

		EvenRowStyle = re.NewStyle().Foreground(lightGray)
	)

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == table.HeaderRow:
				style = style.Inherit(HeaderStyle)
			case row%2 == 0:
				style = style.Inherit(EvenRowStyle)
			case row%2 != 0:
				style = style.Inherit(OddRowStyle)
			}

			switch ColumnID(col) {
			case idColumn:
				style = style.Inherit(IDColumnStyle)
			case xmlTagColumn:
				style = style.Inherit(XMLTagColumnStyle)
			case lengthColumn:
				style = style.Inherit(LengthColumnStyle)
			case dataColumn:
				style = style.Inherit(DataColumnStyle)
			}

			return style
		}).
		Headers("ID", "XML Tag", "Length", "Data").
		Rows(data...).
		String()
}
