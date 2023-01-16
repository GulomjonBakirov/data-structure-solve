package array

func Generate(numRows int) [][]int {
	result := make([][]int, numRows+1)

	for i := 1; i < numRows+1; i++ {
		subArr := []int{}
		C := 1
		for k := 1; k < i+1; k++ {
			subArr = append(subArr, C)
			C = (C * (i - k) / k)

		}

		result[i] = subArr

	}
	result[0] = []int{1}

	return append(result[:1], result[2:]...)
}

func GetRow(rowIndex int) []int {

	result := Generate(rowIndex + 1)
	return result[rowIndex]
}
