// package main

// import (
// 	"content"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

//  func main() {
//  	fmt.Println("Welcome to csvdata file .")
//  	content := "This is a file  $CSV DATA.xlsx "
//  	 file,err := os.Create("./$CSV DATA.xlsx")
// 	checkNilErr(err)
	

//  	 length,err := io.WriteString(file,content)
//      checkNilErr(err)

//  	 fmt.Println("length is: ",length)
//  	 defer file.Close()
// 	 readFile("./CSV DATA.xlsx")
//  }
// 	 func readFile(filname string) {
// 		databyte, err := ioutil.ReadFile(filname)
	 
// 	 checkNilErr(err)
// 	 fmt.Println("Text data inside a file is \n", string(databyte))
//  }
//    func checkNilErr(err error){
//        if err !=nil {
//  		panic(err)
//  	 }
//    }

package main

import ( "fmt"
     

)

   

func main() {
    f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Create a new sheet.
    index, err := f.NewSheet("Sheet2")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Set value of a cell.
    f.SetCellValue("Sheet2", "A2", "Hello world.")
    f.SetCellValue("Sheet1", "B2", 100)
    // Set active sheet of the workbook.
    f.SetActiveSheet(index)
    // Save spreadsheet by the given path.
    if err := f.SaveAs("CSV DATA.xlsx"); err != nil {
        fmt.Println(err)
    }
}
