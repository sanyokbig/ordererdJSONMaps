package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

//easyjson:json
type S struct {
	I              int
	StringToFilter OrderedMapStringToFilter
}

//easyjson:json
type Filter struct {
	Name        string
	Value       int
	StringToInt OrderedMapStringToInt
}

func main() {
	s := S{
		I: 1,
		StringToFilter: OrderedMapStringToFilter{
			"b": Filter{
				Name:  "filter B",
				Value: 2,
				StringToInt: OrderedMapStringToInt{
					"z": 12,
					"x": 10,
					"y": 11,
				},
			},
			"c": Filter{
				Name:  "filter B",
				Value: 2,
				StringToInt: OrderedMapStringToInt{
					"z": 12,
					"x": 10,
					"y": 11,
				},
			},
			"a": Filter{
				Name:  "filter B",
				Value: 2,
				StringToInt: OrderedMapStringToInt{
					"z": 12,
					"x": 10,
					"y": 11,
				},
			}},
	}

	results := map[string]int{}
	for i := 0; i < 15; i++ {
		bytes, _ := s.MarshalJSON()
		hash := fmt.Sprintf("%X", md5.Sum(bytes))
		log.Printf("%v %v", hash, string(bytes))
		if count, ok := results[hash]; ok {
			results[hash] = count + 1
		} else {
			results[hash] = 1
		}
	}
	log.Println("results: ")
	for hash, count := range results {
		log.Printf("%v: %v", hash, count)
	}
}
