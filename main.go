package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func rgb(i int) (int, int, int) {
	// from https://www.schemecolor.com/light-pink-pastels-gradient.php
	var lavenderBlush = [3]int{252, 237, 240}
	var piggyPink = [3]int{253, 226, 231}
	var palePink = [3]int{253, 215, 222}
	var lightRed = [3]int{254, 203, 212}
	var Pink = [3]int{254, 192, 203}
	var lightPink = [3]int{255, 181, 194}

	var colors = [...][3]int{lavenderBlush, piggyPink, palePink, lightRed, Pink, lightPink, Pink, lightRed, palePink, piggyPink, lavenderBlush}

	var factoredIndex = float64(i) * 0.01

	var baseColor = int(factoredIndex) % 9
	var nextColor = (baseColor + 1) % 9
	var gradient = factoredIndex - float64(int(factoredIndex))

	return int(float64(colors[nextColor][0])*gradient + float64(colors[baseColor][0])*(float64(1)-gradient)),
		int(float64(colors[nextColor][1])*gradient + float64(colors[baseColor][1])*(float64(1)-gradient)),
		int(float64(colors[nextColor][2])*gradient + float64(colors[baseColor][2])*(float64(1)-gradient))
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to pinkify with pipes.")
		fmt.Println("Usage: cat boring.txt | pinky")
	}

	reader := bufio.NewReader(os.Stdin)
	j := 0
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)
		j++
	}

}
