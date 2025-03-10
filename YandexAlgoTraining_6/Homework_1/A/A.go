package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in, out := getBuffers()

	defer out.Flush()

	x1, y1 := ScanFloat64(in), ScanFloat64(in)
	x2, y2 := ScanFloat64(in), ScanFloat64(in)
	x, y := ScanFloat64(in), ScanFloat64(in)

	if y > y2 && x < x1 {
		out.WriteString("NW\n")
	} else if y > y2 && x > x2 {
		out.WriteString("NE\n")
	} else if y < y1 && x > x2 {
		out.WriteString("SE\n")
	} else if y < y1 && x < x1 {
		out.WriteString("SW\n")
	} else if y >= y2 {
		out.WriteString("N\n")
	} else if y <= y1 {
		out.WriteString("S\n")
	} else if x >= x2 {
		out.WriteString("E\n")
	} else {
		out.WriteString("W\n")
	}
}

func ScanFloat64(in *bufio.Reader) float64 {
	var a int
	line, _, _ := in.ReadLine()
	a, _ = strconv.Atoi(string(line))
	return float64(a)
}

func getBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}
