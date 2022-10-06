package checker_test

import (
	"regexp"
)

func suspiciousCharRange() {
	/*! suspicious char range `$-%` in [$-%] */
	regexp.MustCompile(`[$-%]`)

	/*! suspicious char range ` -!` in [ -!] */
	regexp.MustCompile(`[ -!]`)

	/*! suspicious char range `❤-❥` in [❤-❥] */
	regexp.MustCompile(`[❤-❥]`)

	/*! suspicious char range `,-.` in [0-9,-.] */
	regexp.MustCompile("\\\\?(?:{[0-9,-.]*}|{q})")
}

func altDups() {
	/*! `x` is duplicated in x|x */
	regexp.MustCompile(`x|x`)

	/*! `[a-z]` is duplicated in [a-z]|[a-z]|[0-9] */
	regexp.MustCompile(`([a-z]|[a-z]|[0-9])`)
}

func charClassDuplicates() {
	/*! `a` is duplicated in [aba] */
	regexp.MustCompile(`x[aba]y`)

	/*! `\141` intersects with `a` in [\141a] */
	regexp.MustCompile(`[\141a]`)
	/*! `a` intersects with `\x61` in [a\x61] */
	regexp.MustCompile(`[a\x61]`)
	/*! `a` intersects with `\x{61}` in [^a\x{61}] */
	regexp.MustCompile(`[^a\x{61}]`)

	/*! `a-c` intersects with `b` in [a-cb] */
	regexp.MustCompile(`[a-cb]`)
	/*! `a-b` is duplicated in [^a-ba-b] */
	regexp.MustCompile(`[^a-ba-b]`)

	/*! `\x{61}-\x{63}` intersects with `c` in [\x{61}-\x{63}c] */
	regexp.MustCompile(`[\x{61}-\x{63}c]`)
	/*! `\x61-\x63` intersects with `c` in [\x61-\x63c] */
	regexp.MustCompile(`[\x61-\x63c]`)
	/*! `\141-\143` intersects with `c` in [\141-\143c] */
	regexp.MustCompile(`[\141-\143c]`)

	/*! `\d` intersects with `5` in [\d5] */
	regexp.MustCompile(`[\d5]`)
	/*! `\d` intersects with `5-6` in [5-6\d] */
	regexp.MustCompile(`[5-6\d]`)
	/*! `\w` intersects with `_` in [\w_] */
	regexp.MustCompile(`[\w_]`)
	/*! `\w` intersects with `a-d` in [\w%a-d] */
	regexp.MustCompile(`[\w%a-d]`)
	/*! `\D` intersects with `g` in [\Dg] */
	regexp.MustCompile(`[\Dg]`)
	/*! `\D` intersects with `❤` in [\D❤5] */
	regexp.MustCompile(`[\D❤5]`)
	/*! `\s` intersects with `\t` in [\s\t] */
	regexp.MustCompile(`[\s\t]`)
	/*! `\s` intersects with `\n` in [\n\s] */
	regexp.MustCompile(`[\n\s]`)

	/*! `1-5` intersects with `2-3` in [1-52-34] */
	regexp.MustCompile(`[1-52-34]`)
	/*! `1-5` intersects with `2-3` in [42-31-5] */
	regexp.MustCompile(`[42-31-5]`)

	/*! `\W` intersects with `❤` in [\w\W❤] */
	regexp.MustCompile(`[\w\W❤]`)

	/*! `|` is duplicated in [a|b|m|k] */
	regexp.MustCompile("(?i)\x1b\\[([0-9]{1,2}(;[0-9]{1,2})?)?[a|b|m|k]")

	/*! `\w` intersects with `_` in [\w!#$%&'*+/=?^_`{|}~-] */
	regexp.MustCompile("^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.+a)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$")
}

func repeatedQuantifier() {
	/*! repeated greedy quantifier in (a+)+ */
	regexp.MustCompile(`(a+)+`)
	/*! repeated greedy quantifier in (?:[ab]*)+ */
	regexp.MustCompile(`(?:[ab]*)+`)
	/*! repeated greedy quantifier in ((ab)+)* */
	regexp.MustCompile(`((ab)+)*`)
}

func redundantFlags() {
	/*! redundant flag m in (?m) */
	regexp.MustCompile(`(?m)(?m)`)

	/*! redundant flag i in (?i:foo) */
	regexp.MustCompile(`(?ims:(?i:foo))(?im:bar)`)

	/*! redundant flag i in (?ims:flags1) */
	regexp.MustCompile(`(?i)(?ims:flags1)(?m:flags2)`)

	/*! redundant flag m in (?m:a|b(?s:foo)) */
	regexp.MustCompile(`((?m)(?m:a|b(?s:foo))(?i)x)`)
}

