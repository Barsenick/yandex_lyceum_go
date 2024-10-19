package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	var new map[string]string = make(map[string]string, len(m))
	for i, val := range m {
		new[val] = i
	}
	return new
}
