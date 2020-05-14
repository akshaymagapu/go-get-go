/**
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
 */

package exercise

type Trie struct {
	val string
	child map[string]*Trie
	done bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		val: "",
		child: make(map[string]*Trie),
		done: false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node:= this
	for _, val:= range word {
		tmpValue := string(val)
		if node.child[tmpValue] == nil {
			node.child[tmpValue] = &Trie{
				val:   tmpValue,
				child: make(map[string]*Trie),
				done : false,
			}
		}
		node = node.child[tmpValue]
	}
	node.done = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, val:= range word {
		tmpValue := string(val)
		if node.child[tmpValue] == nil {
			return false
		}
		node = node.child[tmpValue]
	}
	return node.done
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, val:= range prefix {
		tmpValue := string(val)
		if node.child[tmpValue] == nil {
			return false
		}
		node = node.child[tmpValue]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */