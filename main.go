package main

import (
	"fmt"
	// "strings"
	// "golang.org/x/text/cases"
	// "math/rand"
	// "strings"
	// "time"
)

// func main() {
// 	var firstName string = "john"
// 	lastName := "wick"

// 	fmt.Printf("Hello %s %s !!\n" , firstName, lastName)
// // fmt.Println("Hellow World!!!")
// }

// func main() {
// 	fruits := [4]string{"banana", "mango", "apple", "grape"}

// 	fmt.Println(len(fruits))
// 	fmt.Println(fruits)
// }

// func main() {
// 	numbers := [...]int{1, 12, 43, 14, 98, 101}

// 	fmt.Println(len(numbers))
// 	fmt.Println(numbers)
// }

// func main() {
// 	var char [...]string

// 	char[0] = "test"

// 	fmt.Println(char)
// }

// func main() {
// 	chars := []string{"lorem", "ipsum", "dolor", "sit", "amet"}

// 	chars = append(chars, "bogor")

// 	for _, char := range chars {
// 		fmt.Println("char : ", char)
// 	}
// }

// func main() {
// 	var chicken map[string]int
// 	chicken = map[string]int{}

// 	chicken["januari"] = 10
// 	chicken["februari"] = 18

// 	delete(chicken, "februari")

// 	fmt.Println("Januari : ", chicken["januari"])
// 	fmt.Println(chicken)

// 	var chickens = []map[string]string{
// 		{"name": "blue chick", "age": "1"},
// 		{"name": "red chick", "age": "1"},
// 		{"name": "green chick", "age": "1"},
// 	}

// 	for _, chicken := range chickens {
// 		fmt.Println(chicken["name"], chicken["age"])
// 	}
// }

// func main() {
// 	var names = []string{"John", "Wick"}

// 	printMessage("Hi", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(nameString)
// }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number: ", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number: ", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number: ", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int() % (max - min + 1) + min
// 	return value
// }

// func main() {
// 	var diameter float64 = 15
// 	var area, circumference = calculate(diameter)

// 	fmt.Printf("Luas: %.2f \n", area)
// 	fmt.Printf("Keliling: %2.f \n", circumference)
// }

// func calculate(d float64) (float64, float64) {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d / 2, 2)
// 	// hitung keliling
// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// predefined return value
// func calculate(d float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(d/2, 2)

// 	circumference = math.Pi * d

// 	return
// }

// variadic func
// func main() {
// 	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
// 	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)

// 	fmt.Printf(msg)
// }

// func calculate(numbers ...int) float64 {
// 	var total int = 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))

// 	return avg
// }

// variadic func with normal param
// func main() {
// 	yourHobbies("agam", "basket", "travel")
// }

// func yourHobbies(name string, hobbies ...string) {
// 	var hobbiesAsString = strings.Join(hobbies, ",")

// 	fmt.Printf("Hello my name is %s\n", name)
// 	fmt.Printf("My hobbies are : %s\n", hobbiesAsString)
// }

// func inside variable
// func main() {
// 	var getMinMax = func(n []int) (int, int) {
// 		var min, max int

// 		for i, e := range n {
// 			switch {
// 				case i == 0:
// 					max, min = e, e
// 				case e > max:
// 					max = e
// 				case e < min:
// 					min = e
// 			}
// 		}

// 		return min, max
// 	}

// 	var numbers = []int{2,3,4,5,6,7,8}
// 	var min, max = getMinMax(numbers)
// 	fmt.Printf("data: %v\nmin : %v\nmax : %v\n", numbers, min, max)
// }

// func as a param
// func main() {
// 	var data = []string{"wich", "john", "ethan"}
// 	var dataContainsO = filter(data, func(each string) bool {
// 		return strings.Contains(each, "o")
// 	})

// 	var dataLength5 = filter(data, func(each string) bool {
// 		return len(each) == 5
// 	})

// 	fmt.Println("data : ", data)
// 	fmt.Println("filter contain o : ", dataContainsO)
// 	fmt.Println("filter count char : ", dataLength5)
// }

// func filter(data []string, callback func(string) bool) []string {
// 	var result []string
// 	for _, each := range data {
// 		if filtered := callback(each); filtered {
// 			result = append(result, each)
// 		}
// 	}

// 	return result
// }

// pointer
// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA

// 	fmt.Println("numberA (value)  :", numberA)
// 	fmt.Println("numberA (address)  :", &numberA)
// 	fmt.Println("numberB (value)  :", *numberB)
// 	fmt.Println("numberB (address)  :", numberB)

// 	fmt.Println("")

// 	numberA = 5
// 	fmt.Println("numberA (value)  :", numberA)
// 	fmt.Println("numberA (address)  :", &numberA)
// 	fmt.Println("numberB (value)  :", *numberB)
// 	fmt.Println("numberB (address)  :", numberB)

// }

// parameter pointer

func main() {
	var number = 4
	fmt.Println("before : ", number)

	change(&number, 10)

	fmt.Println("after : ", number)
}

func change(original *int, value int) {

}
