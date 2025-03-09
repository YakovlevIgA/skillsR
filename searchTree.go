package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

// Insert вставляет новое значение в дерево
func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{Value: value}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.Root.insert(newNode)
	}
}

// Вспомогательная функция вставки
func (node *TreeNode) insert(newNode *TreeNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Left.insert(newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			node.Right.insert(newNode)
		}
	}
}

// Search ищет значение в дереве и возвращает true, если оно найдено
func (bst *BinarySearchTree) Search(value int) bool {
	return bst.Root != nil && bst.Root.search(value)
}

// Вспомогательная функция поиска
func (node *TreeNode) search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.search(value)
	} else if value > node.Value {
		return node.Right.search(value)
	} else {
		return true
	}
}

// Delete удаляет узел с заданным значением из дерева
func (bst *BinarySearchTree) Delete(value int) {
	bst.Root = bst.Root.delete(value)
}

// Вспомогательная функция удаления узла
func (node *TreeNode) delete(value int) *TreeNode {
	if node == nil {
		return nil
	}

	// Найден ли узел для удаления?
	if value < node.Value {
		node.Left = node.Left.delete(value)
	} else if value > node.Value {
		node.Right = node.Right.delete(value)
	} else {
		// Узел найден

		// Узел с одним или без потомков
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Узел с двумя потомками, получаем наименьшее значение из правого поддерева
		minRight := node.Right.findMin()
		node.Value = minRight.Value
		node.Right = node.Right.delete(minRight.Value)
	}

	return node
}

// findMin находит минимальный элемент в дереве
func (node *TreeNode) findMin() *TreeNode {
	if node.Left != nil {
		return node.Left.findMin()
	}
	return node
}

// Preorder (прямой обход): текущий узел, левое поддерево, правое поддерево
func (bst *BinarySearchTree) Preorder() {
	bst.Root.preorder()
}

func (node *TreeNode) preorder() {
	if node != nil {
		fmt.Print(node.Value, " ")
		node.Left.preorder()
		node.Right.preorder()
	}
}

// Inorder (центрированный обход): левое поддерево, текущий узел, правое поддерево
func (bst *BinarySearchTree) Inorder() {
	bst.Root.inorder()
}

func (node *TreeNode) inorder() {
	if node != nil {
		node.Left.inorder()
		fmt.Print(node.Value, " ")
		node.Right.inorder()
	}
}

// Postorder (обратный обход): левое поддерево, правое поддерево, текущий узел
func (bst *BinarySearchTree) Postorder() {
	bst.Root.postorder()
}

func (node *TreeNode) postorder() {
	if node != nil {
		node.Left.postorder()
		node.Right.postorder()
		fmt.Print(node.Value, " ")
	}
}

func main() {
	bst := &BinarySearchTree{}
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(70)
	bst.Insert(60)
	bst.Insert(80)

	// Поиск
	fmt.Println("Search for 40:", bst.Search(40))
	fmt.Println("Search for 100:", bst.Search(100))

	// Удаление
	bst.Delete(20)
	fmt.Println("Inorder after deleting 20:")
	bst.Inorder() // Вывод: 30 40 50 60 70 80
	fmt.Println()

	bst.Delete(30)
	fmt.Println("Inorder after deleting 30:")
	bst.Inorder() // Вывод: 40 50 60 70 80
	fmt.Println()

	bst.Delete(50)
	fmt.Println("Inorder after deleting 50:")
	bst.Inorder() // Вывод: 40 60 70 80
	fmt.Println()

	// Обходы
	fmt.Println("Preorder:")
	bst.Preorder() // Вывод: 40 60 50 70 80
	fmt.Println()

	fmt.Println("Postorder:")
	bst.Postorder() // Вывод: 40 60 80 70 50
	fmt.Println()
}
