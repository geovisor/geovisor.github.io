package main

import (
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

func main() {
	jsonFile := "data.json"
	excelFile := "058-邮编.xlsx"
	os.Remove(jsonFile)
	fmt.Println("Start format excel to data.json")
	file, err := os.Create(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	xlFile, err := xlsx.OpenFile(excelFile)
	sheet1 := xlFile.Sheets[0]
	var buffer bytes.Buffer
	buffer.WriteString("{\"message\":\"\",\"value\":[")
	for _, row := range sheet1.Rows[1:] {
		t6 := row.Cells[5].String()
		if strings.Contains(t6, "$") && "" != row.Cells[3].String() {
			buffer.WriteString("{")
			money := strings.Split(strings.Split(t6, " ")[1], ".")[0]
			buffer.WriteString("\"q\":\"" + row.Cells[1].String() + "\",")
			buffer.WriteString("\"l\":\"" + row.Cells[2].String() + "\",")
			buffer.WriteString("\"x\":\"" + row.Cells[3].String() + "\",")
			buffer.WriteString("\"y\":\"" + row.Cells[0].String() + "\",")
			buffer.WriteString("\"f\":\"" + money + "\"")
			buffer.WriteString("},")
			fmt.Println(row.Cells[1].String() + "-" + row.Cells[2].String() + "-" + row.Cells[3].String() + "-" + money)
		}
	}
	buffer.Truncate(buffer.Len() - 1)
	buffer.WriteString("]}")
	file.Write([]byte(buffer.Bytes()))
}

