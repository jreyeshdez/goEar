// Go Script that downloads songs from goear.com
// http://www.goear.com
//
// Julian Reyes
// http://github.com/jreyeshdez
//
// License: MIT
// Copyright (C) 2016 by Julian Reyes
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
    "flag"
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "net/http"
    "io"
)

var urlpattern = "http://www.goear.com/plimiter.php?f=%s"

func getDownloadUrl(code string) string {
	urlFormatted := fmt.Sprintf(urlpattern, code)
	return urlFormatted
}

func parseUrl(url string) (string,string) {
	code := strings.Split(url,"/")[4]
	newUrl := getDownloadUrl(code) 
	fileName := strings.Split(url,"/")[5] + ".mp3"	
	return fileName, newUrl
}

func download(url string) {
	fileName, link := parseUrl(strings.Replace(url,"\n","",1))
	out, err := os.Create(fileName)
	if err != nil {
        	log.Fatal(err)
    	}
	defer out.Close()	
	resp, err := http.Get(link)
 	if err != nil {
        	log.Fatal(err)
    	}
	defer resp.Body.Close()
	fmt.Printf("Downloading: %v\n", fileName)
	io.Copy(out, resp.Body)
}

func readFile(file string){
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		download(scanner.Text())	
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)	
	}
}

func usage(){
	fmt.Fprintf(os.Stderr, "usage: %s [inputfile]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Input file is missing..")
		os.Exit(1)	
	}
	fmt.Printf("Opening: %s\n", os.Args[1]) 
	readFile(os.Args[1])    
}
