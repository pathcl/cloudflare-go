package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

const (
	empty = ""
	tab   = "\t"
)

// PrettyJSON make Json look pretty
func PrettyJSON(data interface{}) (string, error) {

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
func main() {
	// Construct a new API object
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	// Fetch user details on the account
	u, err := api.UserDetails()
	if err != nil {
		log.Fatal(err)
	}

	// Print user details
	uinfo, err := PrettyJSON(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uinfo)

	// Fetch the zone ID
	id, err := api.ZoneIDByName("sanmartin.io") // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	r, err := PrettyJSON(zone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
