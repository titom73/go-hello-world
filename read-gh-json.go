package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// User data structure
type User struct {
	Login string `json:"login"`
}

//PR data structure JSON
type PR struct {
	Name   string `json:"title"`
	Author User   `json:"user"`
	ID     int    `json:"number"`
	State  string `json:"state"`
}

// Milestone data structure
type Milestone struct {
	PRs []PR `json:"items"`
}

// MilestoneByte data model
type MilestoneByte struct {
	Content byte
}

// GetMilestone to get content from GH API engine
func GetMilestone(repository string, milestone string) []byte {
	var APIMilestone string = "https://api.github.com/search/issues?per_page=100&q=milestone:" + milestone + "+type:pr+repo:" + repository

	resp, err := http.Get(APIMilestone)
	if err != nil {
		// handle error
		fmt.Println("Error getting data from GH")
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// result = bodyBytes
	return bodyBytes
}

// ReadJSONFile to test reading a JSON file
func ReadJSONFile(file string) []byte {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	bodyBytes, _ := ioutil.ReadAll(jsonFile)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return bodyBytes
}

func main() {

	fmt.Print("-> Enter Repository name (org/name): ")
	var RepositoryName string
	fmt.Scanln(&RepositoryName)
	fmt.Print("-> Enter milestone: ")
	var MilestoneName string
	fmt.Scanln(&MilestoneName)
	fmt.Println("")

	// read our opened jsonFile as a byte array.
	MilestoneByteValue := GetMilestone(RepositoryName, MilestoneName)
	// MilestoneJSONBodyString := string(MilestoneByteValue)

	// we initialize our Milestone array
	var milestones Milestone

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(MilestoneByteValue, &milestones)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(milestones.PRs); i++ {
		// fmt.Printf(milestone.PRs[i])
		fmt.Println("PR name: " + milestones.PRs[i].Name)
		fmt.Println("PR state: " + milestones.PRs[i].State)
		fmt.Println("PR ID: " + strconv.Itoa(milestones.PRs[i].ID))
		fmt.Println("Author: " + milestones.PRs[i].Author.Login)
		fmt.Println("---")
	}
}
