// Code generated by "precompile.go". DO NOT EDIT.

package rulesdata

import "github.com/quasilyte/go-ruleguard/ruleguard/ir"

var PrecompiledRules = &ir.File{
	PkgPath:       "gorules",
	CustomDecls:   []string{},
	BundleImports: []ir.BundleImport{},
	RuleGroups: []ir.RuleGroup{
		ir.RuleGroup{
			Line:        11,
			Name:        "deferUnlambda",
			MatcherName: "m",
			DocTags: []string{
				"style",
				"experimental",
			},
			DocSummary: "Detects deferred function literals that can be simplified",
			DocBefore:  "defer func() { f() }()",
			DocAfter:   "defer f()",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           12,
					SyntaxPattern:  "defer func() { $f($*args) }()",
					ReportTemplate: "can rewrite as `defer $f($args)`",
					WhereExpr: ir.FilterExpr{
						Line: 13,
						Op:   ir.FilterAndOp,
						Src:  "m[\"f\"].Node.Is(`Ident`) && m[\"f\"].Text != \"panic\" && m[\"f\"].Text != \"recover\" && m[\"args\"].Const",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line: 13,
								Op:   ir.FilterAndOp,
								Src:  "m[\"f\"].Node.Is(`Ident`) && m[\"f\"].Text != \"panic\" && m[\"f\"].Text != \"recover\"",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line: 13,
										Op:   ir.FilterAndOp,
										Src:  "m[\"f\"].Node.Is(`Ident`) && m[\"f\"].Text != \"panic\"",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  13,
												Op:    ir.FilterVarNodeIsOp,
												Src:   "m[\"f\"].Node.Is(`Ident`)",
												Value: "f",
												Args: []ir.FilterExpr{
													ir.FilterExpr{
														Line:  13,
														Op:    ir.FilterStringOp,
														Src:   "`Ident`",
														Value: "Ident",
													},
												},
											},
											ir.FilterExpr{
												Line: 13,
												Op:   ir.FilterNeqOp,
												Src:  "m[\"f\"].Text != \"panic\"",
												Args: []ir.FilterExpr{
													ir.FilterExpr{
														Line:  13,
														Op:    ir.FilterVarTextOp,
														Src:   "m[\"f\"].Text",
														Value: "f",
													},
													ir.FilterExpr{
														Line:  13,
														Op:    ir.FilterStringOp,
														Src:   "\"panic\"",
														Value: "panic",
													},
												},
											},
										},
									},
									ir.FilterExpr{
										Line: 13,
										Op:   ir.FilterNeqOp,
										Src:  "m[\"f\"].Text != \"recover\"",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  13,
												Op:    ir.FilterVarTextOp,
												Src:   "m[\"f\"].Text",
												Value: "f",
											},
											ir.FilterExpr{
												Line:  13,
												Op:    ir.FilterStringOp,
												Src:   "\"recover\"",
												Value: "recover",
											},
										},
									},
								},
							},
							ir.FilterExpr{
								Line:  13,
								Op:    ir.FilterVarConstOp,
								Src:   "m[\"args\"].Const",
								Value: "args",
							},
						},
					},
				},
				ir.Rule{
					Line:           16,
					SyntaxPattern:  "defer func() { $pkg.$f($*args) }()",
					ReportTemplate: "can rewrite as `defer $pkg.$f($args)`",
					WhereExpr: ir.FilterExpr{
						Line: 17,
						Op:   ir.FilterAndOp,
						Src:  "m[\"f\"].Node.Is(`Ident`) && m[\"args\"].Const && m[\"pkg\"].Object.Is(`PkgName`)",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line: 17,
								Op:   ir.FilterAndOp,
								Src:  "m[\"f\"].Node.Is(`Ident`) && m[\"args\"].Const",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  17,
										Op:    ir.FilterVarNodeIsOp,
										Src:   "m[\"f\"].Node.Is(`Ident`)",
										Value: "f",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  17,
												Op:    ir.FilterStringOp,
												Src:   "`Ident`",
												Value: "Ident",
											},
										},
									},
									ir.FilterExpr{
										Line:  17,
										Op:    ir.FilterVarConstOp,
										Src:   "m[\"args\"].Const",
										Value: "args",
									},
								},
							},
							ir.FilterExpr{
								Line:  17,
								Op:    ir.FilterVarObjectIsOp,
								Src:   "m[\"pkg\"].Object.Is(`PkgName`)",
								Value: "pkg",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  17,
										Op:    ir.FilterStringOp,
										Src:   "`PkgName`",
										Value: "PkgName",
									},
								},
							},
						},
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        25,
			Name:        "ioutilDeprecated",
			MatcherName: "m",
			DocTags: []string{
				"style",
				"experimental",
			},
			DocSummary: "Detects deprecated io/ioutil package usages",
			DocBefore:  "ioutil.ReadAll(r)",
			DocAfter:   "io.ReadAll(r)",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           26,
					SyntaxPattern:  "ioutil.ReadAll($_)",
					ReportTemplate: "ioutil.ReadAll is deprecated, use io.ReadAll instead",
				},
				ir.Rule{
					Line:           29,
					SyntaxPattern:  "ioutil.ReadFile($_)",
					ReportTemplate: "ioutil.ReadFile is deprecated, use os.ReadFile instead",
				},
				ir.Rule{
					Line:           32,
					SyntaxPattern:  "ioutil.WriteFile($_, $_, $_)",
					ReportTemplate: "ioutil.WriteFile is deprecated, use os.WriteFile instead",
				},
				ir.Rule{
					Line:           35,
					SyntaxPattern:  "ioutil.ReadDir($_)",
					ReportTemplate: "ioutil.ReadDir is deprecated, use os.ReadDir instead",
				},
				ir.Rule{
					Line:           38,
					SyntaxPattern:  "ioutil.NopCloser($_)",
					ReportTemplate: "ioutil.NopCloser is deprecated, use io.NopCloser instead",
				},
				ir.Rule{
					Line:           41,
					SyntaxPattern:  "ioutil.Discard",
					ReportTemplate: "ioutil.Discard is deprecated, use io.Discard instead",
				},
			},
		},
		ir.RuleGroup{
			Line:        49,
			Name:        "badLock",
			MatcherName: "m",
			DocTags: []string{
				"diagnostic",
				"experimental",
			},
			DocSummary: "Detects suspicious mutex lock/unlock operations",
			DocBefore:  "mu.Lock(); mu.Unlock()",
			DocAfter:   "mu.Lock(); defer mu.Unlock()",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           53,
					SyntaxPattern:  "$mu1.Lock(); $mu2.Unlock()",
					ReportTemplate: "defer is missing, mutex is unlocked immediately",
					WhereExpr: ir.FilterExpr{
						Line: 54,
						Op:   ir.FilterEqOp,
						Src:  "m[\"mu1\"].Text == m[\"mu2\"].Text",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  54,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu1\"].Text",
								Value: "mu1",
							},
							ir.FilterExpr{
								Line:  54,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu2\"].Text",
								Value: "mu2",
							},
						},
					},
					LocationVar: "mu2",
				},
				ir.Rule{
					Line:           58,
					SyntaxPattern:  "$mu1.RLock(); $mu2.RUnlock()",
					ReportTemplate: "defer is missing, mutex is unlocked immediately",
					WhereExpr: ir.FilterExpr{
						Line: 59,
						Op:   ir.FilterEqOp,
						Src:  "m[\"mu1\"].Text == m[\"mu2\"].Text",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  59,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu1\"].Text",
								Value: "mu1",
							},
							ir.FilterExpr{
								Line:  59,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu2\"].Text",
								Value: "mu2",
							},
						},
					},
					LocationVar: "mu2",
				},
				ir.Rule{
					Line:           64,
					SyntaxPattern:  "$mu1.Lock(); defer $mu2.RUnlock()",
					ReportTemplate: "suspicious unlock, maybe Unlock was intended?",
					WhereExpr: ir.FilterExpr{
						Line: 65,
						Op:   ir.FilterEqOp,
						Src:  "m[\"mu1\"].Text == m[\"mu2\"].Text",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  65,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu1\"].Text",
								Value: "mu1",
							},
							ir.FilterExpr{
								Line:  65,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu2\"].Text",
								Value: "mu2",
							},
						},
					},
					LocationVar: "mu2",
				},
				ir.Rule{
					Line:           69,
					SyntaxPattern:  "$mu1.RLock(); defer $mu2.Unlock()",
					ReportTemplate: "suspicious unlock, maybe RUnlock was intended?",
					WhereExpr: ir.FilterExpr{
						Line: 70,
						Op:   ir.FilterEqOp,
						Src:  "m[\"mu1\"].Text == m[\"mu2\"].Text",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  70,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu1\"].Text",
								Value: "mu1",
							},
							ir.FilterExpr{
								Line:  70,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu2\"].Text",
								Value: "mu2",
							},
						},
					},
					LocationVar: "mu2",
				},
				ir.Rule{
					Line:           75,
					SyntaxPattern:  "$mu1.Lock(); defer $mu2.Lock()",
					ReportTemplate: "maybe defer $mu1.Unlock() was intended?",
					WhereExpr: ir.FilterExpr{
						Line: 76,
						Op:   ir.FilterEqOp,
						Src:  "m[\"mu1\"].Text == m[\"mu2\"].Text",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  76,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu1\"].Text",
								Value: "mu1",
							},
							ir.FilterExpr{
								Line:  76,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu2\"].Text",
								Value: "mu2",
							},
						},
					},
					LocationVar: "mu2",
				},
				ir.Rule{
					Line:           80,
					SyntaxPattern:  "$mu1.RLock(); defer $mu2.RLock()",
					ReportTemplate: "maybe defer $mu1.RUnlock() was intended?",
					WhereExpr: ir.FilterExpr{
						Line: 81,
						Op:   ir.FilterEqOp,
						Src:  "m[\"mu1\"].Text == m[\"mu2\"].Text",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  81,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu1\"].Text",
								Value: "mu1",
							},
							ir.FilterExpr{
								Line:  81,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"mu2\"].Text",
								Value: "mu2",
							},
						},
					},
					LocationVar: "mu2",
				},
			},
		},
		ir.RuleGroup{
			Line:        90,
			Name:        "httpNoBody",
			MatcherName: "m",
			DocTags: []string{
				"style",
				"experimental",
			},
			DocSummary: "Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative",
			DocBefore:  "http.NewRequest(\"GET\", url, nil)",
			DocAfter:   "http.NewRequest(\"GET\", url, http.NoBody)",
			Rules: []ir.Rule{
				ir.Rule{
					Line:            91,
					SyntaxPattern:   "http.NewRequest($method, $url, $nil)",
					ReportTemplate:  "http.NoBody should be preferred to the nil request body",
					SuggestTemplate: "http.NewRequest($method, $url, http.NoBody)",
					WhereExpr: ir.FilterExpr{
						Line: 92,
						Op:   ir.FilterEqOp,
						Src:  "m[\"nil\"].Text == \"nil\"",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  92,
								Op:    ir.FilterVarTextOp,
								Src:   "m[\"nil\"].Text",
								Value: "nil",
							},
							ir.FilterExpr{
								Line:  92,
								Op:    ir.FilterStringOp,
								Src:   "\"nil\"",
								Value: "nil",
							},
						},
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        102,
			Name:        "preferDecodeRune",
			MatcherName: "m",
			DocTags: []string{
				"performance",
				"experimental",
			},
			DocSummary: "Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation",
			DocBefore:  "r := []rune(s)[0]",
			DocAfter:   "r, _ := utf8.DecodeRuneInString(s)",
			DocNote:    "See Go issue for details: https://github.com/golang/go/issues/45260",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           103,
					SyntaxPattern:  "[]rune($s)[0]",
					ReportTemplate: "consider replacing $$ with utf8.DecodeRuneInString($s)",
					WhereExpr: ir.FilterExpr{
						Line:  104,
						Op:    ir.FilterVarTypeIsOp,
						Src:   "m[\"s\"].Type.Is(`string`)",
						Value: "s",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  104,
								Op:    ir.FilterStringOp,
								Src:   "`string`",
								Value: "string",
							},
						},
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        112,
			Name:        "sloppyLen",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects usage of `len` when result is obvious or doesn't make sense",
			DocBefore:  "len(arr) <= 0",
			DocAfter:   "len(arr) == 0",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           113,
					SyntaxPattern:  "len($_) >= 0",
					ReportTemplate: "$$ is always true",
				},
				ir.Rule{
					Line:           114,
					SyntaxPattern:  "len($_) < 0",
					ReportTemplate: "$$ is always false",
				},
				ir.Rule{
					Line:           115,
					SyntaxPattern:  "len($x) <= 0",
					ReportTemplate: "$$ can be len($x) == 0",
				},
			},
		},
		ir.RuleGroup{
			Line:        122,
			Name:        "valSwap",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects value swapping code that are not using parallel assignment",
			DocBefore:  "*tmp = *x; *x = *y; *y = *tmp",
			DocAfter:   "*x, *y = *y, *x",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           123,
					SyntaxPattern:  "$tmp := $y; $y = $x; $x = $tmp",
					ReportTemplate: "can re-write as `$y, $x = $x, $y`",
				},
			},
		},
		ir.RuleGroup{
			Line:        131,
			Name:        "switchTrue",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects switch-over-bool statements that use explicit `true` tag value",
			DocBefore:  "switch true {...}",
			DocAfter:   "switch {...}",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           132,
					SyntaxPattern:  "switch true { $*_ }",
					ReportTemplate: "replace 'switch true {}' with 'switch {}'",
				},
				ir.Rule{
					Line:           134,
					SyntaxPattern:  "switch $x; true { $*_ }",
					ReportTemplate: "replace 'switch $x; true {}' with 'switch $x; {}'",
				},
			},
		},
		ir.RuleGroup{
			Line:        142,
			Name:        "flagDeref",
			MatcherName: "m",
			DocTags: []string{
				"diagnostic",
			},
			DocSummary: "Detects immediate dereferencing of `flag` package pointers",
			DocBefore:  "b := *flag.Bool(\"b\", false, \"b docs\")",
			DocAfter:   "var b bool; flag.BoolVar(&b, \"b\", false, \"b docs\")",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           143,
					SyntaxPattern:  "*flag.Bool($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.BoolVar",
				},
				ir.Rule{
					Line:           144,
					SyntaxPattern:  "*flag.Duration($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.DurationVar",
				},
				ir.Rule{
					Line:           145,
					SyntaxPattern:  "*flag.Float64($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.Float64Var",
				},
				ir.Rule{
					Line:           146,
					SyntaxPattern:  "*flag.Int($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.IntVar",
				},
				ir.Rule{
					Line:           147,
					SyntaxPattern:  "*flag.Int64($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.Int64Var",
				},
				ir.Rule{
					Line:           148,
					SyntaxPattern:  "*flag.String($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.StringVar",
				},
				ir.Rule{
					Line:           149,
					SyntaxPattern:  "*flag.Uint($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.UintVar",
				},
				ir.Rule{
					Line:           150,
					SyntaxPattern:  "*flag.Uint64($*_)",
					ReportTemplate: "immediate deref in $$ is most likely an error; consider using flag.Uint64Var",
				},
			},
		},
		ir.RuleGroup{
			Line:        157,
			Name:        "emptyStringTest",
			MatcherName: "m",
			DocTags: []string{
				"style",
				"experimental",
			},
			DocSummary: "Detects empty string checks that can be written more idiomatically",
			DocBefore:  "len(s) == 0",
			DocAfter:   "s == \"\"",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           158,
					SyntaxPattern:  "len($s) != 0",
					ReportTemplate: "replace `$$` with `$s != \"\"`",
					WhereExpr: ir.FilterExpr{
						Line:  159,
						Op:    ir.FilterVarTypeIsOp,
						Src:   "m[\"s\"].Type.Is(`string`)",
						Value: "s",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  159,
								Op:    ir.FilterStringOp,
								Src:   "`string`",
								Value: "string",
							},
						},
					},
				},
				ir.Rule{
					Line:           162,
					SyntaxPattern:  "len($s) == 0",
					ReportTemplate: "replace `$$` with `$s == \"\"`",
					WhereExpr: ir.FilterExpr{
						Line:  163,
						Op:    ir.FilterVarTypeIsOp,
						Src:   "m[\"s\"].Type.Is(`string`)",
						Value: "s",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  163,
								Op:    ir.FilterStringOp,
								Src:   "`string`",
								Value: "string",
							},
						},
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        171,
			Name:        "stringXbytes",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects redundant conversions between string and []byte",
			DocBefore:  "copy(b, []byte(s))",
			DocAfter:   "copy(b, s)",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           172,
					SyntaxPattern:  "copy($_, []byte($s))",
					ReportTemplate: "can simplify `[]byte($s)` to `$s`",
				},
			},
		},
		ir.RuleGroup{
			Line:        180,
			Name:        "indexAlloc",
			MatcherName: "m",
			DocTags: []string{
				"performance",
			},
			DocSummary: "Detects strings.Index calls that may cause unwanted allocs",
			DocBefore:  "strings.Index(string(x), y)",
			DocAfter:   "bytes.Index(x, []byte(y))",
			DocNote:    "See Go issue for details: https://github.com/golang/go/issues/25864",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           181,
					SyntaxPattern:  "strings.Index(string($x), $y)",
					ReportTemplate: "consider replacing $$ with bytes.Index($x, []byte($y))",
					WhereExpr: ir.FilterExpr{
						Line: 182,
						Op:   ir.FilterAndOp,
						Src:  "m[\"x\"].Pure && m[\"y\"].Pure",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  182,
								Op:    ir.FilterVarPureOp,
								Src:   "m[\"x\"].Pure",
								Value: "x",
							},
							ir.FilterExpr{
								Line:  182,
								Op:    ir.FilterVarPureOp,
								Src:   "m[\"y\"].Pure",
								Value: "y",
							},
						},
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        190,
			Name:        "wrapperFunc",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects function calls that can be replaced with convenience wrappers",
			DocBefore:  "wg.Add(-1)",
			DocAfter:   "wg.Done()",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           191,
					SyntaxPattern:  "$wg.Add(-1)",
					ReportTemplate: "use WaitGroup.Done method in `$$`",
					WhereExpr: ir.FilterExpr{
						Line:  192,
						Op:    ir.FilterVarTypeIsOp,
						Src:   "m[\"wg\"].Type.Is(`sync.WaitGroup`)",
						Value: "wg",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  192,
								Op:    ir.FilterStringOp,
								Src:   "`sync.WaitGroup`",
								Value: "sync.WaitGroup",
							},
						},
					},
				},
				ir.Rule{
					Line:           195,
					SyntaxPattern:  "$buf.Truncate(0)",
					ReportTemplate: "use Buffer.Reset method in `$$`",
					WhereExpr: ir.FilterExpr{
						Line:  196,
						Op:    ir.FilterVarTypeIsOp,
						Src:   "m[\"buf\"].Type.Is(`bytes.Buffer`)",
						Value: "buf",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  196,
								Op:    ir.FilterStringOp,
								Src:   "`bytes.Buffer`",
								Value: "bytes.Buffer",
							},
						},
					},
				},
				ir.Rule{
					Line:           199,
					SyntaxPattern:  "http.HandlerFunc(http.NotFound)",
					ReportTemplate: "use http.NotFoundHandler method in `$$`",
				},
				ir.Rule{
					Line:           201,
					SyntaxPattern:  "strings.SplitN($_, $_, -1)",
					ReportTemplate: "use strings.Split method in `$$`",
				},
				ir.Rule{
					Line:           202,
					SyntaxPattern:  "strings.Replace($_, $_, $_, -1)",
					ReportTemplate: "use strings.ReplaceAll method in `$$`",
				},
				ir.Rule{
					Line:           203,
					SyntaxPattern:  "strings.Map(unicode.ToTitle, $_)",
					ReportTemplate: "use strings.ToTitle method in `$$`",
				},
				ir.Rule{
					Line:           205,
					SyntaxPattern:  "bytes.SplitN(b, []byte(\".\"), -1)",
					ReportTemplate: "use bytes.Split method in `$$`",
				},
				ir.Rule{
					Line:           206,
					SyntaxPattern:  "bytes.Replace($_, $_, $_, -1)",
					ReportTemplate: "use bytes.ReplaceAll method in `$$`",
				},
				ir.Rule{
					Line:           207,
					SyntaxPattern:  "bytes.Map(unicode.ToUpper, $_)",
					ReportTemplate: "use bytes.ToUpper method in `$$`",
				},
				ir.Rule{
					Line:           208,
					SyntaxPattern:  "bytes.Map(unicode.ToLower, $_)",
					ReportTemplate: "use bytes.ToLower method in `$$`",
				},
				ir.Rule{
					Line:           209,
					SyntaxPattern:  "bytes.Map(unicode.ToTitle, $_)",
					ReportTemplate: "use bytes.ToTitle method in `$$`",
				},
				ir.Rule{
					Line:           211,
					SyntaxPattern:  "draw.DrawMask($_, $_, $_, $_, nil, image.Point{}, $_)",
					ReportTemplate: "use draw.Draw method in `$$`",
				},
			},
		},
		ir.RuleGroup{
			Line:        219,
			Name:        "regexpMust",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`",
			DocBefore:  "re, _ := regexp.Compile(\"const pattern\")",
			DocAfter:   "re := regexp.MustCompile(\"const pattern\")",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           220,
					SyntaxPattern:  "regexp.Compile($pat)",
					ReportTemplate: "for const patterns like $pat, use regexp.MustCompile",
					WhereExpr: ir.FilterExpr{
						Line:  221,
						Op:    ir.FilterVarConstOp,
						Src:   "m[\"pat\"].Const",
						Value: "pat",
					},
				},
				ir.Rule{
					Line:           224,
					SyntaxPattern:  "regexp.CompilePOSIX($pat)",
					ReportTemplate: "for const patterns like $pat, use regexp.MustCompilePOSIX",
					WhereExpr: ir.FilterExpr{
						Line:  225,
						Op:    ir.FilterVarConstOp,
						Src:   "m[\"pat\"].Const",
						Value: "pat",
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        233,
			Name:        "badCall",
			MatcherName: "m",
			DocTags: []string{
				"diagnostic",
			},
			DocSummary: "Detects suspicious function calls",
			DocBefore:  "strings.Replace(s, from, to, 0)",
			DocAfter:   "strings.Replace(s, from, to, -1)",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           234,
					SyntaxPattern:  "strings.Replace($_, $_, $_, $zero)",
					ReportTemplate: "suspicious arg 0, probably meant -1",
					WhereExpr: ir.FilterExpr{
						Line: 235,
						Op:   ir.FilterEqOp,
						Src:  "m[\"zero\"].Value.Int() == 0",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  235,
								Op:    ir.FilterVarValueIntOp,
								Src:   "m[\"zero\"].Value.Int()",
								Value: "zero",
							},
							ir.FilterExpr{
								Line:  235,
								Op:    ir.FilterIntOp,
								Src:   "0",
								Value: int64(0),
							},
						},
					},
					LocationVar: "zero",
				},
				ir.Rule{
					Line:           237,
					SyntaxPattern:  "bytes.Replace($_, $_, $_, $zero)",
					ReportTemplate: "suspicious arg 0, probably meant -1",
					WhereExpr: ir.FilterExpr{
						Line: 238,
						Op:   ir.FilterEqOp,
						Src:  "m[\"zero\"].Value.Int() == 0",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  238,
								Op:    ir.FilterVarValueIntOp,
								Src:   "m[\"zero\"].Value.Int()",
								Value: "zero",
							},
							ir.FilterExpr{
								Line:  238,
								Op:    ir.FilterIntOp,
								Src:   "0",
								Value: int64(0),
							},
						},
					},
					LocationVar: "zero",
				},
				ir.Rule{
					Line:           241,
					SyntaxPattern:  "strings.SplitN($_, $_, $zero)",
					ReportTemplate: "suspicious arg 0, probably meant -1",
					WhereExpr: ir.FilterExpr{
						Line: 242,
						Op:   ir.FilterEqOp,
						Src:  "m[\"zero\"].Value.Int() == 0",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  242,
								Op:    ir.FilterVarValueIntOp,
								Src:   "m[\"zero\"].Value.Int()",
								Value: "zero",
							},
							ir.FilterExpr{
								Line:  242,
								Op:    ir.FilterIntOp,
								Src:   "0",
								Value: int64(0),
							},
						},
					},
					LocationVar: "zero",
				},
				ir.Rule{
					Line:           244,
					SyntaxPattern:  "bytes.SplitN($_, $_, $zero)",
					ReportTemplate: "suspicious arg 0, probably meant -1",
					WhereExpr: ir.FilterExpr{
						Line: 245,
						Op:   ir.FilterEqOp,
						Src:  "m[\"zero\"].Value.Int() == 0",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  245,
								Op:    ir.FilterVarValueIntOp,
								Src:   "m[\"zero\"].Value.Int()",
								Value: "zero",
							},
							ir.FilterExpr{
								Line:  245,
								Op:    ir.FilterIntOp,
								Src:   "0",
								Value: int64(0),
							},
						},
					},
					LocationVar: "zero",
				},
				ir.Rule{
					Line:           248,
					SyntaxPattern:  "append($_)",
					ReportTemplate: "no-op append call, probably missing arguments",
				},
			},
		},
		ir.RuleGroup{
			Line:        255,
			Name:        "assignOp",
			MatcherName: "m",
			DocTags: []string{
				"style",
			},
			DocSummary: "Detects assignments that can be simplified by using assignment operators",
			DocBefore:  "x = x * 2",
			DocAfter:   "x *= 2",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           256,
					SyntaxPattern:  "$x = $x + 1",
					ReportTemplate: "replace `$$` with `$x++`",
					WhereExpr: ir.FilterExpr{
						Line:  256,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           257,
					SyntaxPattern:  "$x = $x - 1",
					ReportTemplate: "replace `$$` with `$x--`",
					WhereExpr: ir.FilterExpr{
						Line:  257,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           259,
					SyntaxPattern:  "$x = $x + $y",
					ReportTemplate: "replace `$$` with `$x += $y`",
					WhereExpr: ir.FilterExpr{
						Line:  259,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           260,
					SyntaxPattern:  "$x = $x - $y",
					ReportTemplate: "replace `$$` with `$x -= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  260,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           262,
					SyntaxPattern:  "$x = $x * $y",
					ReportTemplate: "replace `$$` with `$x *= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  262,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           263,
					SyntaxPattern:  "$x = $x / $y",
					ReportTemplate: "replace `$$` with `$x /= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  263,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           264,
					SyntaxPattern:  "$x = $x % $y",
					ReportTemplate: "replace `$$` with `$x %= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  264,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           265,
					SyntaxPattern:  "$x = $x & $y",
					ReportTemplate: "replace `$$` with `$x &= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  265,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           266,
					SyntaxPattern:  "$x = $x | $y",
					ReportTemplate: "replace `$$` with `$x |= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  266,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           267,
					SyntaxPattern:  "$x = $x ^ $y",
					ReportTemplate: "replace `$$` with `$x ^= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  267,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           268,
					SyntaxPattern:  "$x = $x << $y",
					ReportTemplate: "replace `$$` with `$x <<= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  268,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           269,
					SyntaxPattern:  "$x = $x >> $y",
					ReportTemplate: "replace `$$` with `$x >>= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  269,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
				ir.Rule{
					Line:           270,
					SyntaxPattern:  "$x = $x &^ $y",
					ReportTemplate: "replace `$$` with `$x &^= $y`",
					WhereExpr: ir.FilterExpr{
						Line:  270,
						Op:    ir.FilterVarPureOp,
						Src:   "m[\"x\"].Pure",
						Value: "x",
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        277,
			Name:        "preferWriteByte",
			MatcherName: "m",
			DocTags: []string{
				"performance",
				"experimental",
			},
			DocSummary: "Detects WriteRune calls with byte literal argument and reports to use WriteByte instead",
			DocBefore:  "w.WriteRune('\\n')",
			DocAfter:   "w.WriteByte('\\n')",
			Rules: []ir.Rule{
				ir.Rule{
					Line:           278,
					SyntaxPattern:  "$w.WriteRune($c)",
					ReportTemplate: "consider replacing $$ with $w.WriteByte($c)",
					WhereExpr: ir.FilterExpr{
						Line: 279,
						Op:   ir.FilterAndOp,
						Src:  "m[\"w\"].Type.Implements(\"io.ByteWriter\") && (m[\"c\"].Const && m[\"c\"].Value.Int() < 256)",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line:  279,
								Op:    ir.FilterVarTypeImplementsOp,
								Src:   "m[\"w\"].Type.Implements(\"io.ByteWriter\")",
								Value: "w",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  279,
										Op:    ir.FilterStringOp,
										Src:   "\"io.ByteWriter\"",
										Value: "io.ByteWriter",
									},
								},
							},
							ir.FilterExpr{
								Line: 279,
								Op:   ir.FilterAndOp,
								Src:  "(m[\"c\"].Const && m[\"c\"].Value.Int() < 256)",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  279,
										Op:    ir.FilterVarConstOp,
										Src:   "m[\"c\"].Const",
										Value: "c",
									},
									ir.FilterExpr{
										Line: 279,
										Op:   ir.FilterLtOp,
										Src:  "m[\"c\"].Value.Int() < 256",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  279,
												Op:    ir.FilterVarValueIntOp,
												Src:   "m[\"c\"].Value.Int()",
												Value: "c",
											},
											ir.FilterExpr{
												Line:  279,
												Op:    ir.FilterIntOp,
												Src:   "256",
												Value: int64(256),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		ir.RuleGroup{
			Line:        287,
			Name:        "preferFprint",
			MatcherName: "m",
			DocTags: []string{
				"performance",
				"experimental",
			},
			DocSummary: "Detects fmt.Sprint(f|ln) calls which can be replaced with fmt.Fprint(f|ln)",
			DocBefore:  "w.Write([]byte(fmt.Sprintf(\"%x\", 10)))",
			DocAfter:   "fmt.Fprintf(w, \"%x\", 10)",
			Rules: []ir.Rule{
				ir.Rule{
					Line:            296,
					SyntaxPattern:   "$w.Write([]byte($fmt.Sprint($*args)))",
					ReportTemplate:  "fmt.Fprint($w, $args) should be preferred to the $$",
					SuggestTemplate: "fmt.Fprint($w, $args)",
					WhereExpr: ir.FilterExpr{
						Line: 297,
						Op:   ir.FilterAndOp,
						Src:  "m[\"w\"].Type.Implements(\"io.Writer\") &&\n\tm[\"fmt\"].Text == \"fmt\" && m[\"fmt\"].Object.Is(`PkgName`)",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line: 297,
								Op:   ir.FilterAndOp,
								Src:  "m[\"w\"].Type.Implements(\"io.Writer\") &&\n\tm[\"fmt\"].Text == \"fmt\"",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  297,
										Op:    ir.FilterVarTypeImplementsOp,
										Src:   "m[\"w\"].Type.Implements(\"io.Writer\")",
										Value: "w",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  297,
												Op:    ir.FilterStringOp,
												Src:   "\"io.Writer\"",
												Value: "io.Writer",
											},
										},
									},
									ir.FilterExpr{
										Line: 298,
										Op:   ir.FilterEqOp,
										Src:  "m[\"fmt\"].Text == \"fmt\"",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  298,
												Op:    ir.FilterVarTextOp,
												Src:   "m[\"fmt\"].Text",
												Value: "fmt",
											},
											ir.FilterExpr{
												Line:  298,
												Op:    ir.FilterStringOp,
												Src:   "\"fmt\"",
												Value: "fmt",
											},
										},
									},
								},
							},
							ir.FilterExpr{
								Line:  298,
								Op:    ir.FilterVarObjectIsOp,
								Src:   "m[\"fmt\"].Object.Is(`PkgName`)",
								Value: "fmt",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  298,
										Op:    ir.FilterStringOp,
										Src:   "`PkgName`",
										Value: "PkgName",
									},
								},
							},
						},
					},
				},
				ir.Rule{
					Line:            302,
					SyntaxPattern:   "$w.Write([]byte($fmt.Sprintf($*args)))",
					ReportTemplate:  "fmt.Fprintf($w, $args) should be preferred to the $$",
					SuggestTemplate: "fmt.Fprintf($w, $args)",
					WhereExpr: ir.FilterExpr{
						Line: 303,
						Op:   ir.FilterAndOp,
						Src:  "m[\"w\"].Type.Implements(\"io.Writer\") &&\n\tm[\"fmt\"].Text == \"fmt\" && m[\"fmt\"].Object.Is(`PkgName`)",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line: 303,
								Op:   ir.FilterAndOp,
								Src:  "m[\"w\"].Type.Implements(\"io.Writer\") &&\n\tm[\"fmt\"].Text == \"fmt\"",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  303,
										Op:    ir.FilterVarTypeImplementsOp,
										Src:   "m[\"w\"].Type.Implements(\"io.Writer\")",
										Value: "w",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  303,
												Op:    ir.FilterStringOp,
												Src:   "\"io.Writer\"",
												Value: "io.Writer",
											},
										},
									},
									ir.FilterExpr{
										Line: 304,
										Op:   ir.FilterEqOp,
										Src:  "m[\"fmt\"].Text == \"fmt\"",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  304,
												Op:    ir.FilterVarTextOp,
												Src:   "m[\"fmt\"].Text",
												Value: "fmt",
											},
											ir.FilterExpr{
												Line:  304,
												Op:    ir.FilterStringOp,
												Src:   "\"fmt\"",
												Value: "fmt",
											},
										},
									},
								},
							},
							ir.FilterExpr{
								Line:  304,
								Op:    ir.FilterVarObjectIsOp,
								Src:   "m[\"fmt\"].Object.Is(`PkgName`)",
								Value: "fmt",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  304,
										Op:    ir.FilterStringOp,
										Src:   "`PkgName`",
										Value: "PkgName",
									},
								},
							},
						},
					},
				},
				ir.Rule{
					Line:            308,
					SyntaxPattern:   "$w.Write([]byte($fmt.Sprintln($*args)))",
					ReportTemplate:  "fmt.Fprintln($w, $args) should be preferred to the $$",
					SuggestTemplate: "fmt.Fprintln($w, $args)",
					WhereExpr: ir.FilterExpr{
						Line: 309,
						Op:   ir.FilterAndOp,
						Src:  "m[\"w\"].Type.Implements(\"io.Writer\") &&\n\tm[\"fmt\"].Text == \"fmt\" && m[\"fmt\"].Object.Is(`PkgName`)",
						Args: []ir.FilterExpr{
							ir.FilterExpr{
								Line: 309,
								Op:   ir.FilterAndOp,
								Src:  "m[\"w\"].Type.Implements(\"io.Writer\") &&\n\tm[\"fmt\"].Text == \"fmt\"",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  309,
										Op:    ir.FilterVarTypeImplementsOp,
										Src:   "m[\"w\"].Type.Implements(\"io.Writer\")",
										Value: "w",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  309,
												Op:    ir.FilterStringOp,
												Src:   "\"io.Writer\"",
												Value: "io.Writer",
											},
										},
									},
									ir.FilterExpr{
										Line: 310,
										Op:   ir.FilterEqOp,
										Src:  "m[\"fmt\"].Text == \"fmt\"",
										Args: []ir.FilterExpr{
											ir.FilterExpr{
												Line:  310,
												Op:    ir.FilterVarTextOp,
												Src:   "m[\"fmt\"].Text",
												Value: "fmt",
											},
											ir.FilterExpr{
												Line:  310,
												Op:    ir.FilterStringOp,
												Src:   "\"fmt\"",
												Value: "fmt",
											},
										},
									},
								},
							},
							ir.FilterExpr{
								Line:  310,
								Op:    ir.FilterVarObjectIsOp,
								Src:   "m[\"fmt\"].Object.Is(`PkgName`)",
								Value: "fmt",
								Args: []ir.FilterExpr{
									ir.FilterExpr{
										Line:  310,
										Op:    ir.FilterStringOp,
										Src:   "`PkgName`",
										Value: "PkgName",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

