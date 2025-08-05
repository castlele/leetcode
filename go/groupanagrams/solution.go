package groupanagrams

import (
	"sort"
)

// source: https://leetcode.com/problems/group-anagrams
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	sort := func(str string) string {
		bytes := []byte(str)
		sort.Sort(sortedStr(bytes))

		return string(bytes)
	}

	for _, str := range strs {
		sorted := sort(str)

		if group, ok := groups[sorted]; ok {
			group = append(group, str)
			groups[sorted] = group
		} else {
			groups[sorted] = []string{str}
		}
	}

	var res [][]string

	for _, group := range groups {
		res = append(res, group)
	}

	return res
}

type sortedStr []byte

func (s sortedStr) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortedStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedStr) Len() int {
	return len(s)
}
