package input

import (
	"testing"
)

type TestFunc func(t *testing.T)

func ExpectReset(t *testing.T, tree *MatchTree) {
	t.Run("Reset", func(t *testing.T) {
		tree.Reset()
		if tree.current != tree.root {
			t.Errorf("expected reset to root")
		}
	})
}

func GenTreeExpectMatch(
	t *testing.T,
	tree *MatchTree,
	s string,
	expected TreeResult,
) TestFunc {
	return func(t *testing.T) {
		for _, c := range s {
			if !tree.Match(c) {
				t.Errorf("expected match for %c", c)
			}
		}
		if tree.CurrentResult() != expected {
			t.Errorf("expected %d, got %d", expected, tree.CurrentResult())
		}
		ExpectReset(t, tree)
	}
}

func GenTreeDontExpectMatch(
	t *testing.T,
	tree *MatchTree,
	s string,
) TestFunc {
	return func(t *testing.T) {
		matchedAll := true
		for _, c := range s {
			if !tree.Match(c) {
				matchedAll = false
				break
			}
		}

		if matchedAll {
			t.Errorf("expected no match")
		}

		if tree.CurrentResult() != NoMatch {
			t.Errorf("expected %d, got %d", NoMatch, tree.CurrentResult())
		}
		ExpectReset(t, tree)
	}
}

func TestMatchTree(t *testing.T) {
	const (
		MatchSingleCharacter TreeResult = iota
		MatchLongerString
		MatchSimilarString
		MatchSpecialCharacters
		DontMatchSingleCharacter
		DontMatchMidTree
	)

	const SingleCharacter = "a"
	const LongerString = "abcdefg"
	const SimilarString = "abcfeg"
	const WithSpecialCharacters = "!#$((#*@))"

	const DontMatch = "z"
	const DontMatchMid = "abcz"

	tree := NewMatchTree([]MatchTreeElement{
		{Value: SingleCharacter, Result: MatchSingleCharacter},
		{Value: LongerString, Result: MatchLongerString},
		{Value: SimilarString, Result: MatchSimilarString},
		{Value: WithSpecialCharacters, Result: MatchSpecialCharacters},
	})

	t.Run(
		"SingleCharacter",
		GenTreeExpectMatch(t, tree, SingleCharacter, MatchSingleCharacter),
	)
	t.Run(
		"LongerString",
		GenTreeExpectMatch(t, tree, LongerString, MatchLongerString),
	)
	t.Run(
		"SimilarString",
		GenTreeExpectMatch(t, tree, SimilarString, MatchSimilarString),
	)
	t.Run(
		"WithSpecialCharacters",
		GenTreeExpectMatch(t, tree, WithSpecialCharacters, MatchSpecialCharacters),
	)
	t.Run(
		"DontMatchSingleCharacter",
		GenTreeDontExpectMatch(t, tree, DontMatch),
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
		if tree.current != tree.root {
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
}
