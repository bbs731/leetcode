package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dictm := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		dictm[wordList[i]] = true
	}
	if dictm[endWord] == false {
		return 0
	}

	wordList = wordList[:0]
	wordList = append(wordList, beginWord)

	for k := range dictm {
		if k != beginWord && k != endWord {
			wordList = append(wordList, k)
		}
	}
	wordList = append(wordList, endWord)

	n := len(wordList)
	matrix := make([][]bool, n)
	visited := make([]bool, n)
	distance := make([]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// check numsList[i] 和 numsList[j] 的编辑距离是不是差 1
			if diffone(wordList[i], wordList[j]) {
				matrix[i][j] = true
				matrix[j][i] = true
			}
		}
	}

	var bfs func()
	stack := make([]int, 0)
	bfs = func() {
		for len(stack) != 0 {
			top := stack[0]
			stack = stack[1:]

			for i := 0; i < n; i++ {
				if i != top && visited[i] == false && matrix[top][i] == true {
					visited[i] = true
					stack = append(stack, i)
					distance[i] = distance[top] + 1
				}
			}
		}
	}
	visited[0] = true
	distance[0] = 1
	stack = append(stack, 0)
	bfs()

	return distance[n-1]
}

/*
这个 diffone 的发放， 正常是过不去时间限制的！

创建图，有技巧，！ 看官网的答案！
https://leetcode-cn.com/problems/word-ladder/solution/dan-ci-jie-long-by-leetcode-solution/

*/
func diffone(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count == 1
}
