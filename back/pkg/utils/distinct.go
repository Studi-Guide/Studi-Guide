package utils

//Distinct removes doubled values from string arrays
func Distinct(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
