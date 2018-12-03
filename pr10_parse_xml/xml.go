package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/xml"
)

// ParseXML just learning how to parse xml of a post
type ParseXML struct {
	XMLName 	xml.Name  	`xml:"post"`
	ID			string		`xml:"id,attr"`
	Content	string		`xml:"content"`
	Author	author		`xml:"author"`
	XML 		string		`xml:",innerxml"`
}
type author struct{
	ID			string		`xml:"id,attr"`
	Name		string		`xml:",chardata"`
}

func parseXMLHandler(w http.ResponseWriter, r *http.Request){
	parsexml := ParseXML{}
	err := xml.NewDecoder(r.Body).Decode(&parsexml)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Fprintln(w, parsexml.XMLName)
}

func main() {
	mux := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/xml", parseXMLHandler)
	mux.ListenAndServe()
}