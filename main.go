package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
	
	for scanner.Scan(){
		CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal("Error: Could not read from the input: %v\n",err)
	}
}

func CheckDomain(domain string) {
	var hasSPF, hasDMARC, hasMX bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil{
		log.Printf("Error: %v\n",err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err!= nil{
		log.Printf("Error: %v\n",err)
	}
	
}