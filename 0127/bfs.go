package leetcode

/*
这个版本有错误！ 看 bfs1.go
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dictm := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		dictm[wordList[i]] = true
	}
	if dictm[endWord] == false {
		return 0
	}

	if dictm[beginWord] == false {
		wordList = append([]string{beginWord}, wordList...)
	}
	n := len(wordList)
	numsList := make([]int, len(wordList))
	matrix := make([][]bool, n)
	visited := make([]bool, n)
	distance := make([]int, n)
	endwordindex := 0
	startwordindex := 0

	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}

	for i := 0; i < len(wordList); i++ {
		num := 0
		for k := 0; k < len(wordList[i]); k++ {
			num |= 1 << uint(wordList[i][k]-'a'+1) //错误： 这里，根本没有考虑， letter 的先后顺序，所以这种方法是无效的
		}
		numsList[i] = num
		if wordList[i] == endWord {
			endwordindex = i
		}
		if wordList[i] == beginWord {
			startwordindex = i
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// check numsList[i] 和 numsList[j] 的编辑距离是不是差 1
			tmp := numsList[i] & numsList[j]
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

	visited[startwordindex] = true
	stack = append(stack, startwordindex)
	bfs()

	if visited[endwordindex] == false {
		return 0
	}
	return distance[endwordindex] + 1
}

func countones(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n &= n - 1
	}
	return cnt
}

func diffone(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count == 1
}
