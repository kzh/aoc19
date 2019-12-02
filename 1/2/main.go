package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(file)
	fuel := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return
		}

		line := scanner.Text()
		mass, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
			return
		}

		temp := int(math.Floor(mass/3)) - 2
		for temp > 0 {
			fuel += temp
			temp = int(math.Floor(float64(temp/3)) - 2)
		}
	}

	log.Println(fuel)
}
