
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455954456383488>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestF(t *testing.T) {
	vR := NewTreeVertex("r", nil)
	vR.Threshold = 2

	vD := vR.AddDescendant(NewTreeVertex("D", nil))
	vD.Threshold = 2
	for _, id := range []string{"A", "B", "C"} {
		vD.AddDescendant(NewTreeVertex(id, nil))
	}

	vE := vR.AddDescendant(NewTreeVertex("E", nil))
	vE.Threshold = 2
	for _, id := range []string{"a", "b", "c"} {
		vE.AddDescendant(NewTreeVertex(id, nil))
	}

	vF := vR.AddDescendant(NewTreeVertex("F", nil))
	vF.Threshold = 2
	for _, id := range []string{"1", "2", "3"} {
		vF.AddDescendant(NewTreeVertex(id, nil))
	}

	permutations := vR.ToTree().Permute()
//对于具有r-（d，e）的子树，我们有9个组合（每个子树有3个组合，其中d和e是根）
//对于具有r-（d，f）的子树，我们有9个来自相同逻辑的组合
//对于具有r-（e，f）的子树，我们也有9个组合
//共27种组合
	assert.Equal(t, 27, len(permutations))

	listCombination := func(i Iterator) []string {
		var traversal []string
		for {
			v := i.Next()
			if v == nil {
				break
			}
			traversal = append(traversal, v.Id)
		}
		return traversal
	}

//第一个组合是组合图上最左的遍历
	expectedScan := []string{"r", "D", "E", "A", "B", "a", "b"}
	assert.Equal(t, expectedScan, listCombination(permutations[0].BFS()))

//最后一个组合是组合图上最右的遍历
	expectedScan = []string{"r", "E", "F", "b", "c", "2", "3"}
	assert.Equal(t, expectedScan, listCombination(permutations[26].BFS()))
}

