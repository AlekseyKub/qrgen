// createSvg
package qrgen

import (
	"fmt"
	"os"
)

func CreateSvgRect(arr [][]int, name string) {

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
	for i, row := range arr {
		for a, poz := range row {
			if (poz == 1) || (poz == 3) && !(((i < 7) && (a < 7)) || (i < 7) && (a > len(arr)-8) || ((i > len(arr)-8) && (a < 7))) {
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

	//Отрисовка поисковых узоров
	x = border
	y = border
	//левый верхний
	rect := fmt.Sprintf(`<rect x="%d" y="%d" width="70" height="70" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")
	x += sizePoint
	y += sizePoint
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="50" height="50" fill="white"/>`, x, y)
	file.WriteString(rect + "\n")
	x += sizePoint
	y += sizePoint
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="30" height="30" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")

	//правый верхний
	x = ((len(arr))-7)*sizePoint + border
	y = border
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="70" height="70" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")
	x += sizePoint
	y += sizePoint
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="50" height="50" fill="white"/>`, x, y)
	file.WriteString(rect + "\n")
	x += sizePoint
	y += sizePoint
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="30" height="30" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")

	//нижний левый
	x = border
	y = ((len(arr))-7)*sizePoint + border
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="70" height="70" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")
	x += sizePoint
	y += sizePoint
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="50" height="50" fill="white"/>`, x, y)
	file.WriteString(rect + "\n")
	x += sizePoint
	y += sizePoint
	rect = fmt.Sprintf(`<rect x="%d" y="%d" width="30" height="30" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")

	file.WriteString(endSvg)

}
func CreateSvgCircle(arr [][]int, name string) {

	sizePoint := 10
	border := 40
	sizeBakground := (len(arr))*sizePoint + border*2
	x := border + 5
	y := border + 5

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
	maxX := ((len(arr))*sizePoint + border) - sizePoint - 5

	//отрисовка массива в SVG
	for i, row := range arr {
		for a, poz := range row {
			if (poz == 1) || (poz == 3) && !(((i < 7) && (a < 7)) || (i < 7) && (a > len(arr)-8) || ((i > len(arr)-8) && (a < 7))) {
				rect := fmt.Sprintf(`<circle cx="%d" cy="%d" r="5" fill="black"/>`, x, y)
				file.WriteString(rect + "\n")
			}
			if x == maxX {
				x = border + 5
				y += sizePoint
				break
			}
			x += sizePoint
		}
	}

	//Отрисовка поисковых узоров
	x = border + 35
	y = border + 35
	//левый верхний
	rect := fmt.Sprintf(`<circle cx="%d" cy="%d" r="35" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="25" fill="white"/>`, x, y)
	file.WriteString(rect + "\n")
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="15" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")

	//правый верхний
	x = (((len(arr))-7)*sizePoint + border) + 35
	y = border + 35
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="35" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="25" fill="white"/>`, x, y)
	file.WriteString(rect + "\n")
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="15" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")

	//нижний левый
	x = border + 35
	y = (((len(arr))-7)*sizePoint + border) + 35
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="35" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="25" fill="white"/>`, x, y)
	file.WriteString(rect + "\n")
	rect = fmt.Sprintf(`<circle cx="%d" cy="%d" r="15" fill="black"/>`, x, y)
	file.WriteString(rect + "\n")

	file.WriteString(endSvg)

}
