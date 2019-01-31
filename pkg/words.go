package words

func WordCount(s string) map[string]int { // HL
	m := map[string]int{}

	var word string
	maxlen := len(s)
	for i := 0; i < maxlen; i++ {
		char := string(s[i])
		if char != " " {
			word += char
		} else {
			m[word]++
			word = ""
		}
	}
	if word != "" {
		m[word]++
	}

	return m
} // HL
