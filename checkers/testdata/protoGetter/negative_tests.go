package checker_test

import (
	"fmt"

	pb "github.com/go-critic/go-critic/checkers/testdata/_importable/proto"
)

func _(t *pb.Test) interface{} {
	func(...interface{}) {}(t.GetS(), t.GetS())

	t.S, _ = "test", "test"
	_, t.S = "test", "test"
	_, t.S, _ = "test", "test", "test"
	_, _, t.S = "test", "test", "test"
	t.Embedded.S = "test"

	println(&t.S, &t.Embedded.S)

	println(t.GetS(), t.GetEmbedded().GetS())

	var many []*pb.Test
	manyIndex := 42

	fmt.Println(
		t.GetS(),
		t.GetEmbedded(),
		t.GetEmbedded().GetS(),
		t.GetEmbedded().GetEmbedded().GetS(),
		t.GetEmbedded().GetS(),
		t.GetEmbedded().GetEmbedded().GetS(),
		t.GetEmbedded().GetEmbedded().GetS(),
		t.GetEmbedded().GetEmbedded().GetS(),
		t.GetEmbedded().GetEmbedded().GetEmbedded().GetS(),
		t.GetEmbedded().GetEmbedded().GetEmbedded().GetS(),
		many[0].GetS(),
		many[0].GetEmbedded().GetS(),
		many[0].GetEmbedded().GetS(),
		many[0].GetEmbedded().GetEmbedded().GetS(),
		many[0].GetEmbedded().GetEmbedded().GetEmbedded().GetS(),
		many[0].GetEmbedded().GetEmbedded().GetEmbedded().CustomMethod(),
		many[0].GetEmbedded().GetEmbedded().GetEmbedded().GetS(),
		many[manyIndex].GetS(),
	)
	return t.GetS()
}
