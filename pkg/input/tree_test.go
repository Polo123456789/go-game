package input

import (
	"strconv"
	"testing"
)

type TestFunc func(t *testing.T)

func ExpectReset(t *testing.T, tree *MatchTree) {
	tree.Reset()
	if !tree.IsAtBase() {
		t.Fatalf("expected reset to root")
	}
}

func GenTreeExpectMatch(
	t *testing.T,
	tree *MatchTree,
	s string,
	expected TreeResult,
) TestFunc {
	return func(t *testing.T) {
		ExpectReset(t, tree)
		for _, c := range s {
			if !tree.Match(c) {
				t.Fatalf("expected match for %c", c)
			}
		}
		if tree.CurrentResult() != expected {
			t.Fatalf("expected result to be %d, got %d", expected, tree.CurrentResult())
		}
	}
}

func GenTreeDontExpectMatch(
	t *testing.T,
	tree *MatchTree,
	s string,
) TestFunc {
	return func(t *testing.T) {
		ExpectReset(t, tree)
		matchedAll := true
		for _, c := range s {
			if !tree.Match(c) {
				matchedAll = false
				break
			}
		}

		if matchedAll {
			t.Fatalf("expected no match")
		}

		if tree.CurrentResult() != NoMatch {
			t.Fatalf("expected result to be %d, got %d", NoMatch, tree.CurrentResult())
		}
	}
}

func TestMatchTree(t *testing.T) {
	const (
		MatchSingleCharacter TreeResult = iota
		MatchLongerString
		MatchSimilarString
		MatchSpecialCharacters
		MatchNumberFirst
	)

	const SingleCharacter = "a"
	const LongerString = "abcdefg"
	const SimilarString = "abcfeg"
	const WithSpecialCharacters = "!#$((#*@))"
	const NumberFirst = "0abc"

	const DontMatchSingle = "z"
	const DontMatchMid = "abcz"

	elements := []MatchTreeElement{
		{Value: SingleCharacter, Result: MatchSingleCharacter},
		{Value: LongerString, Result: MatchLongerString},
		{Value: SimilarString, Result: MatchSimilarString},
		{Value: WithSpecialCharacters, Result: MatchSpecialCharacters},
		{Value: NumberFirst, Result: MatchNumberFirst},
	}

	tree := NewMatchTree(elements)

	// Expect full matches
	for _, e := range elements {
		t.Run(
			"Plain Match: "+e.Value,
			GenTreeExpectMatch(t, tree, e.Value, e.Result),
		)
	}

	t.Run(
		"DontMatchSingleCharacter",
		GenTreeDontExpectMatch(t, tree, DontMatchSingle),
	)

	t.Run(
		"DontMatchMid",
		GenTreeDontExpectMatch(t, tree, DontMatchMid),
	)

	t.Run("MatchOrReset", func(t *testing.T) {
		tree.Reset()
		if !tree.MatchOrReset('a') {
			t.Errorf("expected a match to start")
		}
		if tree.MatchOrReset('z') {
			t.Errorf("expected no match")
		}
		if !tree.IsAtBase() {
			t.Errorf("expected reset to root")
		}
	})

	t.Run("CanContinueMatching", func(t *testing.T) {
		tree.Reset()
		if !tree.CanContinueMatching() {
			t.Errorf("Tree should be able to match if it hasn't started yet")
		}
		if !tree.Match('a') {
			t.Errorf("expected match")
		}
		if !tree.CanContinueMatching() {
			t.Errorf("Tree should be able to match")
		}
	})

	t.Run("WithNumericModifier", func(t *testing.T) {
		// Any modifier ending with a 0 will fail, because the tree matches
		// the string "0abc", and the strings in the tree take precedence
		modifiers := []int{
			1, 2, 3, 4, 5,
			6, 7, 8, 9,
			12,
			23,
			34235,
			4343,
			54,
			873,
		}

		for _, i := range modifiers {
			numericModifier := strconv.Itoa(i)
			expected := i

			for _, e := range elements {
				value := numericModifier + e.Value
				t.Run(
					"With Numeric Modifier: "+numericModifier+e.Value,
					GenTreeExpectMatch(t, tree, value, e.Result),
				)
				if tree.NumericModifier() != expected {
					t.Errorf("expected modifier to be %d, got %d", expected, tree.NumericModifier())
				}
			}
		}

	})
}
