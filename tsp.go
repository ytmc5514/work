package main

import (
        "encoding/csv"
        "fmt"
        "io"
        "os"
)

func fileReader() {

        var fp *os.File
        if len(os.Args) < 2 {
                fp = os.Stdin
        } else {
                var err error
                fp, err = os.Open("input_0.csv")
                if err != nil {
                        panic(err)
                }
                defer fp.Close()
        }

        reader := csv.NewReader(fp)
        reader.Comma = '\t'
        reader.LazyQuotes = true 
                for {
                record, err := reader.Read()
                if err == io.EOF {
                        break
                } else if err != nil {
                        panic(err)
                }
                fmt.Println(record)
        }
}


func distance ( city1, city2 ){
	