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

	temp := ScanIntSlice(2, in)
	n, r := temp[0], temp[1]

	arr := ScanIntSlice(n, in)
	left, cnt := 0, 0

	var right uint64 = 1
	for ; right < n; right++ {
		for arr[right]-arr[left] > r {
			left++
		}
		cnt += left
	}

	out.WriteString(fmt.Sprint(cnt) + "\n")
}

func ScanInt(in *bufio.Reader) int {
	var a int
	line, _, _ := in.ReadLine()
	a, _ = strconv.Atoi(string(line))
	return a
}

func ScanIntSlice(n uint64, in *bufio.Reader) []uint64 {
	arr := make([]uint64, n)
	bytes, _, _ := in.ReadLine()
	nums := strings.Split(string(bytes), " ")

	for i := range arr {
		num, _ := strconv.Atoi(nums[i])
		arr[i] = uint64(num)
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
