package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, out := getBuffers()

	defer out.Flush()

	n := ScanInt(in)

	arr := ScanIntSlice(n, in)

	var prev int64 = 0

	for i := range arr {
		out.WriteString(fmt.Sprintf("%d ", prev+arr[i]))
		prev += arr[i]
	}
}

func ScanInt(in *bufio.Reader) int {
	var a int
	line, _, _ := in.ReadLine()
	a, _ = strconv.Atoi(string(line))
	return a
}

func ScanIntSlice(n int, in *bufio.Reader) []int64 {
	arr := make([]int64, n)
	bytes, _, _ := in.ReadLine()
	nums := strings.Split(string(bytes), " ")

	for i := range arr {
		num, _ := strconv.Atoi(nums[i])
		arr[i] = int64(num)
	}
	return arr
}

func getBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}
