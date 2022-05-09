package shorthash

func GenerateShortHashes(dictionary string, curLen int) (result []string) {
	if curLen == 0 || dictionary == "" {
		return []string{}
	}

	for _, v := range dictionary {
		result = append(result, string(v))
		for _, value := range GenerateShortHashes(dictionary, curLen-1) {
			result = append(result, string(v)+value)
		}
	}

	return result
}
