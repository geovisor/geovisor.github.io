package main

import(
    "fmt"
    "bytes"
    "os"
    "strings"
    "github.com/tealeg/xlsx"
)
func main(){
    fmt.Println("HelloWorld.")
    os.Remove("data.json")
    file, err := os.Create("data.json")
    excelFileName := "tj.xlsx"
    xlFile, err2 := xlsx.OpenFile(excelFileName)
    if err2 != nil {
        fmt.Println("open failed: %s\n", err)
    }
    for _, sheet := range xlFile.Sheets[1:] {
        fmt.Printf("Sheet Name: %s\n", sheet.Name)
    }
    sheet1 := xlFile.Sheets[0]
    var buffer bytes.Buffer
    buffer.WriteString("{\"message\":\"\",\"value\":[")
    for _, row := range sheet1.Rows[1:] {
         t1 := row.Cells[0].String()
         t2 := row.Cells[1].String()
         t3 := row.Cells[2].String()
         t4 := row.Cells[3].String()
         t5 := row.Cells[4].String()
         if strings.Contains(row.Cells[5].String(), "$") && "" != t4 {
             t6 := strings.Split(strings.Split(row.Cells[5].String(), " ")[1], ".")[0] + "元"
             fmt.Println(t1+ "-" + t2 + "-" + t3 +  "-" + t4 + "-" + t5 + "-" + t6)
             buffer.WriteString("{")
             buffer.WriteString("\"q\":\"" + t2 + "\",")
             buffer.WriteString("\"l\":\"" + t3 + "\",")
             buffer.WriteString("\"x\":\"" + t4 + "\",")
             buffer.WriteString("\"y\":\"" + t1 + "\",")
             buffer.WriteString("\"f\":\"" + t6 + "\"")
             buffer.WriteString("},")
         }
    }
    buffer.WriteString("]}")
    file.Write([]byte(buffer.Bytes()))
}
