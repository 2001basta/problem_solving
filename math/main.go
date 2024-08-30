package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERORR INPUT")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	str_data := strings.Fields(string(data))
	if len(str_data) == 0 {
		fmt.Println("ERORR FILLE")
		return
	}
	s := []int{}
	for _, i := range str_data {
		x, er := strconv.Atoi(i)
		if er != nil {
			log.Fatal(er)
		}
		s = append(s, x)
	}
	sort.Ints(s)
	averge, median, variance, standard_deviation := math_skil(s)
	fmt.Println("Average:", averge)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", standard_deviation)
}

func math_skil(ar_nb []int) (int, int, int, int) {
	x := len(ar_nb)
	var m float64
	if x%2 == 0 {
		m = float64(ar_nb[x/2])
	} else {
		m = float64(ar_nb[x/2] + ar_nb[x/2-1])
	}
	median := int(math.Round(m))
	var sum int
	for _, i := range ar_nb {
		sum += i
	}
	averge := float64(sum) / float64(x)

	var varian float64
	for _, i := range ar_nb {
		varian += math.Pow((float64(i) - averge), 2)
	}
	varian /= float64(x)
	variance := int(math.Round(varian))
	standard_deviation := int(math.Round(math.Sqrt(varian)))

	return int(math.Round(averge)), median, variance, standard_deviation
}
