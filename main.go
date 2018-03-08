package main

import (
	"crypto/md5"
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

	for i := 0; i < 15; i++ {
		bytes, _ := s.MarshalJSON()
		log.Printf("%X %v", md5.Sum(bytes), string(bytes))
	}
}
