package formatter

import (
	"testing"
)

func TestFormatTable(t *testing.T) {
	tests := []struct {
		name     string
		data     [][]string
		expected string
	}{
		{
			name: "simple table",
			data: [][]string{
				{"Name", "Age"},
				{"Alice", "30"},
				{"Bob", "25"},
			},
			expected: "" +
				"+-------+-----+\n" +
				"| Name  | Age |\n" +
				"+-------+-----+\n" +
				"| Alice | 30  |\n" +
				"| Bob   | 25  |\n" +
				"+-------+-----+\n",
		},
		{
			name:     "empty table",
			data:     [][]string{},
			expected: "",
		},
		{
			name: "single column",
			data: [][]string{
				{"Items"},
				{"Apple"},
				{"Banana"},
			},
			expected: "" +
				"+--------+\n" +
				"| Items  |\n" +
				"+--------+\n" +
				"| Apple  |\n" +
				"| Banana |\n" +
				"+--------+\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTable(tt.data)
			if got != tt.expected {
				t.Errorf("FormatTable() = \n%v\nwant\n%v", got, tt.expected)
			}
		})
	}
}

func TestFormatCSV(t *testing.T) {
	tests := []struct {
		name      string
		data      [][]string
		delimiter string
		expected  string
	}{
		{
			name: "simple data",
			data: [][]string{
				{"Name", "Age", "City"},
				{"Alice", "30", "New York"},
				{"Bob", "25", "Los Angeles"},
			},
			delimiter: ",",
			expected:  "Name,Age,City\nAlice,30,New York\nBob,25,Los Angeles\n",
		},
		{
			name: "with quotes needed",
			data: [][]string{
				{"Name", "Description"},
				{"Product", "A \"great\" item"},
				{"Service", "Includes, many options"},
			},
			delimiter: ",",
			expected:  "Name,Description\nProduct,\"A \"\"great\"\" item\"\nService,\"Includes, many options\"\n",
		},
		{
			name: "custom delimiter",
			data: [][]string{
				{"Name", "Age", "City"},
				{"Alice", "30", "New York"},
				{"Bob", "25", "Los Angeles"},
			},
			delimiter: ";",
			expected:  "Name;Age;City\nAlice;30;New York\nBob;25;Los Angeles\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatCSV(tt.data, tt.delimiter)
			if got != tt.expected {
				t.Errorf("FormatCSV() = \n%v\nwant\n%v", got, tt.expected)
			}
		})
	}
}

// We skip testing FormatList and FormatNumberedList to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
