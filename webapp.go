package main

import (
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/upload", uploadHandler)
	m.Run()
}

func connect() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
}


func uploadHandler(w http.ResponseWriter, r *http.Request) {

 	// the FormFile function takes in the POST input id file
 	file, header, err := r.FormFile("uploadfile")

 	if err != nil {
 		fmt.Fprintln(w, err)
 		return
 	}

 	defer file.Close()

 	out, err := os.Create("/tmp/uploadedfile")
 	if err != nil {
 		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
 		return
 	}

 	defer out.Close()

 	// write the content from POST to the file
 	_, err = io.Copy(out, file)
 	if err != nil {
 		fmt.Fprintln(w, err)
 	}

 	fmt.Fprintf(w, "File uploaded successfully : ")
 	fmt.Fprintf(w, header.Filename)
 }
