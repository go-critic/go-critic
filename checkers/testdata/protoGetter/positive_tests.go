package checker_test

import (
	"fmt"

	pb "github.com/go-critic/go-critic/checkers/testdata/_importable/proto"
)

func _(t *pb.Test) interface{} {
	/*! proto message field read without getter: "t.S" should be "t.GetS()" */
	func(...interface{}) {}(t.GetS(), t.S)

	t.S, _ = "test", "test"
	_, t.S = "test", "test"
	_, t.S, _ = "test", "test", "test"
	_, _, t.S = "test", "test", "test"
	t.Embedded.S = "test"

	println(&t.S, &t.Embedded.S)

	/*! proto message field read without getter: "t.S" should be "t.GetS()" */
	/*! proto message field read without getter: "t.Embedded.S" should be "t.GetEmbedded().GetS()" */
	println(t.S, t.Embedded.S)

	var many []*pb.Test
	manyIndex := 42

	fmt.Println(
		/*! proto message field read without getter: "t.S" should be "t.GetS()" */
		t.S,
		/*! proto message field read without getter: "t.Embedded" should be "t.GetEmbedded()" */
		t.Embedded,
		/*! proto message field read without getter: "t.Embedded.S" should be "t.GetEmbedded().GetS()" */
		t.Embedded.S,
		/*! proto message field read without getter: "t.Embedded.Embedded.S" should be "t.GetEmbedded().GetEmbedded().GetS()" */
		t.Embedded.Embedded.S,
		/*! proto message field read without getter: "t.GetEmbedded().S" should be "t.GetEmbedded().GetS()" */
		t.GetEmbedded().S,
		/*! proto message field read without getter: "t.GetEmbedded().Embedded.S" should be "t.GetEmbedded().GetEmbedded().GetS()" */
		t.GetEmbedded().Embedded.S,
		/*! proto message field read without getter: "t.GetEmbedded().Embedded.GetS()" should be "t.GetEmbedded().GetEmbedded().GetS()" */
		t.GetEmbedded().Embedded.GetS(),
		/*! proto message field read without getter: "t.GetEmbedded().GetEmbedded().S" should be "t.GetEmbedded().GetEmbedded().GetS()" */
		t.GetEmbedded().GetEmbedded().S,
		/*! proto message field read without getter: "t.GetEmbedded().GetEmbedded().Embedded.S" should be "t.GetEmbedded().GetEmbedded().GetEmbedded().GetS()" */
		t.GetEmbedded().GetEmbedded().Embedded.S,
		/*! proto message field read without getter: "t.GetEmbedded().GetEmbedded().GetEmbedded().S" should be "t.GetEmbedded().GetEmbedded().GetEmbedded().GetS()" */
		t.GetEmbedded().GetEmbedded().GetEmbedded().S,
		/*! proto message field read without getter: "many[0].S" should be "many[0].GetS()" */
		many[0].S,
		/*! proto message field read without getter: "many[0].Embedded.S" should be "many[0].GetEmbedded().GetS()" */
		many[0].Embedded.S,
		/*! proto message field read without getter: "many[0].GetEmbedded().S" should be "many[0].GetEmbedded().GetS()" */
		many[0].GetEmbedded().S,
		/*! proto message field read without getter: "many[0].GetEmbedded().Embedded.S" should be "many[0].GetEmbedded().GetEmbedded().GetS()" */
		many[0].GetEmbedded().Embedded.S,
		/*! proto message field read without getter: "many[0].GetEmbedded().Embedded.Embedded.S" should be "many[0].GetEmbedded().GetEmbedded().GetEmbedded().GetS()" */
		many[0].GetEmbedded().Embedded.Embedded.S,
		/*! proto message field read without getter: "many[0].GetEmbedded().Embedded.Embedded.CustomMethod()" should be "many[0].GetEmbedded().GetEmbedded().GetEmbedded().CustomMethod()" */
		many[0].GetEmbedded().Embedded.Embedded.CustomMethod(),
		/*! proto message field read without getter: "many[0].GetEmbedded().Embedded.Embedded.GetS()" should be "many[0].GetEmbedded().GetEmbedded().GetEmbedded().GetS()" */
		many[0].GetEmbedded().Embedded.Embedded.GetS(),
		/*! proto message field read without getter: "many[manyIndex].S" should be "many[manyIndex].GetS()" */
		many[manyIndex].S,
	)
	/*! proto message field read without getter: "t.S" should be "t.GetS()" */
	return t.S
}
