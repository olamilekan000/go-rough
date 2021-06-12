package helpers

func Remove(items []string, itemIndex int) []string {
	newitems := []string{}

	for index, val := range items {
		if index != itemIndex {
			newitems = append(newitems, val)
		}
	}

	return newitems
}

func AddTasks(tasks *[]string, val string) {
	*tasks = append(*tasks, val)
}
