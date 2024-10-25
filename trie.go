package lgo

import "fmt"

// Node đại diện cho một node trong Trie
type NodeTrie struct {
	children map[rune]*NodeTrie
	isEnd    bool
}

// Trie là cấu trúc gốc của cây Trie
type Trie struct {
	root *NodeTrie
}

// NewNode khởi tạo một node mới
func NewNode() *NodeTrie {
	return &NodeTrie{
		children: make(map[rune]*NodeTrie),
		isEnd:    false,
	}
}

// NewTrie khởi tạo một Trie mới
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

// Insert thêm một từ vào Trie
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		// Nếu ký tự chưa có trong children, tạo node mới
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = NewNode()
		}
		// Di chuyển đến node tiếp theo
		currentNode = currentNode.children[char]
	}
	// Đánh dấu đây là kết thúc của một từ
	currentNode.isEnd = true
}

// Search kiểm tra xem từ có tồn tại trong Trie hay không
func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, char := range word {
		// Nếu ký tự không có trong children, trả về false
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		// Di chuyển đến node tiếp theo
		currentNode = currentNode.children[char]
	}
	// Kiểm tra xem node hiện tại có phải là kết thúc của một từ
	return currentNode.isEnd
}

// StartsWith kiểm tra xem có từ nào bắt đầu với prefix hay không
func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root
	for _, char := range prefix {
		// Nếu ký tự không có trong children, trả về false
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		// Di chuyển đến node tiếp theo
		currentNode = currentNode.children[char]
	}
	return true
}

// Main function để thử nghiệm Trie
func main() {
	trie := NewTrie()

	// Thêm từ vào Trie
	words := []string{"go", "golang", "gopher", "good", "game"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Kiểm tra sự tồn tại của từ
	fmt.Println(trie.Search("go"))     // true
	fmt.Println(trie.Search("golang")) // true
	fmt.Println(trie.Search("goph"))   // false

	// Kiểm tra tiền tố
	fmt.Println(trie.StartsWith("go"))  // true
	fmt.Println(trie.StartsWith("gop")) // true
	fmt.Println(trie.StartsWith("gam")) // true
	fmt.Println(trie.StartsWith("xyz")) // false
}
