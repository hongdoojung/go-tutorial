package main

func main() {
	count, total := sum(1, 7, 3, 5, 9)
	println(count, total)
}

func sum(nums ...int) (int, int) {
	result := 0
	count := 0
	for _, n := range nums {
		result += n
		count++
	}
	return count, result
}
