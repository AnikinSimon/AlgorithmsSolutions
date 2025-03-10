package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	in, out := getBuffers()

	defer out.Flush()

	temp := ScanIntSlice(2, in)
	n, k := temp[0], temp[1]

	arr := ScanIntSlice(n, in)
	l, sm, cnt := -1, 0, 0

	for r := range arr {
		sm += arr[r]
		for sm > k {
			l++
			sm -= arr[l]
		}
		if sm == k {
			cnt++
		}
	}

	out.WriteString(fmt.Sprint(cnt) + "\n")
}

func ScanInt(in *bufio.Reader) int {
	var a int
	line, _, _ := in.ReadLine()
	a, _ = strconv.Atoi(string(line))
	return a
}

func ScanIntSlice(n int, in *bufio.Reader) []int {
	arr := make([]int, n)
	bytes, _, _ := in.ReadLine()
	nums := strings.Split(string(bytes), " ")

	for i := range arr {
		num, _ := strconv.Atoi(nums[i])
		arr[i] = num
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
