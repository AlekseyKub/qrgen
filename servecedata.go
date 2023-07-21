// servicedata
package qrgen

// import (
// 	"fmt"
// )

// Добавление служебной информации о маске и уровне коррекции
func CreateMaskAndCorrectionLevel(arr [][]int, ver, cor, number int) {
	data := maskAndCorrectionLevel[cor][number]

	createUpLeft(arr, data)
	createDownLeft(arr, data)
	createUpRight(arr, data)
	createBlackPoin(arr)
}

// добавление маски и уровня коррекции слева сверху
func createUpLeft(arr [][]int, data int) {
	x := 0
	for i := 0; i < 8; i++ {
		if i == 6 {
			x++
		}
		if HasBit(data, i) {
			arr[x][8] = 3
			x++
		} else {
			arr[x][8] = 2
			x++
		}
	}
	x = 0
	for i := 14; i > 7; i-- {
		if i == 8 {
			x++
		}
		if HasBit(data, i) {
			arr[8][x] = 3
			x++
		} else {
			arr[8][x] = 2
			x++
		}
	}
}

// Добавление снизу слева
func createDownLeft(arr [][]int, data int) {
	x := 14
	for i := len(arr) - 1; i > len(arr)-8; i-- {
		if HasBit(data, x) {
			arr[i][8] = 3
			x--
		} else {
			arr[i][8] = 2
			x--
		}
	}
}

// Добавление сверху справа
func createUpRight(arr [][]int, data int) {
	x := 0
	for i := len(arr) - 1; i > len(arr)-9; i-- {
		if HasBit(data, x) {
			arr[8][i] = 3
			x++
		} else {
			arr[8][i] = 2
			x++
		}
	}
}

// Добавить чернй модуль
func createBlackPoin(arr [][]int) {
	arr[len(arr)-8][8] = 3
}
func HasBit(n int, pos int) bool {
	val := n & (1 << pos)
	return (val > 0)
}
