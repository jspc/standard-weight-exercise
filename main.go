package main

import (
	"log"
)

var (
	young = map[Gender]float64{
		Male:   0.9,
		Female: 0.85,
	}
	older = map[Gender]float64{
		Male:   0.95,
		Female: 0.9,
	}
)

func main() {
	i, err := GetInput()
	if err != nil {
		log.Panic(err)
	}

	std := standardWeight(i.AgeGroup, i.Height, i.Weight, i.Gender)
	log.Printf("Standard Weight: %+v\n", std)

	idx := standardWeightIndex(i.Weight, std)
	log.Printf("Index: %+v\n", idx)

	ty := classify(idx)
	log.Printf("Type: %q\n", ty)
}

func classify(idx float64) string {
	switch {
	case idx < 85:
		return "thin"

	case idx < 95:
		return "a bit thin"

	case idx < 115:
		return "standard"

	case idx < 125:
		return "a little fat"
	}

	return "fat"
}

func standardWeight(ageGroup AgeGroup, height, weight float64, gender Gender) float64 {
	multiplier := genderMultiplier(ageGroup, gender)

	return (height - 100) * multiplier
}

func genderMultiplier(ageGroup AgeGroup, gender Gender) float64 {
	if ageGroup == LTE35 {
		return young[gender]
	}

	return older[gender]
}

func standardWeightIndex(now, std float64) float64 {
	return now / std * 100
}
