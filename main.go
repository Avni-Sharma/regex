package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "${status.dbCredentials.user}:${status.dbCredentials.password}@${status.dbConnectionIP}"
	//str1 := "this is a [sample] [[string]] with [SOME] special words"

	//re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	re := regexp.MustCompile(`\$\{.*?\}`)
	fmt.Printf("Pattern: %v\n", re.String())      // print pattern
	fmt.Println("Matched:", re.MatchString(str1)) // true

	fmt.Println("\nText between square brackets:")
	submatchall := re.FindAllString(str1, -1)
	fmt.Println(submatchall)

	for i, element := range submatchall {
		element = strings.Trim(element, "${")
		element = strings.Trim(element, "}")
		fmt.Println(element)
		fmt.Println(i)
		str1 = strings.Replace(str1, element, "parsedvalue", i+1)
		// instead of parsed value there will be a function call
		fmt.Println("RESULT", str1)
	}
	fmt.Println("RESULT", str1)

}
