package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"mygoexam/ocr"
)

const address = ":5050"

func main() {

	//test()

	http.HandleFunc("/upload", uploadHandler)

	log.Println("Server started and listening at:" + address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		panic(err)
	}

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		uploadFile(w, r)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload 10 MB.
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")
	defer file.Close()

	if err != nil {
		fmt.Println("Error Retrieving the file from http request.")
		fmt.Println(err)
		return
	}

	// read all of the contents of our uploaded file.
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fileContent := string(fileBytes)

	lines := strings.Split(fileContent, "\n\n")

	parsedLine := ""
	newLine := ""
	for _, l := range lines {
		newLine, err = ocr.ReadLine(strings.Trim(l, "\n"))
		if err != nil {
			//Error 500
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("File is not well formed."))
			return
		}
		parsedLine += newLine + "\n"

		log.Println(parsedLine)
	}

	w.WriteHeader(http.StatusCreated) //201 created because is a POST.
	w.Write([]byte(parsedLine))

}

func test() {

	line := "    _  _     _  _  _  _  _ " + "\n" +
		"  | _| _||_||_ |_   ||_||_|" + "\n" +
		"  ||_  _|  | _||_|  ||_| _|"

	fmt.Println("Input:")
	fmt.Println(line)

	parsedLine, err := ocr.ReadLine(line)
	if err != nil {
		panic(err)
	}
	fmt.Println("Output:")
	fmt.Println(parsedLine)

	fmt.Println("Reading file:")
	lines := readFile()
	fmt.Println("File output:")
	for _, line := range lines {
		parsedLine, err = ocr.ReadLine(strings.Trim(line, "\n"))
		fmt.Println(parsedLine)
	}

}

func readFile() []string {
	content, err := os.ReadFile("entradas.txt")
	if err != nil {
		panic(err)
	}
	digits := strings.Split(string(content), "\n\n")
	return digits
}
