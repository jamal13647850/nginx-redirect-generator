package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	var domain, loc, rt string

	fmt.Println("Please enter domain in xlsx:")
	fmt.Scanln(&domain)

	fmt.Println("Please enter xls file name: //eg: redirects.xlsx")
	fmt.Scanln(&loc)

	fmt.Println("Please enter redirect type: //eg: 301")
	fmt.Scanln(&rt)

	xlFile, err := xlsx.OpenFile(loc)
	if err != nil {
		fmt.Printf("1- Error opening file: %v\n", err)
		return
	}

	f, err := os.OpenFile("redirects.conf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("2- Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			src := row.Cells[0].String()
			dest := row.Cells[1].String()
			src, _ = unquote(src)
			dest, _ = unquote(dest)
			stringData := fmt.Sprintf("\n location %s { \n \tadd_header X-Redirect-By \"jamalghasemi.com\"; \n \treturn %s %s;\n}", src, rt, dest)
			search := "location " + domain
			newString := strings.Replace(stringData, search, "location ", 1)
			f.WriteString(newString)
		}
	}
}

func unquote(s string) (string, error) {
	return s, nil // unquoting is not implemented in this example
}
