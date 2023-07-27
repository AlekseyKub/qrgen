// dataprocessing
package qrgen

import (
	"fmt"
)

// заполнение данными QR кода
func AddingDataQR(qr [][]int, data []byte, ver int) {
	direction := 1 //направление заполнения (1 вверх)
	sizeQr := SizeQrCode[ver]

	NumBitData := 0             //номер байт
	maxBitData := len(data) * 8 //номер испльзуемого байта

	for y := sizeQr - 1; y >= 0; y -= 2 {
		if y == 6 {
			y -= 1
		}

		if direction == 1 { //заполняем вверх
			for i := sizeQr - 1; i >= 0; i-- {
				if qr[i][y] == 0 {
					if NumBitData == maxBitData {
						break
					}
					qr[i][y] = hasNumBit(data, NumBitData)
					NumBitData++
				}
				if qr[i][y-1] == 0 {
					if NumBitData == maxBitData {
						break
					}
					qr[i][y-1] = hasNumBit(data, NumBitData)
					NumBitData++
				}
				if i == 0 {
					direction = 0
				}
			}
		} else {
			for i := 0; i <= sizeQr-1; i++ {
				if qr[i][y] == 0 {
					if NumBitData == maxBitData {
						break
					}
					qr[i][y] = hasNumBit(data, NumBitData)
					NumBitData++
				}
				if qr[i][y-1] == 0 {
					if NumBitData == maxBitData {
						break
					}
					qr[i][y-1] = hasNumBit(data, NumBitData)
					NumBitData++
				}
				if i == sizeQr-1 {
					direction = 1
				}
			}
		}
	}
}

// расчетблоков коррекции
func AddingDataOut(dataout []byte, data []byte, ver int, cor int) []byte {

	numberBlock := NumberOfBlocks[ver][cor]                     //количество блоков данных
	byteOnBlock := NumberOfBytesCorrection[ver][cor]            //количество байт коррекции на каждый блок
	byteOnBlockData := (MaxDataBit[ver][cor] / numberBlock) / 8 //кол. байт в блоке данных
	extendedBlocks := (MaxDataBit[ver][cor] / 8) % numberBlock  //количество дополненных блоков данных

	// fmt.Print("количество блоков данных = ")
	// fmt.Println(numberBlock)
	// fmt.Print("количество байт коррекции = ")
	// fmt.Println(byteOnBlock)
	// fmt.Print("кол. байт в блоке данных = ")
	// fmt.Println(byteOnBlockData)
	// fmt.Print("дополненные блоки = ")
	// fmt.Println(extendedBlocks)

	blockCorrSlise := make([]byte, 0, numberBlock*byteOnBlock)
	numExtendedBlock := 0 //номер дополненного блока

	a := 0
	b := 0
	for i := 0; i < numberBlock; i++ { // 4 блока 0-3   2 доп.блока 2-3
		if (extendedBlocks > 0) && (i >= numberBlock-extendedBlocks) {
			a = (i * byteOnBlockData) + numExtendedBlock
			b = ((i + 1) * byteOnBlockData) + (numExtendedBlock + 1)
			numExtendedBlock++
		} else {
			a = i * byteOnBlockData
			b = ((i + 1) * byteOnBlockData)
		}

		data1 := data[a:b]
		blockCorrSlise = append(blockCorrSlise, CreateBlock(data1, byteOnBlock)...)
	}

	//добавление данных по блокам к выходному масссибу
	numExtendedBlock = 0
	for i := 0; i < byteOnBlockData; i++ {
		x := i                             //номер байта из блока
		for y := 0; y < numberBlock; y++ { //номер блока
			if y > (numberBlock - extendedBlocks) {
				x += 1
				numExtendedBlock++
			}
			dataout = append(dataout, data[x])
			x += byteOnBlockData
		}
	}

	if extendedBlocks > 0 {
		numExtendedBlock = 0
		y := ((numberBlock - extendedBlocks) + 1) * (byteOnBlockData)
		for i := 0; i < extendedBlocks; i++ {
			dataout = append(dataout, data[y])
			y += byteOnBlockData + 1
		}
	}

	//добавление блоков корррекции по блокам к выходному масссибу
	for i := 0; i < byteOnBlock; i++ {
		x := i //
		for y := 0; y < numberBlock; y++ {
			dataout = append(dataout, blockCorrSlise[x])
			x += byteOnBlock
		}
	}
	return dataout
}

// Расчет блока коррекции
// Принимает блок для которого необходимо посчитать уровень коррекции
//
//	и количество байт в блоке уровня коррекции (возвращает блок коррекции)
func CreateBlock(block []byte, byteOnBlock int) []byte {

	blockCorr := make([]byte, 0, len(block))    //создаем срез по количеству байт
	blockZero := make([]byte, byteOnBlock)      //создаем срез и "0"
	blockCorr = append(blockCorr, block...)     //в начало среза вставляем исходный блок
	blockCorr = append(blockCorr, blockZero...) //дополняем нолями

	fistByte := 0
	polynom := genPolynoms[byteOnBlock-1]

	for i := 0; i < len(block); i++ {
		if blockCorr[0] == 0 {
			blockCorr = blockCorr[1:len(blockCorr)]
			blockCorr = append(blockCorr, 0)
		} else {
			fistByte = galoisFieldsReverse[blockCorr[0]]
			blockCorr = blockCorr[1:len(blockCorr)]
			blockCorr = append(blockCorr, 0)

			for a := 0; a < byteOnBlock; a++ {
				x := polynom[a] + fistByte
				if x > 254 {
					x = x % 255
				}
				x = galoisFields[x]
				blockCorr[a] ^= byte(x)
			}
		}
	}
	blockCorr = blockCorr[:byteOnBlock]
	return blockCorr
}

// Проверка на то помещаются ли данные в QR выбранной версии
func SizeСheck(data []byte, ver, cor int) bool {
	x := len(data) * 8
	return (x < MaxDataBit[ver][cor])
}

// Распечатать по битам
func printBit(x int) {
	for i := 7; i >= 0; i-- {
		if hasBit(x, i) {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}

// Проверяет 1 или 0 в проверяемом разряде.
func hasBit(n int, pos int) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func hasNumBit(data []byte, num int) int {
	n := 0
	n = int(data[num/8])
	val := 0
	switch num % 8 {
	case 0:
		val = n & (1 << 7)
	case 1:
		val = n & (1 << 6)
	case 2:
		val = n & (1 << 5)
	case 3:
		val = n & (1 << 4)
	case 4:
		val = n & (1 << 3)
	case 5:
		val = n & (1 << 2)
	case 6:
		val = n & (1 << 1)
	case 7:
		val = n & (1 << 0)
	}
	if val > 0 {
		return 1
	} else {
		return 0
	}
}
