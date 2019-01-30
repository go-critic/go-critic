package checkers

import (
	"go/ast"
	"regexp"
	"strings"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "codegenComment"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects malformed 'code generated' file comments"
	info.Before = `// This file was automatically generated by foogen`
	info.After = `// Code generated by foogen. DO NOT EDIT.`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		patterns := []string{
			"this file was auto(?:matically)? generated",
			"this file was generated by",
			"this file is auto(?:matically)? generated",
			// TODO(Quasilyte): more of these.
		}
		re := regexp.MustCompile("(?i)" + strings.Join(patterns, "|"))
		return &codegenCommentChecker{
			ctx:          ctx,
			badCommentRE: re,
		}
	})
}

type codegenCommentChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext

	badCommentRE *regexp.Regexp
}

func (c *codegenCommentChecker) WalkFile(f *ast.File) {
	if f.Doc == nil {
		return
	}
	for _, comment := range f.Doc.List {
		if c.badCommentRE.MatchString(comment.Text) {
			c.warn(comment)
			return
		}
	}
}

func (c *codegenCommentChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "comment should match `Code generated .* DO NOT EDIT.` regexp")
}
