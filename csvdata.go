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

// 

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func uploadForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/upload.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func uploadXLSX(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, _, err := r.FormFile("xlsxfile")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, "Failed to parse Excel file", http.StatusInternalServerError)
		return
	}

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		http.Error(w, "Failed to read rows", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<h2>Excel File Content:</h2><table border='1'>")
	for _, row := range rows {
		fmt.Fprint(w, "<tr>")
		for _, cell := range row {
			fmt.Fprintf(w, "<td>%s</td>", cell)
		}
		fmt.Fprint(w, "</tr>")
	}
	fmt.Fprint(w, "</table>")
}

func main() {
	http.HandleFunc("/", uploadForm)
	http.HandleFunc("/upload", uploadXLSX)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
