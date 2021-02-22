/*Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false*/

package Easy

import (
	"reflect"
)

func IsAnagram(s string, t string) bool {
	sMap := make(map[string]int, 0)
	tMap := make(map[string]int, 0)
	for i := 0; i < len(s); i++ {
		sMap[string(s[i])] = sMap[string(s[i])] + 1
	}
	for i := 0; i < len(t); i++ {
		tMap[string(t[i])] = tMap[string(t[i])] + 1
	}

	return reflect.DeepEqual(sMap, tMap)
}
