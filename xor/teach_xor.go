package xor

var network [][]float64
var trainingData map[float64]float64

func getTrainingData(count int) {
	trainingData = make(map[float64]float64, count)
	for i := 0; i < count; i++ {

	}
}

// TODO - create interfaces for functions like train, getTrainingData to adhere to
func train() {
	
}

func inputOutputToKeyValue(x1, x2, y bool) {
	if x1 && x2 {
		trainingData[3] = 1
		return
	}
	if x1 && !x2 {
		trainingData[2] = 1
		return
	}
	if !x1 && x2 {
		trainingData[1] = 1
		return
	}
	trainingData[0] = 0
}
