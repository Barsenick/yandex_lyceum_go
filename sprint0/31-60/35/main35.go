package main

func CountingSort(contacts []string) map[string]int {
	occurrences := make(map[string]int)

	for _, num := range contacts {
		occurrences[num]++
	}

	return occurrences
}
