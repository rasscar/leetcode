func braceExpansionII(expression string) []string {
	index := 0
	resultMap := parse(expression, &index)

	result := make([]string, 0, len(resultMap))
	for k := range resultMap {
		result = append(result, k)
	}

	sort.Strings(result)
	return result
}

func parse(s string, index *int) map[string]bool {
	resultSet := make(map[string]bool)
	currentSet := make(map[string]bool)

	for *index < len(s) {
		char := s[*index]

		if char == '{' {
			*index++
			subSet := parse(s, index)
			currentSet = crossProduct(currentSet, subSet)
		} else if char == ',' {
			*index++
			for k := range currentSet {
				resultSet[k] = true
			}
			currentSet = make(map[string]bool)
		} else if char == '}' {
			*index++
			for k := range currentSet {
				resultSet[k] = true
			}
			return resultSet
		} else {
			letterSet := map[string]bool{string(char): true}
			currentSet = crossProduct(currentSet, letterSet)
			*index++
		}
	}

	for k := range currentSet {
		resultSet[k] = true
	}
	return resultSet
}

func crossProduct(setA, setB map[string]bool) map[string]bool {
	if len(setA) == 0 {
		return setB
	}
	if len(setB) == 0 {
		return setA
	}

	product := make(map[string]bool)
	for a := range setA {
		for b := range setB {
			product[a+b] = true
		}
	}
	return product
}
