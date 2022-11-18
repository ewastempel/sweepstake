package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type Participant struct {
	Name     string `json:"name"`
	BetCount int    `json:"betCount"`
}

type Participants struct {
	Participants []Participant `json:"participants"`
}

type Country struct {
	Name string `json:"name"`
}

type Countries struct {
	Countries []Country `json:"countries"`
}

func getCountries(countriesJSON string) map[int]string {
	result := make(map[int]string)
	jsonFile, err := os.Open(countriesJSON)
	if err != nil {
		fmt.Println(err)
	}

	countriesArray, _ := io.ReadAll(jsonFile)

	var countries Countries
	json.Unmarshal(countriesArray, &countries)

	defer jsonFile.Close()

	fmt.Println("Countries count:", len(countries.Countries))
	for i := 0; i < len(countries.Countries); i++ {
		//fmt.Println("Country name:", countries.Countries[i].Name)
		//fmt.Println(countries.Countries[i].Name)
		result[i] = countries.Countries[i].Name
	}

	//fmt.Println(result)
	return result
}

func getParticipants(participantsJSON string) map[string]int {
	result := make(map[string]int)
	jsonFile, err := os.Open(participantsJSON)
	if err != nil {
		fmt.Println(err)
	}

	participantsArray, _ := io.ReadAll(jsonFile)

	var participants Participants
	json.Unmarshal(participantsArray, &participants)

	defer jsonFile.Close()

	fmt.Println("Participants count:", len(participants.Participants))
	var betCount int
	for i := 0; i < len(participants.Participants); i++ {
		betCount += participants.Participants[i].BetCount
		result[participants.Participants[i].Name] = participants.Participants[i].BetCount
	}
	fmt.Println("Total bet count:", betCount)

	//fmt.Println(result)
	return result
}

func AssignCountriesToParticipants(countries map[int]string, participants map[string]int) (bets map[string][]string) {
	bets = make(map[string][]string)
	remainingCountries := countries
	fmt.Println("Participating Countries:", remainingCountries)
	var keys []int
	// Because it is iterating through a map, the elements are returned randomly and therefore the keys are already shuffled
	for key, _ := range remainingCountries {
		keys = append(keys, key)
	}
	//fmt.Println(keys)

	// Shuffling the keys again (not relying on the map randomness)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	// Print the countries after shuffling the keys
	//for key := range keys {
	//	fmt.Println(remainingCountries[keys[key]])
	//}
	//fmt.Println(keys)

	index := 0
	for name, count := range participants {
		//fmt.Println("name:", name, ", bets:", count)
		var participantCountries []string
		for i := 0; i < count; i++ {
			participantCountries = append(participantCountries, remainingCountries[keys[index]])
			delete(remainingCountries, keys[index])
			index++
		}
		bets[name] = participantCountries
	}
	fmt.Println("Iterated through", index, "countries,", "remaining countries:", remainingCountries)
	fmt.Println("")
	return
}

func main() {
	list := AssignCountriesToParticipants(getCountries("data/countries.json"), getParticipants("data/participants.json"))
	fmt.Println("Final List:")
	for name, countries := range list {
		fmt.Print(name, ": ")
		for _, country := range countries {
			fmt.Print(country, ", ")
		}
		fmt.Println("")
	}
}
