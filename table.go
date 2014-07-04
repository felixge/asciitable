package asciitable

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func NewTable() *Table {
	return &Table{}
}

type Table struct {
	data [][]string
}

func (t *Table) AddRow(vals ...string) {
	t.data = append(t.data, vals)
}

func (t *Table) AddSeparator() {
	t.data = append(t.data, nil)
}

func (t *Table) numColumns() int {
	num := 0
	for _, row := range t.data {
		if l := len(row); l > num {
			num = l
		}
	}
	return num
}

func (t *Table) lengths() []int {
	lengths := make([]int, t.numColumns())
	for _, row := range t.data {
		for i, val := range row {
			if l := len(val); l > lengths[i] {
				lengths[i] = l
			}
		}
	}
	return lengths
}

func (t *Table) WriteTo(w io.Writer) (n int64, err error) {
	bw := bufio.NewWriter(w)
	writeStr := func(str string) bool {
		var wn int
		wn, err = bw.WriteString(str)
		n += int64(wn)
		return err == nil
	}

	lengths := t.lengths()
	defer func() {
		flushErr := bw.Flush()
		if err == nil {
			err = flushErr
		}
	}()
	separator := "+"
	for _, l := range lengths {
		separator += strings.Repeat("-", l+2) + "+"
	}
	if !writeStr(separator + "\n") {
		return
	}
	for _, row := range t.data {
		if row == nil {
			writeStr(separator)
		} else {
			if !writeStr("|") {
				return
			}
			for c, l := range lengths {
				val := " " + row[c]
				if pad := l - len(val); pad > 0 {
					val += strings.Repeat(" ", pad+1)
				}
				val += " |"
				if !writeStr(val) {
					return
				}
			}
		}
		writeStr("\n")
	}
	if !writeStr(separator) {
		return
	}
	return
}

func (t *Table) String() string {
	buf := &bytes.Buffer{}
	_, err := t.WriteTo(buf)
	if err != nil {
		panic("Bug")
	}
	return buf.String()
}
