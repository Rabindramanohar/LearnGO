package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//  map containing (i, i*i) for each i between 1 and n (inclusive)
	var num int
	fmt.Print("Enter an integral number: ")
	fmt.Scanln(&num)

	mapResult := getSquareMap(num)
	fmt.Printf("map %v\n", mapResult)

	// WPG divisible by 7 but are not a multiple of 5, between 2000 and 3200
	getNumberDiv7(20, 32)

	// WP which takes a string of comma-separated numbers and returns a slice of int
	sliceOfInts, err := convertStringToSlice("12, 23, 34, 45, 56, 67, 78, 89, 90")
	if err != nil {
		fmt.Println("Error in converting String to Slice of Ints : ", err)
	} else {
		fmt.Println("Converted to Int: ", sliceOfInts)
	}

	// convert to upper case
	strInput := []string{"Hello world", "Practice makes perfect"}
	myToUpperCase(strInput)

    // concepts of map
    understandMap();

}

func getSquareMap(n int) map[int]int {
	sqMap := make(map[int]int)

	for i := 1; i <= n; i++ {
		sqMap[i] = i * i
	}

	return sqMap
}

func getNumberDiv7(start int, end int) {
	fmt.Println("divisible by 7 elements")

	for i := start; i <= end; i++ {
		if i%7 == 0 && i%5 != 0 {
			fmt.Print(i)

			if i < 3200 {
				fmt.Print(", ")
			}
		}

	}

	fmt.Println()
}

func convertStringToSlice(input string) ([]int, error) {
	strArr := strings.Split(input, ", ")
	var ans []int

	for _, ch := range strArr {
		num, err := strconv.Atoi(strings.TrimSpace(ch))

		if err != nil {
			return nil, err
		}

		ans = append(ans, num)
	}

	return ans, nil
}

func myToUpperCase(input []string) {
	var ans []string
	for _, word := range input { // "_" we use to bland identifier when we don't want index
		ans = append(ans, strings.ToUpper(word))
	}
	for _, line := range ans {
		fmt.Printf(line + " ")
	}
}

func understandMap() {
    // map initialization
    var map1 map[string]int;

    // empty map declaration
    map2 := make(map[string]string)

    // map declaration using make function boolean as value
    map3 := make(map[int]bool)

    // adding items to maps
    map1 = map[string]int{"one": 1, "two": 2, "three": 3};
    map2["name"] = "Rabindra"
    map3[1] = true;

    // access items from map
    val1 := map1["one"];
    val2 := map2["name"];
    val3 := map3[1];

    fmt.Println("value from map1: ", val1);
    fmt.Println("value from map2: ", val2);
    fmt.Println("value from map3: ", val3);

    // Iterate over map
    fmt.Println("\nIterating over map");
    for key, value := range map1 {
        fmt.Printf("%s: %d\n", key, value);
    }

    // delete items from map
    delete(map1, "two");
    delete(map2, "name");

    fmt.Println("map1 after deleted: ", map1);
    fmt.Println("map2 after deleted: ", map2);

    // truncate map3
    map1 = make(map[string]int);
    fmt.Println("map3 after truncate: ", map1);

    // merge two maps
    map4 := map[string]int{"four": 4, "five": 5};
    for key, value := range map4 {
        map1[key] = value;
    }
    fmt.Println("map1 after merging with map4: ", map1);

}
