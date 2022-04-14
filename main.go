package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Boardはナンプレのboardを指す
type Board [9][9]int

func pretty(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i % 3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func duplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Board) bool {
	// 行チェック
	for i := 0; i < 9; i++ {
		// 出現頻度
		var c[10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}

	//列チェック
	for i := 0; i < 9; i++ {
		// 出現頻度
		var c[10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if duplicated(c) {
			return false
		}
	}

	//3x3のチェック
	for i := 0; i < 9; i+=3 {
		for j := 0; j < 9; j+=3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}
	return true
}

func solved(b Board) bool {
	if !verify(b) {
		return false
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board) bool {
	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// マスが0の場合
			if b[i][j] == 0 {
				for c := 9; c >= 1; c-- {
					b[i][j] = c
					// もし数がルール通りなら
					if verify(*b) {
						// 更に探索する
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false 
			}
		}
	}
	return false
}

func main() {
	b := Board{
		{0, 5, 0, 0, 8, 3, 0, 1, 7},
		{0, 0, 0, 1, 0, 0, 4, 0, 0},
		{3, 0, 4, 0, 0, 5, 6, 0, 8},
		{0, 0, 0, 0, 3, 0, 0, 0, 9},
		{0, 9, 0, 8, 2, 4, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 7, 0},
		{0, 0, 9, 0, 0, 0, 0, 5, 0},
		{0, 0, 7, 2, 9, 0, 0, 8, 6},
		{1, 0, 3, 6, 0, 7, 2, 0, 4},
	}
	fmt.Printf("%+v\n", pretty(b))
}