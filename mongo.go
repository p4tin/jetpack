package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

type TaxInfo struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Zipcode string        `bson:"zip"`
	State   string        `bson:"state"`
	Tax     string        `bson:"tax"`
}

func MongoMain() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("GeoDB").C("TaxInfo")

	_, err = c.Upsert(bson.M{"_id": "999999"}, &TaxInfo{Zipcode: "999999", State: "ZZ", Tax: "0.099"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Upsert(bson.M{"_id": "111111"}, &TaxInfo{Zipcode: "111111", State: "ZZ", Tax: "0.011"})
	if err != nil {
		log.Fatal(err)
	}

	result := TaxInfo{}
	err = c.Find(bson.M{"zip": os.Args[1]}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("State: %s, Rate: %s\n", result.State, result.Tax)

	c.Remove(bson.M{"_id": "111111"})
	c.Remove(bson.M{"_id": "999999"})
}