func flagClear() {
	/*! clearing unset flag i in (?-i) */
	regexp.MustCompile(`(?-i)x`)

	/*! clearing unset flag i in (?-i) */
	regexp.MustCompile(`(?i:foo)(?-i)bar`)

	/*! clearing unset flag m in (?-mi) */
	/*! clearing unset flag i in (?-mi) */
	regexp.MustCompile(`(?i:(?m:fo(?-i)o))(?-mi)bar`)

	/*! clearing unset flag i in (?i-ii) */
	regexp.MustCompile(`(?i-ii)`)

	/*! clearing unset flag i in (?-i) */
	regexp.MustCompile(`(?:(?i)foo)(?-i)`)

	/*! clearing unset flag i in (?-i) */
	regexp.MustCompile(`((?i)(?-i))(?-i)`)

	/*! clearing unset flag i in (?-i) */
	regexp.MustCompile(`(?:(?i)(?-i))(?-i)`)

	/*! clearing unset flag s in (?m-s) */
	regexp.MustCompile(`(?m-s)(?:tags)/(\S+)$`)
}

func suspiciousAltAnchor() {
	/*! ^ applied only to `foo` in ^foo|bar|baz */
	regexp.MustCompile(`^foo|bar|baz`)

	/*! ^ applied only to `a` in ^a|bc */
	regexp.MustCompile(`(?:^a|bc)c`)

	/*! $ applied only to `baz` in foo|bar|baz$ */
	regexp.MustCompile(`foo|bar|baz$`)

	/*! $ applied only to `bc` in a|bc$ */
	regexp.MustCompile(`c(?:a|bc$)`)
}

func badRegexps() {
	/*! suspicious char range `=-_` in [/.@!~#$%^&*:";?\\+=-_,{}\[\]<>！￥…（）—=、“”：；？。，《》] */
	/*! `=-_` intersects with `=` in [/.@!~#$%^&*:";?\\+=-_,{}\[\]<>！￥…（）—=、“”：；？。，《》] */
	regexp.MustCompile(`[/.@!~#$%^&*:";?\\+=-_,{}\[\]<>！￥…（）—=、“”：；？。，《》]`)

	/*! `e` is duplicated in [com|org|edu|net] */
	regexp.MustCompile(`^(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)

	/*! `\s` intersects with `\t` in [\s\t] */
	regexp.MustCompile(`(?m)^([\s\t]*)([\*\-\+]|\d\.)\s+`)

	/*! suspicious char range `%-\/` in [a-z0-9_.?&=%-\/] */
	regexp.MustCompile(`^[^\/][a-z0-9_.?&=%-\/]+$`)

	/*! suspicious char range `+-\.` in [a-z0-9+-\.] */
	regexp.MustCompile(`^([a-z][a-z0-9+-\.]*):(\/\/)?.+$`)

	/*! `_` is duplicated in [a-zA-Z\.\-_0-9_] */
	regexp.MustCompile(`^[a-zA-Z\.\-_0-9_]+$`)

	/*! `\w` intersects with `\d` in [!\w\d\.\+\-] */
	regexp.MustCompile(`^https?://itunes.apple.com/(?:(\w+)/)?app/(?:[!\w\d\.\+\-]+/)?id(\d+)`)

	/*! `\w` intersects with `\d` in [_=\w\d\.&;] */
	/*! `\w` intersects with `\d` in [_\w\d\.] */
	regexp.MustCompile(`^https?://play.google.com/store/apps/details\?(?:[_=\w\d\.&;]*[;|&])?id=([_\w\d\.]+)`)

	/*! $ applied only to `b` in a|b$ */
	regexp.MustCompile(`^(?:(?:https?:\/\/)?google\.com)?\/(a|b$)`)
}

func danglingAnchor() {
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`a^`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`a^b`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`^^foo`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`foo?|bar^`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`(?i:a)^foo`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`(?i)^(?:foo|bar|^baz)`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`(?i)^(?m)foobar^baz`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`(?i:foo|((?:f|b|(foo|^bar)^)))`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile(`(?i)(?m)\n^foo|bar|baz`)
	/*! dangling or redundant ^, maybe \^ is intended? */
	regexp.MustCompile("(?ms)^```(?:(?P<type>yaml)\\w*\\n(?P<content>.+?)|\\w*\\n(?P<content>\\{.+?\\}))\\n^```")
}
