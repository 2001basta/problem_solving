package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	str_data := strings.Fields(string(data))

	fmt.Println(variance(str_data))
}

func variance(str []string) (int, int, int) {
	s := []float64{}
	for _, i := range str {
		x, er := strconv.Atoi(i)
		if er != nil {
			log.Fatal(er)
		}
		s = append(s, float64(x))
	}
	x := len(s)
	var sum float64
	for _, i := range s {
		sum += i
	}
	averge := sum / float64(x)
	var variance float64
	for _,i := range s {
		variance += math.Pow((i-averge), 2) / float64(x)
	}
	standr := check(math.Sqrt(variance))
	xn := check(variance)
	y := check(averge)
	return xn, y , standr
}
func check(n float64) int{
	if int(n*10)%10 >= 5{
		return  int(n+1)
	}
	return int(n)
}