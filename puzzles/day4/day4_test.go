package puzzles

import "testing"

func TestPart1SampleInput(t *testing.T) {
	input := [][][]int{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	want := 2

	answer := SovleDay4Part1(input)

	if answer != want {
		t.Fatalf("expected %v got %v", want, answer)
	}
}

func TestPart2SampleInput(t *testing.T) {
	input := [][][]int{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	want := 4

	answer := SovleDay4Part2(input)

	if answer != want {
		t.Fatalf("expected %v got %v", want, answer)
	}
}
