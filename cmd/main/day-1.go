package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() {
	file := "./day-1-input.txt"
	elves := separateEvlesFromFile(file)

	maxCalories := maxIntFromElvesArray(elves, 0)
	secondMostCalories := maxIntFromElvesArray(elves, maxCalories)
	thirdMostCalories := maxIntFromElvesArray(elves, secondMostCalories)

	topThree := thirdMostCalories + secondMostCalories + maxCalories

	fmt.Printf("first: %d, second: %d, third: %d\n", maxCalories, secondMostCalories, thirdMostCalories)

	fmt.Printf("Max calories among the elves: %d\n", maxCalories)
	fmt.Printf("Calories across the top 3 elves: %d", topThree)
}

func separateEvlesFromFile(filename string) [][]int {
	calorie_list := calorieList(filename)
	var elf []int
	var elves [][]int

	for _, v := range calorie_list {
		if v == "" {
			elves = append(elves, elf)
			elf = nil
			continue
		}
		calories, _ := strconv.Atoi(v)

		elf = append(elf, calories)
	}

	elves = append(elves, elf)

	return elves
}

func calorieList(filename string) []string {
	var list []string
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}

func maxIntFromElvesArray(elves [][]int, topLevel int) int {
	calories := 0
	maxCalories := 0
	for _, elf := range elves {
		for _, calorie := range elf {
			calories += calorie
		}

		// this is gross. refactor
		if calories > maxCalories && topLevel == 0 {
			maxCalories = calories
		} else if topLevel != 0 && calories < topLevel && calories > maxCalories {
			maxCalories = calories
		}
		calories = 0
	}

	return maxCalories
}
