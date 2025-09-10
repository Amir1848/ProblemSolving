package longestsubstringwithoutrepeatingcharacters

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	runeIndexesMap := map[rune][]int{}

	var maxLength int

	startIndex := 0
	for i, c := range s {

		indexes, found := runeIndexesMap[c]
		if found {
			newStartIndex := indexes[len(indexes)-1] + 1
			if newStartIndex > startIndex {
				startIndex = newStartIndex
			}
		}

		lengthOfSubString := i - startIndex + 1
		if lengthOfSubString > maxLength {
			maxLength = lengthOfSubString
		}

		indexes = append(indexes, i)
		runeIndexesMap[c] = indexes
	}

	return maxLength
}
