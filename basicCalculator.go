package Problem_Solving

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	nb, str := sapare(s)
	if len(nb) == 1 {
		return nb[0]
	}
	fmt.Println(nb, str)
	n := nb[0]
	i, j := 1, 0
	for {
		if str[j] == "+" {
			r := nb[i]
			if j < len(str)-1 {
				k := j
				for k < len(str)-1 && (str[k+1] == "*" || str[k+1] == "/") {
					k++
					j++
					i++
					if str[k] == "*" {
						r *= nb[i]
					}
					if str[k] == "/" {
						r /= nb[i]
					}
				}
			}
			fmt.Println(r)
			n += r

		} else if str[j] == "-" {
			r := nb[i]
			if j < len(str)-1 {
				k := j
				for k < len(str)-1 && (str[k+1] == "*" || str[k+1] == "/") {
					k++
					j++
					i++
					if str[k] == "*" {
						r *= nb[i]
					}
					if str[k] == "/" {
						r /= nb[i]
					}
				}
			}
			n -= r
		} else if str[j] == "*" {
			n *= nb[i]
		} else if str[j] == "/" {
			n /= nb[i]
		}
		i++
		j++
		if i == len(nb) {
			break
		}
		if j == len(str) {
			break
		}
	}
	return n
}

func sapare(s string) ([]int, []string) {
	nb := []int{}
	str := []string{}
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '-' {
			k := i + 1
			count := 0
			if i < len(s)-1 {
				for k < len(s) && s[k] >= '0' && s[k] <= '9' {
					k++
					count++
				}
			}
			if k == i+1 {
				n, _ := strconv.Atoi(s[0 : i+2])
				nb = append(nb, n)
				i++
			} else {
				n, _ := strconv.Atoi(s[i:k])
				nb = append(nb, n)
				i += count
			}

		} else if s[i] >= '0' && s[i] <= '9' {
			k := i + 1
			count := 0
			if i < len(s)-1 {
				for k < len(s) && s[k] >= '0' && s[k] <= '9' {
					count++
					k++
				}
			}
			if k == i+1 {
				n, _ := strconv.Atoi(string(s[i]))
				nb = append(nb, n)
			} else {
				n, _ := strconv.Atoi(s[i:k])
				nb = append(nb, n)
				i += count
			}

		} else if s[i] != ' ' {
			str = append(str, string(s[i]))
		}
	}
	return nb, str
}
