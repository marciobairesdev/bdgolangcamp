package main

func filterByAgeRange(fromAge, toAge int, ages []int) []int {
	response := make([]int, 0)
	for _, ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge {
			response = append(response, ageToCheck)
		}
	}
	return response
}
