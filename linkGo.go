package main

import (
	"fmt"
)

type LinkNode struct {
	nextNode *LinkNode   //下一个结点
	data     interface{} //放数据
}

type LinkTable struct {
	head   *LinkNode //头结点 不放数据
	length int       //链表长度
}

//创建链表
func CreateLinkTable() *LinkTable {
	newLinkTable := new(LinkTable) //：= 声明并赋值
	firstNode := new(LinkNode)

	newLinkTable.head = firstNode

	newLinkTable.length = 0
	return newLinkTable
}

func DeleteLinkTable(curLinkTable *LinkTable) {
	for curLinkTable.length != 0 {
		deleteNode := curLinkTable.head.nextNode
		curLinkTable.head.nextNode = curLinkTable.head.nextNode.nextNode
		curLinkTable.length--
		deleteNode.data = nil
		deleteNode.data = nil

	}
}

//插入指定位置
func insertNode(position int, data interface{}, linkTable *LinkTable) *LinkTable {
	if position <= 0 || position > linkTable.length+1 {
		fmt.Println("error insert position")
		return nil
	}
	target := new(LinkNode)
	target.nextNode = nil
	target.data = data
	temp := linkTable.head
	for i := 0; i < position-1; i++ {
		temp = temp.nextNode
	}
	target.nextNode = temp.nextNode
	temp.nextNode = target
	linkTable.length++
	return linkTable
}

// 删除指定位置
func deleteNode(position int, linkTable *LinkTable) *LinkTable {
	if position <= 0 || position > linkTable.length {
		fmt.Println("error insert position")
		return nil
	}
	temp := linkTable.head
	for i := 0; i < position-1; i++ {
		temp = temp.nextNode
	}
	temp.nextNode = temp.nextNode.nextNode
	linkTable.length--

	return linkTable
}

//查找指定位置
func findNode(positon int, linkTable *LinkTable) *LinkNode {
	if positon <= 0 || positon > linkTable.length {
		fmt.Println("error insert position")
		return nil
	}
	temp := linkTable.head
	for i := 0; i < positon; i++ {
		temp = temp.nextNode
	}
	return temp
}
func printTable(linktable *LinkTable) {
	temp := linktable.head.nextNode
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.nextNode
	}
	fmt.Println("")
}
