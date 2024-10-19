package main

func DeleteLongKeys(m map[string]int) map[string]int {
	var new map[string]int = make(map[string]int, len(m))

	for i := range m {
		if len(i) < 6 {
			new[i] = m[i]
		}
	}

	return new
}
