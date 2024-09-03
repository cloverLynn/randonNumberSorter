package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	highest := 0
	lowest := 0
	attempts := 0
	for i := 0; i < 100; i++ {
		attempt := sortRandom()
		if i == 0 {
			highest = attempt
			lowest = attempt
		} else {
			if attempt > highest {
				highest = attempt
			}
			if attempt < lowest {
				lowest = attempt
			}
		}
		attempts += attempt
	}
	fmt.Println("highest:", highest)
	fmt.Println("lowest:", lowest)
	fmt.Println("average", attempts/100)
}

func sortRandom() int {
	chartMap := createMap(20)
	times := 0
	do := true
	for do {
		solutionsMap := make(map[int]int)
		for i := 0; i < 20; i++ {
			solutionsMap[i] = 0
		}
		randNum := random(solutionsMap)
		spot, _ := determineFirst(randNum, chartMap)
		solutionsMap[spot] = randNum
		for i := 0; i < 20; i++ {
			randNum = random(solutionsMap)
			spot = determineSpot(randNum, solutionsMap, chartMap)
			if spot != -1 {
				solutionsMap[spot] = randNum
				if i == 18 {
					do = false
					times++
					/*fmt.Println(solutionsMap)
					fmt.Println(times)
					fmt.Println("---------------")*/
					return times
				}
			} else {
				i = 21
				/*fmt.Println("---------------")*/
				times++
			}
		}
	}
	return -1
}

// generate randomNumber between 1-999
func random(m map[int]int) int {
	randomNum := 0
	for randomNum == 0 {
		randomNum = rand.IntN(999)
	}
	/*fmt.Println(randomNum)*/
	return randomNum
}

func determineSpot(num int, m map[int]int, cm map[int]int) int {
	spot, _ := determineFirst(num, cm)
	if m[spot] == 0 {
		return spot
	} else {
		if num < m[spot] {
			if spot > 0 && m[spot-1] == 0 {
				return spot - 1

			}
		} else {
			if num > m[spot] {
				if spot < 4 && m[spot+1] == 0 {
					return spot + 1
				}
			}
		}
	}
	return -1
}

func determineFirst(num int, m map[int]int) (int, float64) {
	div := float64(num) / float64(1000)
	div = math.Round(div*100) / 10
	return m[num], div
}

func createMap(slots int) map[int]int {
	mn := 1.0
	mx := 1000.0
	sectionsArr := make(map[int]int)
	sections := mx / float64(slots)
	sectionCounter := 0
	for i := 1; i <= slots; i++ {
		sectionsArr[i-1] = i * int(sections)
	}
	m := make(map[int]int)
	for i := mn; i <= mx; i++ {
		div := (i / mx) * 1000
		if div <= float64(sectionsArr[sectionCounter]) {
			m[int(i)] = sectionCounter
		} else {
			sectionCounter++
			m[int(i)] = sectionCounter
		}
	}
	return m
}
