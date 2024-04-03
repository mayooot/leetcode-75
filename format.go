package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var algorithm = flag.String("a", "", "")

func main() {
	flag.Parse()

	res := convert(*algorithm)
	fmt.Println(res)
}

func convert(origin string) string {
	numberStr := strings.Split(origin, ".")[0]
	name := strings.TrimSpace(strings.Split(origin, ".")[1])

	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if number < 10 {
		numberStr = fmt.Sprintf("%d%d", 0, number)
	}

	name = strings.Replace(name, " ", "", -1)

	return fmt.Sprintf("%s-%s", numberStr, name)
}
