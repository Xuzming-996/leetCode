package _60_170

import "strings"

//168 Excel转列
func convertToTitle(columnNumber int) string {
	column := strings.Builder{}
	for columnNumber > 0 {
		c := columnNumber % 26
		if c == 0 {
			c = 26
			columnNumber -= 1
		}
		column.WriteString(string('A' + c - 1))
		columnNumber /= 26
	}
	tmp := []byte(column.String())
	var tmp1 []byte
	for i := len(tmp) - 1; i >= 0; i-- {
		tmp1 = append(tmp1, tmp[i])
	}

	return string(tmp1)
}
