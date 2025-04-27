package formatter

import (
	"strings"
)

// FormatTable formats a 2D slice of strings as a text table
func FormatTable(data [][]string) string {
	if len(data) == 0 {
		return ""
	}

	// Calculate the maximum width of each column
	numCols := len(data[0])
	colWidths := make([]int, numCols)

	for _, row := range data {
		for i, cell := range row {
			if i < numCols && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// Build the table
	var builder strings.Builder

	// Horizontal separator line
	writeSeparator := func() {
		builder.WriteString("+")
		for _, width := range colWidths {
			builder.WriteString(strings.Repeat("-", width+2))
			builder.WriteString("+")
		}
		builder.WriteString("\n")
	}

	writeSeparator()

	// Write header row
	for i, cell := range data[0] {
		builder.WriteString("| ")
		builder.WriteString(cell)
		builder.WriteString(strings.Repeat(" ", colWidths[i]-len(cell)))
		builder.WriteString(" ")
	}
	builder.WriteString("|\n")

	writeSeparator()

	// Write data rows
	for i := 1; i < len(data); i++ {
		row := data[i]
		for j, cell := range row {
			if j < numCols {
				builder.WriteString("| ")
				builder.WriteString(cell)
				builder.WriteString(strings.Repeat(" ", colWidths[j]-len(cell)))
				builder.WriteString(" ")
			}
		}
		builder.WriteString("|\n")
	}

	writeSeparator()

	return builder.String()
}

// FormatCSV formats a 2D slice of strings as a CSV string
func FormatCSV(data [][]string, delimiter string) string {
	if len(data) == 0 {
		return ""
	}

	if delimiter == "" {
		delimiter = ","
	}

	var builder strings.Builder

	for _, row := range data {
		for i, cell := range row {
			// Escape double quotes by doubling them
			escaped := strings.ReplaceAll(cell, "\"", "\"\"")

			// Wrap in quotes if the cell contains delimiter, quotes, or newlines
			if strings.ContainsAny(cell, delimiter+"\"\n\r") {
				builder.WriteString("\"")
				builder.WriteString(escaped)
				builder.WriteString("\"")
			} else {
				builder.WriteString(escaped)
			}

			if i < len(row)-1 {
				builder.WriteString(delimiter)
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

// FormatList formats a slice of strings as a bulleted list
func FormatList(items []string, bullet string) string {
	if len(items) == 0 {
		return ""
	}

	if bullet == "" {
		bullet = "• "
	}

	var builder strings.Builder

	for _, item := range items {
		builder.WriteString(bullet)
		builder.WriteString(item)
		builder.WriteString("\n")
	}

	return builder.String()
}

// FormatNumberedList formats a slice of strings as a numbered list
func FormatNumberedList(items []string) string {
	if len(items) == 0 {
		return ""
	}

	var builder strings.Builder

	for i, item := range items {
		builder.WriteString(strings.TrimSpace(strings.Join([]string{
			strings.TrimSpace(string(rune('1' + i))),
			".",
			strings.TrimSpace(item),
		}, " ")))
		builder.WriteString("\n")
	}

	return builder.String()
}
