package input

import (
	"unicode"
)

type TreeResult int

const NoMatch TreeResult = -1
const DefaultNumericModifier int = 1

type node struct {
	value    rune
	children map[rune]*node
	result   TreeResult
}

type MatchTree struct {
	root                     *node
	current                  *node
	numericModifier          int
	noCurrentNumericModifier bool
}

type MatchTreeElement struct {
	Value  string
	Result TreeResult
}

func NewMatchTree(elems []MatchTreeElement) *MatchTree {
	var tree MatchTree
	tree.root = &node{0, make(map[rune]*node), NoMatch}
	tree.current = tree.root
	tree.numericModifier = DefaultNumericModifier
	for _, elem := range elems {
		tree.Add(elem)
	}
	return &tree
}

func (tree *MatchTree) Add(elem MatchTreeElement) {
	curr := tree.root
	for _, b := range elem.Value {
		next, ok := curr.children[b]
		if !ok {
			curr.children[b] = &node{
				value:    b,
				children: make(map[rune]*node),
				result:   NoMatch,
			}
			next = curr.children[b]
		}
		curr = next
	}
	curr.result = elem.Result
}

func (tree *MatchTree) Reset() {
	tree.current = tree.root
	tree.numericModifier = DefaultNumericModifier
	tree.noCurrentNumericModifier = true
}

func (tree *MatchTree) IsAtBase() bool {
	return tree.current == tree.root
}

func (tree *MatchTree) Match(c rune) bool {
	next, ok := tree.current.children[c]
	if ok {
		tree.current = next
		return true
	}

	// Try to match a numeric modifier
	if !tree.IsAtBase() || !unicode.IsDigit(c) {
		return false
	}

	asInt := int(c - '0')
	if tree.noCurrentNumericModifier {
		tree.numericModifier = asInt
		tree.noCurrentNumericModifier = false
	} else {
		tree.numericModifier = tree.numericModifier*10 + asInt
	}
	return true
}

func (tree *MatchTree) MatchOrReset(c rune) bool {
	matched := tree.Match(c)
	if !matched {
		tree.Reset()
	}
	return matched
}

func (tree *MatchTree) CurrentResult() TreeResult {
	return tree.current.result
}

func (tree *MatchTree) NumericModifier() int {
	return tree.numericModifier
}

func (tree *MatchTree) CanContinueMatching() bool {
	return len(tree.current.children) != 0
}
