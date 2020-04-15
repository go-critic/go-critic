package checker_test

import (
	"regexp"
)

func multiPass() {
	// 1. `[a-a]` -> `[a]`
	// 2. `[a]` -> `a`
	/*! can re-write `[a-a]` as `a` */
	regexp.MustCompile(`[a-a]`)

	// 1. `(?:a|b|c)` -> `(?:[abc])`
	// 2. `(?:[abc])` -> `[abc]`
	/*! can re-write `(?:a|b|c)` as `[abc]` */
	regexp.MustCompile(`(?:a|b|c)`)
}

func altCommonPrefixSuffix() {
	/*! can re-write `foo|fo` as `foo?` */
	regexp.MustCompile(`foo|fo`)

	/*! can re-write `(?:http|https)://` as `(?:https?)://` */
	regexp.MustCompile(`(?:http|https)://`)

	/*! can re-write `xpath|path` as `x?path` */
	regexp.MustCompile(`xpath|path`)

	// Should also work with multi-byte runes.
	/*! can re-write `❤path|path` as `❤?path` */
	regexp.MustCompile(`❤path|path`)
	/*! can re-write `fo|fo❤` as `fo❤?` */
	regexp.MustCompile(`fo|fo❤`)
}

// xx* -> x+
func merge() {
	/*! can re-write `x[abcd][abcd]*y` as `x[abcd]+y` */
	regexp.MustCompile(`x[abcd][abcd]*y`)

	/*! can re-write `axx*y` as `ax+y` */
	regexp.MustCompile(`axx*y`)
}

// (?:x) -> x
func ungroup() {
	/*! can re-write `(?:x)+` as `x+` */
	regexp.MustCompile(`(?:x)+`)

	/*! can re-write `(?:[abc])+` as `[abc]+` */
	regexp.MustCompile(`(?:[abc])+`)
}

// Replaces duplicated expression x with x{n}, when n is a number of duplications.
func repeat() {
	// Always replace several spaces with repetition.
	/*! can re-write `  ` as ` {2}` */
	regexp.MustCompile(`  `)
	/*! can re-write `   ` as ` {3}` */
	regexp.MustCompile(`   `)
	/*! can re-write `    ` as ` {4}` */
	regexp.MustCompile(`    `)

	/*! can re-write `[a-z][a-z]` as `[a-z]{2}` */
	regexp.MustCompile(`[a-z][a-z]`)
	/*! can re-write `[abc][abc][abc]` as `[abc]{3}` */
	regexp.MustCompile(`[abc][abc][abc]`)

	/*! can re-write `(?:foo|bar)(?:foo|bar)` as `(?:foo|bar){2}` */
	regexp.MustCompile(`(?:foo|bar)(?:foo|bar)`)

	/*! can re-write `aaaaax` as `a{5}x` */
	regexp.MustCompile(`aaaaax`)

	/*! can re-write `\d\d\d` as `\d{3}` */
	regexp.MustCompile(`\d\d\d`)
	/*! can re-write `\.\.\.` as `\.{3}` */
	regexp.MustCompile(`\.\.\.`)
	/*! can re-write `\.\.\.\.` as `\.{4}` */
	regexp.MustCompile(`\.\.\.\.`)

	/*! can re-write `....` as `.{4}` */
	regexp.MustCompile(`....`)
	/*! can re-write `.....x` as `.{5}x` */
	regexp.MustCompile(`.....x`)
}

// Replaces the char class with equivalent expression.
func replaceCharClass() {
	/*! can re-write `foo[0-9]+` as `foo\d+` */
	regexp.MustCompile(`foo[0-9]+`)

	/*! can re-write `[0-9]` as `\d` */
	regexp.MustCompile(`[0-9]`)
	/*! can re-write `[[:word:]]` as `\w` */
	regexp.MustCompile(`[[:word:]]`)
	/*! can re-write `[[:^word:]]` as `\W` */
	regexp.MustCompile(`[[:^word:]]`)
	/*! can re-write `[[:digit:]]` as `\d` */
	regexp.MustCompile(`[[:digit:]]`)
	/*! can re-write `[[:^digit:]]` as `\D` */
	regexp.MustCompile(`[[:^digit:]]`)
	/*! can re-write `[[:space:]]` as `\s` */
	regexp.MustCompile(`[[:space:]]`)
	/*! can re-write `[[:^space:]]` as `\S` */
	regexp.MustCompile(`[[:^space:]]`)

	/*! can re-write `[^\D]` as `\d` */
	regexp.MustCompile(`[^\D]`)
	/*! can re-write `[^[:^word:]]` as `\w` */
	regexp.MustCompile(`[^[:^word:]]`)
}

// [x] -> x
func unwrapCharClass() {
	/*! can re-write `[x]` as `x` */
	regexp.MustCompile(`[x]`)
	/*! can re-write `[\d]` as `\d` */
	regexp.MustCompile(`[\d]`)
	/*! can re-write `[]]` as `\]` */
	regexp.MustCompile(`[]]`)
	/*! can re-write `[][]` as `\]\[` */
	regexp.MustCompile(`[][]`)
}

