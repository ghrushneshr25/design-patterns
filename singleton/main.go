package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps := make(map[string]int)
		fmt.Println("Initializing database")
		caps["Seoul"] = 1234567
		caps["Tokyo"] = 4567890
		caps["Mexico City"] = 7890123
		instance = &singletonDatabase{capitals: caps}
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func GetTotalPopulationEx(database Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += database.GetPopulation(city)
	}
	return result
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println(pop)
	db2 := GetSingletonDatabase()
	pop = db2.GetPopulation("Tokyo")
	fmt.Println(pop)

	fmt.Println("Total Population", GetTotalPopulation([]string{"Seoul", "Tokyo"}))
}
