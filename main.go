package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
	if err != nil{
		log.Printf("Error: %v\n",err);
	}

	for _, record : range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecods, err := net.LookupTXT("_dmarc." + domain)
	if err != nil{
		lof.Printf("Error",err)
	}

	for _, record : range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	
}