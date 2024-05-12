package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	m := wordCount(s)

	slice, key := sortDescByCount(m)

	result := sortDescAlphabetically(slice, key)

	if len(result) >= 10 {
		return result[0:10]
	}

	return make([]string, 0)
}

func wordCount(s string) map[string]int {
	s += " "

	m := make(map[string]int)

	var builder strings.Builder

	for _, val := range s {
		if val == ' ' || val == '\n' || string(val) == "\t" {
			if _, ok := m[builder.String()]; ok {
				m[builder.String()]++
			} else {
				m[builder.String()] = 1
			}

			builder.Reset()
		} else {
			builder.WriteRune(val)
		}
	}

	return m
}

func sortDescByCount(m map[string]int) ([]string, []int) {
	i := make([]int, len(m))
	v := make([]string, len(m))

	for {
		if len(m)-1 == -1 {
			break
		}

		delete(m, v[len(m)-1])

		for mk, mv := range m {
			if i[len(m)-1] == 0 || mv < i[len(m)-1] {
				i[len(m)-1] = mv
				v[len(m)-1] = mk
			}
		}
	}

	return v, i
}

func sortDescAlphabetically(str []string, key []int) []string {
	slicesStartKey := 0
	previousValue := 0

	for k, v := range key {
		if k > 0 && v != previousValue {
			sort.Strings(str[slicesStartKey:k])
			slicesStartKey = k
		}

		if k == len(key)-1 && v == previousValue {
			sort.Strings(str[slicesStartKey : k+1])
		}

		previousValue = v
	}

	return str
}
