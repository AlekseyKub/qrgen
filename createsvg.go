// createSvg
package qrgen

import (
	"fmt"
	"os"
)

func CreateSvg(arr [][]int, name string) {

	sizePoint := 10
	border := 40
	sizeBakground := (len(arr))*sizePoint + border*2
	x := border
	y := border

	startSvg := fmt.Sprintf(`<svg version="1.1" baseProfile="full" width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, sizeBakground, sizeBakground)
	endSvg := `</svg>`

	file, err := os.Create(name)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(startSvg + "\n")

	//Максимальное значение X
	maxX := ((len(arr))*sizePoint + border) - sizePoint

	//отрисовка массива в SVG
	for _, row := range arr {
		for _, poz := range row {
			if (poz == 1) || (poz == 3) {
				rect := fmt.Sprintf(`<rect x="%d" y="%d" width="10" height="10" fill="black"/>`, x, y)
				file.WriteString(rect + "\n")
			}
			if x == maxX {
				x = border
				y += sizePoint
				break
			}
			x += sizePoint
		}
	}

	file.WriteString(endSvg)

}
