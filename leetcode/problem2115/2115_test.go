package problem2115

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	indeg := make(map[string]int)
	preq := make(map[string]map[string]struct{})
	for i, recipe := range recipes {
		for _, ingredient := range ingredients[i] {
			if preq[ingredient] == nil {
				preq[ingredient] = make(map[string]struct{})
			}
			preq[ingredient][recipe] = struct{}{}
			indeg[recipe]++
		}
	}

	queue := make([]string, 0)
	for _, supply := range supplies {
		queue = append(queue, supply)
	}

	for len(queue) != 0 {
		supply := queue[0]
		queue = queue[1:]
		for k, _ := range preq[supply] {
			indeg[k]--
			if indeg[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	result := make([]string, 0)
	for _, recipe := range recipes {
		if indeg[recipe] == 0 {
			result = append(result, recipe)
		}
	}
	return result
}
