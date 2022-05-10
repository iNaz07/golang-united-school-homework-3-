package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		result = append(result, input[k])
	}
	return
}
