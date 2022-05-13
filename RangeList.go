/**
	RangeList -- 区间列表

	Author: 顾君垚 <1205072443@qq.com>
	blog：hdugujy.xyz
	Create on 2022/5/13 12:10.
**/
package main

import "fmt"


type RangeList struct {
	//TODO: implement
	invals [][2]int
}

func (this *RangeList) Add(rangeElement [2]int) {
	left := rangeElement[0]
	right := rangeElement[1]
	n := len(this.invals)
	tmp := make([][2]int, 0)
	//遍历添加
	for i := 0; i <= n; i++ {
		if i == n || this.invals[i][0] > right {
			tmp = append(tmp, [2]int{left, right})
			for i < n {
				tmp = append(tmp, this.invals[i])
				i++
			}
		} else if this.invals[i][1] < left {
			tmp = append(tmp, this.invals[i])
		} else {
			left = min(left, this.invals[i][0])
			right = max(right, this.invals[i][1])
		}
	}
	this.invals = make([][2]int, len(tmp))
	copy(this.invals, tmp)
}

func (this *RangeList) Remove(rangeElement [2]int) {
	//确定左右端点
	left := rangeElement[0]
	right := rangeElement[1]
	tmp := make([][2]int, 0)
	n := len(this.invals)
	if (n == 0) {
		fmt.Println("error：RangeList中没有元素可以移除！")
		return
	}
	//遍历
	for i := 0; i < n; i++ {
		if this.invals[i][1] <= left || this.invals[i][0] >= right {
			tmp = append(tmp, this.invals[i])
		} else {
			if this.invals[i][0] < left {
				tmp = append(tmp, [2]int{this.invals[i][0], left})
			}
			if this.invals[i][1] > right {
				tmp = append(tmp, [2]int{right, this.invals[i][1]})
			}
		}
	}
	this.invals = make([][2]int, len(tmp))
	//由tmp复制到this.invals
	copy(this.invals, tmp)
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (this *RangeList) Print() {
	//打印列表元素
	n := len(this.invals)
	if (n == 0) {
		fmt.Println("RangeList中无元素")
	}
	for i := 0; i < n; i++ {
		if (i < n-1) {
			fmt.Print("[ ", this.invals[i][0], " , ", this.invals[i][1], " ) ")
		} else {
			fmt.Println("[", this.invals[i][0], ",", this.invals[i][1], ")")
		}

	}
}

func main() {
	rl := RangeList{}
	rl.Add([2]int{1, 5})
	rl.Print()
	// Should display: [1, 5)

	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)

	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)

	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)

	rl.Add([2]int{2, 4})
	rl.Print()
	// Should display: [1, 5) [10, 21)

	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)

	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)

	rl.Remove([2]int{10, 11})
	rl.Print()
	// Should display: [1, 8) [11, 21)

	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)

	rl.Remove([2]int{3, 19})
	rl.Print()
	// Should display: [1, 3) [19, 21)

	//无元素情况下的测试用例
	rl.Remove([2]int{1, 21})
	rl.Print()

	rl.Remove([2]int{1, 21})
	rl.Print()
}
