package asciitable

import "strings"
import "testing"

func Test_Table(t *testing.T) {
	twoSections := [][]string{
		{"First", "Second", "Third"},
		nil,
		{"Foo", "Bar", "Foobar Here"},
		{"Another one", "Can", "Not hurt"},
	}
	threeSections := append(twoSections, nil, []string{"", "", "Cool"})

	tests := []struct {
		Data     [][]string
		Expected string
	}{
		{
			Data: twoSections,
			Expected: strings.TrimSpace(`
+-------------+--------+-------------+
| First       | Second | Third       |
+-------------+--------+-------------+
| Foo         | Bar    | Foobar Here |
| Another one | Can    | Not hurt    |
+-------------+--------+-------------+
`,
			),
		},
		{
			Data: threeSections,
			Expected: strings.TrimSpace(`
+-------------+--------+-------------+
| First       | Second | Third       |
+-------------+--------+-------------+
| Foo         | Bar    | Foobar Here |
| Another one | Can    | Not hurt    |
+-------------+--------+-------------+
|             |        | Cool        |
+-------------+--------+-------------+
`,
			),
		},
	}
	for _, test := range tests {
		table := NewTable()
		for _, row := range test.Data {
			if row != nil {
				table.AddRow(row...)
			} else {
				table.AddSeparator()
			}
		}
		if got := table.String(); got != test.Expected {
			t.Errorf("\nGot:\n%s\nExpected:\n%s\n", got, test.Expected)
		}
	}
}
