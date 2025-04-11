package main

import (
	"fmt"
	"sort"
)

func denseRanking(scores []int, gits []int) []int {

	rankMap := make(map[int]bool)
	for _, score := range scores {
		rankMap[score] = true
	}

	uniqueScores := make([]int, 0, len(rankMap))
	for score := range rankMap {
		uniqueScores = append(uniqueScores, score)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(uniqueScores)))

	result := make([]int, len(gits))
	n := len(uniqueScores)

	idx := n - 1
	for i, s := range gits {
		for idx >= 0 && s >= uniqueScores[idx] {
			idx--
		}
		result[i] = idx + 2 
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	var m int
	fmt.Scan(&m)

	gits := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&gits[i])
	}

	ranks := denseRanking(scores, gits)

	for _, r := range ranks {
		fmt.Printf("%d ", r)
	}
}
