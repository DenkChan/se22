package main

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type LinkTable struct {
	head   *LinkNode
	tail   *LinkNode
	length int
}

/* 创建节点 */
func NewLinkNode(node_data interface{}) *LinkNode {
	// Go语言可以自动决定局部变量作用域是否改变
	return &LinkNode{
		next: nil,
		data: node_data,
	}
}

/* 创建链表 */
func NewLinkTable() LinkTable {
	return LinkTable{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

/* 尾插 */
func (list *LinkTable) append(node_data interface{}) {
	//Go语言中这 相当于给LinkTable添加类成员函数
	new_node := NewLinkNode(node_data)

	if list.length == 0 {
		list.head = new_node
		list.tail = new_node
	} else {
		list.tail.next = new_node
		list.tail = new_node
	}

	list.length++

	return
}

/* pos处插入节点 */
func (list *LinkTable) insert(node_data interface{}, pos int) bool {

	if pos < 0 || pos > list.length {
		return false
	}

	new_node := NewLinkNode(node_data)

	if pos == 0 {
		new_node.next = list.head
		list.head = new_node
	} else {
		cur_node := list.head
		for i := 0; i < pos-1; i++ {
			cur_node = cur_node.next
		}
		new_node.next = cur_node.next
		cur_node.next = new_node
	}

	list.length++

	return true

}

/* 删除pos处节点 */
func (list *LinkTable) delete(position int) bool {

	if position <= 0 || position > list.length {
		return false
	}

	if position == 1 {
		list.head = list.head.next
	} else {
		pre_node := list.head
		for i := 1; i < position-1; i++ {
			pre_node = pre_node.next
		}
		pre_node.next = pre_node.next.next
		if position == list.length {
			list.tail = pre_node
		}

	}

	list.length--

	return true
}

/* 按值查找节点 */
func (list *LinkTable) find(search_data interface{}) *LinkNode {

	cur_node := list.head

	for i := 0; i < list.length; i++ {
		if cur_node.data == search_data {
			return cur_node
		}
		cur_node = cur_node.next
	}

	return nil
}

/* 获取队列长度 */
func (list *LinkTable) get_list_length() int {
	return list.length
}

/* 获取队头节点 */
func (list *LinkTable) get_list_head() *LinkNode {
	return list.head
}

/* 返回队尾节点 */
func (list *LinkTable) get_list_tail() *LinkNode {
	return list.tail
}

/* 打印所有节点 */
func (list *LinkTable) showlist() {
	if list.length == 0 {
		fmt.Println("List has no node.")
	} else {
		cur_node := list.head
		for i := 0; i < list.length; i++ {
			fmt.Print(cur_node.data, " ")
			cur_node = cur_node.next
		}
		fmt.Println()
	}
}
