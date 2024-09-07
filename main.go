package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
	data, err := ReadData(arg[0])
	if err != nil {
		return
	}
	if len(data) <= 1 {
		return
	}
	mdn := Median(data)
	avg := Averge(data)
	vrnc := Vrariance(data, avg)
	std_dvt := math.Sqrt(vrnc)
	fmt.Printf("Average: %.0f\n", math.Round(avg))
	fmt.Printf("Median: %.0f\n", math.Round(mdn))
	fmt.Printf("Variance: %.0f\n", math.Round(vrnc))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(std_dvt))
}

func Averge(nums []int) float64 {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}

func Median(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	if len(nums)%2 == 0 {
		return float64(nums[n/2]+nums[n/2-1]) / 2.0
	} else {
		return float64(nums[n/2])
	}
}

func Vrariance(nums []int, averg float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += math.Pow(float64(num)-(averg), 2)
	}
	return sum / float64(len(nums))
}

func ReadData(file string) ([]int, error) {
	data, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var nums []int
	scan := bufio.NewScanner(data)
	for scan.Scan() {
		num, err := strconv.Atoi(scan.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}
