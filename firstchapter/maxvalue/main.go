package main

func main() {
	num := []int{1, 2, 5, 7, 8, 9}
	max := num[0]
	for _, v := range num[1:] {
		if max < v {
			max = v
		}
	}
	println(max)
}
