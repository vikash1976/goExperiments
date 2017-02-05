package main

import (
    "encoding/csv"
    "text/template"
    "os"
    "strconv"
    "time"
    "log"
    //"fmt"
)

//Record - data structure reprsenting records
type Record struct {
    Date time.Time
    Open float64
}
var tpl *template.Template

//YYYYMMDD a date formater function
func YYYYMMDD (t time.Time) string {
    return t.Format("2006-02-01")
}

var funcMaps = template.FuncMap{
    "yyyyMMDD" : YYYYMMDD,
}
func init() {
    tpl = template.Must(template.New("").Funcs(funcMaps).ParseFiles("stock.gohtml"))
}
func parseCSV (filePath string) []Record {
    src, err := os.Open(filePath)
    
    if err != nil {
        log.Fatalln(err)
    }
    defer src.Close()
    reader := csv.NewReader(src)
    rows, err := reader.ReadAll()
    if err != nil {
        log.Fatalln(err)
    }
    records := make([]Record, 0, len(rows))
    
    for i, row := range rows {
        if i== 0 {
            continue
        }
        date, _ := time.Parse("2006-01-02", row[0])
        open, _ := strconv.ParseFloat(row[1], 64)
        
        records = append(records, Record {
            Date: date,
            Open: open,
        })
    }
    return records
}

func main() {
    records := parseCSV("stock.csv")
    
    err := tpl.ExecuteTemplate(os.Stdout, "stock.gohtml", records)
   
    if err != nil {
        log.Fatalln(err)
    }
    //fmt.Println(time.Parse("2006-01-02", "2014-06-06"))
}