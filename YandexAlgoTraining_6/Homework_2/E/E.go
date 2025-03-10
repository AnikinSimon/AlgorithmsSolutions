package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in, out := getBuffers()

	defer out.Flush()

	n := ScanInt(in)

	arr := ScanIntSlice(n, in)

	sort.Slice(arr, func (i, j int) bool {
		return arr[i] < arr[j]
	})
	
	l := n / 2
	
	if (n % 2 == 0) {
		l--
	}
	
	r := l + 1

	for n > 0 {
		offset := n / 2
		if (n % 2 == 0) {
			offset--
		}
		if (offset <= l) {
			out.WriteString(fmt.Sprint(arr[l], " "))
			l--
		} else {
			out.WriteString(fmt.Sprint(arr[r], " "))
			r++
		}
		n--
	}
}

func ScanInt(in *bufio.Reader) int {
	var a int
	line, _, _ := in.ReadLine()
	a, _ = strconv.Atoi(string(line))
	return a
}

func ScanIntSlice(n int, in *bufio.Reader) []uint64 {
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
