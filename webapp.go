package main

import (
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/upload", uploadhandler)
	m.Run()
}
func uploadhandler(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Capture multipart form file information
	file, handler, err := req.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
	}

	// Read the file into memory
	data, err := ioutil.ReadAll(file)
	// ... check err value for nil

	// Specify the Mongodb database
	my_db := session.DB("Images")
	filename := handler.Filename

	// Create the file in the Mongodb Gridfs instance
	my_file, err := my_db.GridFS("fs").Create(filename)
	// ... check err value for nil

	// Write the file to the Mongodb Gridfs instance
	n, err := my_file.Write(data)
	// ... check err value for nil

	// Close the file
	err = my_file.Close()
	// ... check err value for nil

	// Write a log type message
	fmt.Printf("%d bytes written to the Mongodb instance\n", n)

	// ... other statements redirecting to rest of user flow...
}
 /*func uploadhandler(r *http.Request){

	 session, err := mgo.Dial("127.0.0.1:27017")
	 if err != nil {
		 panic(err)
	 }

	 defer session.Close()
	 db := session.DB("Images")

	 file, header, err:= r.FormFile("uploadfile") // the input file by form
	 defer file.Close()

 }*/