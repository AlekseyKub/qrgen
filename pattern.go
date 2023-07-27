// pattern
package qrgen

// Создание поисковых узоров
func MakeSearchPattern(arr [][]int) {
	pattern := [7][7]int{{3, 3, 3, 3, 3, 3, 3},
		{3, 2, 2, 2, 2, 2, 3},
		{3, 2, 3, 3, 3, 2, 3},
		{3, 2, 3, 3, 3, 2, 3},
		{3, 2, 3, 3, 3, 2, 3},
		{3, 2, 2, 2, 2, 2, 3},
		{3, 3, 3, 3, 3, 3, 3}}
	//Верхний левый
	for i := 0; i < 7; i++ {
		for y := 0; y < 7; y++ {
			arr[i][y] = pattern[i][y]
		}
	}
	//Верхний правый
	for i := 0; i < 7; i++ {
		x := (len(arr)) - 7
		for y := 0; y < 7; y++ {
			arr[i][x] = pattern[i][y]
			x++
		}
	}
	//Нихний левый
	z := 0
	for x := (len(arr)) - 7; x < (len(arr)); x++ {
		for y := 0; y < 7; y++ {
			arr[x][y] = pattern[z][y]
		}
		z++
	}

	// Сознание границы у поисковых узоров
	for i := 0; i < 8; i++ {
		arr[7][i] = 2
	}
	for i := (len(arr)) - 8; i < (len(arr)); i++ {
		arr[7][i] = 2
	}

	for i := 0; i < 7; i++ {
		arr[i][7] = 2
	}

	for i := (len(arr)) - 8; i < (len(arr)); i++ {
		arr[i][7] = 2
	}

	for i := 0; i < 8; i++ {
		arr[(len(arr))-8][i] = 2
	}

	for i := 0; i < 7; i++ {
		arr[i][(len(arr))-8] = 2
	}
}

// Создание выравнивающих узоров
func CreatePattern(coordinates []int, arr [][]int) {
	patternArr := [][]int{{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 3, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3}}

	//Перебираем все координаты исключая те которые попадают на поисковые узоры
	for _, x := range coordinates {
		for _, y := range coordinates {
			if !((x == 6 && y == 6) || (x == 6 && y > (len(arr))-10) || (x > (len(arr))-10 && y == 6)) {
				create(x, y, arr, patternArr)
			}
		}
	}
}

// Создает поисковый узор в нужных координатах
func create(i, z int, arr [][]int, arr2 [][]int) {
	i = i - 2
	z = z - 2
	r := z
	for x := 0; x < (len(arr2)); x++ {
		for y := 0; y < (len(arr2)); y++ {
			arr[i][z] = arr2[x][y]
			z++
		}
		i++
		z = r
	}
}

// Воздание Вертикальной и горизонтальной последовательности 1 и 0
func СreateTimingTemplate(arr [][]int) {
	row := 6
	color := 3
	for i := 8; i < (len(arr))-8; i++ {
		if color == 2 {
			arr[row][i] = color
			arr[i][row] = color
			color = 3
		} else {
			arr[row][i] = color
			arr[i][row] = color
			color = 2
		}
	}
}
