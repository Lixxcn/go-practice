package main

import (
	"fmt"
)

// node 定义链表节点结构
// T any 表示这是一个泛型类型，可以存储任意类型的数据
type node[T any] struct {
	Data T        // 存储节点数据
	next *node[T] // 指向下一个节点的指针
}

// list 定义链表结构
type list[T any] struct {
	start *node[T] // 指向链表第一个节点的指针
}

// add 方法用于向链表添加新节点
// 这个实现使用了递归的方式，将新节点添加到链表末尾
func (l *list[T]) add(data T) {
	// 创建新节点
	n := node[T]{
		Data: data,
		next: nil,
	}

	// 如果链表为空，将新节点设为起始节点
	if l.start == nil {
		l.start = &n
		return
	}

	// 如果链表只有一个节点，将新节点添加到其后
	if l.start.next == nil {
		l.start.next = &n
		return
	}

	// 递归处理：保存当前起始节点，移动到下一个节点，递归添加
	// 递归完成后恢复起始节点位置
	temp := l.start        // 保存当前起始节点（比如12）
	l.start = l.start.next // 移动到下一个节点（比如9）
	l.add(data)            // 递归调用，继续处理
	l.start = temp         // 恢复起始节点位置
}

func main() {
	// 创建一个存储整数的链表
	var myList list[int]
	fmt.Println(myList) // 输出空链表

	// 添加第一个元素：12
	myList.add(12) // 结果：12 -> nil

	// 添加第二个元素：9
	myList.add(9) // 结果：12 -> 9 -> nil

	// 添加第三个元素：3
	// 最终结果：12 -> 9 -> 3 -> nil
	myList.add(3)

	// 添加第四个元素：9
	// 最终结果：12 -> 9 -> 3 -> 9 -> nil
	myList.add(9)

	// 打印链表所有元素
	cur := myList.start
	for {
		fmt.Println("*", cur) // 打印当前节点
		if cur == nil {
			break
		}
		cur = cur.next // 移动到下一个节点
	}
}
