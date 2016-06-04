package main

import (
	"bufio"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("state_abbr_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		state := scanner.Text()
		filename := "TAXRATES_ZIP5_" + state + "201506.csv"
		url := "https://s3-us-west-2.amazonaws.com/taxrates.csv/TAXRATES_ZIP5_" + state + "201506.csv"
		//fmt.Println(url)
		downloadFromUrl(url)
		//read the new file and insert data in mongo
		taxfile, taxerr := os.Open("taximports/" + filename)
		if taxerr != nil {
			log.Fatal(taxerr)
		}
		defer taxfile.Close()

		scanner := bufio.NewScanner(taxfile)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ",")
			write_to_mongo(parts[1], parts[0], parts[4])
			fmt.Println(parts[0] + ", " + parts[1] + ", " + parts[4])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create("taximports/" + fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

type TaxInfo struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Zipcode string        `bson:"zip"`
	State   string        `bson:"state"`
	Tax     string        `bson:"tax"`
}

func write_to_mongo(pZip, pState, pTax string) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("GeoDB").C("TaxInfo")

	_, err = c.Upsert(bson.M{"_id": pZip}, &TaxInfo{Zipcode: pZip, State: pState, Tax: pTax})
	if err != nil {
		log.Fatal(err)
	}
}
