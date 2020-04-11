// https://www.hackerrank.com/challenges/lisa-workbook/problem

package main

import "fmt"

func workbook(n int32, k int32, arr []int32) int32 {
	var pageNumber = int32(1)
	var specialCount = int32(0)

	for _, exCount := range arr {
		for startProb := int32(1); startProb <= exCount; startProb += k {
			var endProb = startProb + k - 1

			if endProb > exCount {
				endProb = exCount
			}

			if startProb <= pageNumber && pageNumber <= endProb {
				specialCount++
			}

			pageNumber++
		}
	}

	return specialCount
}

func main() {
	fmt.Println(workbook(5, 3, []int32{4, 2, 6, 1, 10,}))
}
