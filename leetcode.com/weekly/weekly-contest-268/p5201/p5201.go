package p5201

func wateringPlants(plants []int, capacity int) (stepCount int) {
	capacityNow := capacity
	pPos := 0
	stepCount = 0
	n := len(plants)
	for pPos < n {
		// fmt.Printf("pPos:%+v\n", pPos)
		//打水浇水
		if capacityNow < plants[pPos] {
			capacityNow = capacity - plants[pPos]
			stepCount += pPos*2 + 1
			pPos++
			// fmt.Printf("1,stepCount:%+v,pPos:%+v\n", stepCount, pPos)
			continue
		}
		//浇水
		capacityNow -= plants[pPos]
		stepCount++
		pPos++
		// fmt.Printf("2,stepCount:%+v,pPos:%+v\n", stepCount, pPos)
	}
	return stepCount
}
