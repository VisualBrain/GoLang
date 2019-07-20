package main 

import (
	"fmt"
	"os" //getting us allow to access file system
	"time" //convert unix timestamp to time object
	"io/ioutil" //to read the file contents
	"strings" //allow us to read the contents of files in reader
	"encoding/csv" //to get all contents from csv files
	"strconv" //help us convert non-string to string
	"runtime"
)
//filewatcher
const watchedPath = "./source"
//Parallel Processing - using multiple processor to do the same work
func main() {
	
	runtime.GOMAXPROCS(3)
	
	for {
		d,_ := os.Open(watchedPath) //Open() - works for both files and folders
		files,_ := d.Readdir(-1) //negative - to return all the files contain in particular path
		for INDEX,file := range files {
			//we create the filepath
			fmt.Println("File Is Running on index-------------------------",INDEX)
			filepath := watchedPath+"/"+file.Name();
			 fi,_ := os.Open(filepath) //opening the file at filepath
			 data,_ := ioutil.ReadAll(fi) //reading the content of file // in bytes
			 fi.Close() //closing the file
			 os.Remove(filepath)//removing the file at filepath
			 
			 go func(data string) {  //Annonymous function - goroutine
			 	reader := csv.NewReader(strings.NewReader(data)) //csv.NewReader(r:reader) //strings.NewReader(data) - converts strings of data into Reader
			 	//csv - Read - return first record of file 
			 	//csv - ReadAll - read all records of file and returns slice of those records
			 	records,_ := reader.ReadAll()
			 	for i,record := range records {
			 		fmt.Println("Reading Record-----",i)
			 		invoice := new(Invoice)
			 		invoice.InvoiceNumber = record[0]
			 		invoice.Amount,_ = strconv.ParseFloat(record[1], 64) //bitsize = 64or 32
			 		invoice.PurchaseOrderNumber,_ = strconv.Atoi(record[2]) //string to integer
			 		unixTime,_ := strconv.ParseInt(record[3],10, 64)//convert string to number of base 10 and bitsize -64 (string,base,bitsize)
					invoice.TimeOfPurchase = time.Unix(unixTime, 0)	 //convert integers to seconds 
					fmt.Println("Running:")
					fmt.Printf("Received Amount:$%.2f for Invoice Number : %v \t Purchase Order Number : %d\t on Time : %v\n",invoice.Amount,invoice.InvoiceNumber,	invoice.PurchaseOrderNumber,invoice.TimeOfPurchase)
			 	}
			 }(string(data))
		}
	}
}

type Invoice struct {
	InvoiceNumber string
	Amount float64
	PurchaseOrderNumber int
	TimeOfPurchase time.Time
	
}

//this is using goroutines