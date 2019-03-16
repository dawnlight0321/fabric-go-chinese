
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455954619961344>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAndExists(t *testing.T) {
	v := NewTreeVertex("1", nil)
	u := v.AddDescendant(NewTreeVertex("2", nil)).AddDescendant(NewTreeVertex("4", nil))
	v.AddDescendant(NewTreeVertex("3", nil)).AddDescendant(NewTreeVertex("5", nil))
	assert.Equal(t, u, v.Find("4"))
	assert.True(t, v.Exists("4"))
	assert.Nil(t, v.Find("10"))
	assert.False(t, v.Exists("10"))
	assert.Nil(t, u.Find("1"))
	assert.False(t, u.Exists("1"))
	assert.Equal(t, v, v.Find("1"))
	assert.True(t, v.Exists("1"))
}

func TestIsLeaf(t *testing.T) {
	v := NewTreeVertex("1", nil)
	assert.True(t, v.AddDescendant(NewTreeVertex("2", nil)).IsLeaf())
	assert.False(t, v.IsLeaf())
}

func TestBFS(t *testing.T) {
	v := NewTreeVertex("1", nil)
	v.AddDescendant(NewTreeVertex("2", nil)).AddDescendant(NewTreeVertex("4", nil))
	v.AddDescendant(NewTreeVertex("3", nil)).AddDescendant(NewTreeVertex("5", nil))
	tree := v.ToTree()
	assert.Equal(t, v, tree.Root)
	i := tree.BFS()
	j := 1
	for {
		v := i.Next()
		if v == nil {
			assert.True(t, j == 6)
			break
		}
		assert.Equal(t, fmt.Sprintf("%d", j), v.Id)
		j++
	}
}

func TestClone(t *testing.T) {
	v := NewTreeVertex("1", 1)
	v.AddDescendant(NewTreeVertex("2", 2)).AddDescendant(NewTreeVertex("4", 3))
	v.AddDescendant(NewTreeVertex("3", 4)).AddDescendant(NewTreeVertex("5", 5))

	copy := v.Clone()
//它们是不同的参考文献
	assert.False(t, copy == v)
//他们是平等的。
	assert.Equal(t, v, copy)

	v.AddDescendant(NewTreeVertex("6", 6))
	assert.NotEqual(t, v, copy)
}

func TestReplace(t *testing.T) {
	v := &TreeVertex{
		Id: "r",
		Descendants: []*TreeVertex{
			{Id: "D", Descendants: []*TreeVertex{}},
			{Id: "E", Descendants: []*TreeVertex{}},
			{Id: "F", Descendants: []*TreeVertex{}},
		},
	}

	v.replace("D", &TreeVertex{
		Id: "d",
		Descendants: []*TreeVertex{
			{Id: "a", Descendants: []*TreeVertex{}},
			{Id: "b", Descendants: []*TreeVertex{}},
			{Id: "c", Descendants: []*TreeVertex{}},
		},
	})

	assert.Equal(t, "r", v.Id)
	assert.Equal(t, &TreeVertex{Id: "F", Descendants: []*TreeVertex{}}, v.Descendants[2])
	assert.Equal(t, &TreeVertex{Id: "E", Descendants: []*TreeVertex{}}, v.Descendants[1])
	assert.Equal(t, "D", v.Descendants[0].Id)
	assert.Equal(t, []*TreeVertex{
		{Id: "a", Descendants: []*TreeVertex{}},
		{Id: "b", Descendants: []*TreeVertex{}},
		{Id: "c", Descendants: []*TreeVertex{}},
	}, v.Descendants[0].Descendants)
}

