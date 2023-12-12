package table

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/companieshouse/btd-cli/pkg/btd"
)

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
)

type ColumnId int

const (
	idColumn ColumnId = iota
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
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)

		IdColumnStyle     = re.NewStyle().Width(6).Align(lipgloss.Center)
		XmlTagColumnStyle = re.NewStyle().Width(20).Align(lipgloss.Center)
		LengthColumnStyle = re.NewStyle().Width(8).Align(lipgloss.Center)
		DataColumnStyle   = re.NewStyle().Width(40).MaxWidth(40).MaxHeight(100)

		OddRowStyle = re.NewStyle().Foreground(gray)

		EvenRowStyle = re.NewStyle().Foreground(lightGray)
	)

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			// style := DefaultCellStyle
			var style lipgloss.Style

			switch {
			case row == 0:
				style = style.Inherit(HeaderStyle)
			case row%2 == 0:
				style = style.Inherit(EvenRowStyle)
			case row%2 != 0:
				style = style.Inherit(OddRowStyle)
			}

			switch ColumnId(col) {
			case idColumn:
				style = style.Inherit(IdColumnStyle)
			case xmlTagColumn:
				style = style.Inherit(XmlTagColumnStyle)
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