// \# -> #
func unescape() {
	/*! can re-write `\#` as `#` */
	regexp.MustCompile(`\#`)
	/*! can re-write `\&` as `&` */
	regexp.MustCompile(`\&`)
	/*! can re-write `\!` as `!` */
	regexp.MustCompile(`\!`)
	/*! can re-write `\@` as `@` */
	regexp.MustCompile(`\@`)
	/*! can re-write `\%` as `%` */
	regexp.MustCompile(`\%`)
	/*! can re-write `\>` as `>` */
	regexp.MustCompile(`\>`)
	/*! can re-write `\<` as `<` */
	regexp.MustCompile(`\<`)
	/*! can re-write `\:` as `:` */
	regexp.MustCompile(`\:`)
	/*! can re-write `\;` as `;` */
	regexp.MustCompile(`\;`)
	/*! can re-write `\/` as `/` */
	regexp.MustCompile(`\/`)
	/*! can re-write `\,` as `,` */
	regexp.MustCompile(`\,`)
	/*! can re-write `\=` as `=` */
	regexp.MustCompile(`\=`)

	/*! can re-write `[x\#]` as `[x#]` */
	regexp.MustCompile(`[x\#]`)
	/*! can re-write `[x\.]` as `[x.]` */
	regexp.MustCompile(`[x\.]`)
}

// a|b|c -> [abc]
func charAlt() {
	/*! can re-write `(a|b|c|d)` as `([abcd])` */
	regexp.MustCompile(`(a|b|c|d)`)

	/*! can re-write `a|b` as `[ab]` */
	regexp.MustCompile(`a|b`)
}

// [a-a] -> [a]
// [a-b] -> [ab]
// [a-c] -> [abc]
func unrangeCharClass() {
	/*! can re-write `[xa-a]` as `[xa]` */
	regexp.MustCompile(`[xa-a]`)
	/*! can re-write `[xa-b]` as `[xab]` */
	regexp.MustCompile(`[xa-b]`)
	/*! can re-write `[xa-c]` as `[xabc]` */
	regexp.MustCompile(`[xa-c]`)

	/*! can re-write `[x1-1]` as `[x1]` */
	regexp.MustCompile(`[x1-1]`)
	/*! can re-write `[x1-2]` as `[x12]` */
	regexp.MustCompile(`[x1-2]`)
	/*! can re-write `[x1-3]` as `[x123]` */
	regexp.MustCompile(`[x1-3]`)

	/*! can re-write `[1-3a-c]` as `[123abc]` */
	regexp.MustCompile(`[1-3a-c]`)
}

// x{0,1} -> x?
// x{1,}  -> x+
// x{0,}  -> x*
// x{1}   -> x
// x{0}   ->
func unrepeat() {
	/*! can re-write `x{0}foo` as `foo` */
	regexp.MustCompile(`x{0}foo`)

	/*! can re-write `x{1}` as `x` */
	regexp.MustCompile(`x{1}`)

	/*! can re-write `[abc]{1}` as `[abc]` */
	regexp.MustCompile(`[abc]{1}`)

	/*! can re-write `[0-9]{1,}` as `\d+` */
	regexp.MustCompile(`[0-9]{1,}`)

	/*! can re-write `[0-9]{0,}` as `\d*` */
	regexp.MustCompile(`[0-9]{0,}`)

	/*! can re-write `[0-9]{0,1}` as `\d?` */
	regexp.MustCompile(`[0-9]{0,1}`)
}

func mixed() {
	/*! can re-write `(https?:\/\/[^\s]+)` as `(https?://\S+)` */
	regexp.MustCompile(`(https?:\/\/[^\s]+)`)

	/*! can re-write `((:|ː)\w+(:|ː))` as `(([:ː])\w+([:ː]))` */
	regexp.MustCompile(`((:|ː)` + `\w+(:|ː))`)

	/*! can re-write `^[^\s]+\.[^\s]+$` as `^\S+\.\S+$` */
	regexp.MustCompile(`^[^\s]+\.[^\s]+$`)

	/*! can re-write `(^[.]{1})|([.]{1}$)|([.]{2,})` as `(^[.])|([.]$)|([.]{2,})` */
	regexp.MustCompile(`(^[.]{1})|([.]{1}$)|([.]{2,})`)

	/*! can re-write `'''''(.*)'''''` as `'{5}(.*)'{5}` */
	regexp.MustCompile(`'''''(.*)'''''`)

	/*! can re-write `o|O` as `[oO]` */
	regexp.MustCompile(`o|O`)

	/*! can re-write `(-(c|e)e(?:-\d+)?)$` as `(-([ce])e(?:-\d+)?)$` */
	regexp.MustCompile(`(-(c|e)e(?:-\d+)?)$`)

	/*! can re-write `^[a-z]+\[[0-9]+\]$` as `^[a-z]+\[\d+\]$` */
	regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	/*! can re-write `Copyright.*(\d{4}),?\s([\w -!]*\w)` as `Copyright.*(\d{4}),?\s([\w !]*\w)` */
	regexp.MustCompile(`Copyright.*(\d{4}),?\s([\w -!]*\w)`)

	/*! can re-write `^(?i)(\s*\w+(\.\w+){0,1}\s*)` as `^(?i)(\s*\w+(\.\w+)?\s*)` */
	regexp.MustCompile(`^(?i)(\s*\w+(\.\w+){0,1}\s*)`)

	/*! can re-write `("|'|“|”|’|«|»)` as `(["'“”’«»])` */
	regexp.MustCompile(`("|'|“|”|’|«|»)`)

	/*! can re-write `\p{Han}|[\w]+` as `\p{Han}|\w+` */
	regexp.MustCompile(`\p{Han}|[\w]+`)

	/*! can re-write `\(?[\w\-\.\[\]]\)?` as `\(?[\w\-.\[\]]\)?` */
	regexp.MustCompile(`\(?[\w\-\.\[\]]\)?`)
}
