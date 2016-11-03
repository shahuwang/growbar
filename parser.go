//line growbar.y:1
package main

import __yyfmt__ "fmt"

//line growbar.y:3
import (
	"fmt"
)

//line growbar.y:7
type GrowSymType struct {
	yys        int
	identifier string
	expression *Expression
}

const INT_LITERAL = 57346
const DOUBLE_LITERAL = 57347
const STRING_LITERAL = 57348
const IDENTIFIER = 57349
const FUNCTION = 57350
const IF = 57351
const ELSE = 57352
const ELSIF = 57353
const WHILE = 57354
const FOR = 57355
const RETURN_T = 57356
const BREAK = 57357
const CONTINUE = 57358
const NULL_T = 57359
const LP = 57360
const RP = 57361
const LC = 57362
const RC = 57363
const SEMICOLON = 57364
const COMMA = 57365
const ASSIGN = 57366
const LOGICAL_AND = 57367
const LOGICAL_OR = 57368
const EQ = 57369
const NE = 57370
const GT = 57371
const GE = 57372
const LT = 57373
const LE = 57374
const ADD = 57375
const SUB = 57376
const MUL = 57377
const DIV = 57378
const MOD = 57379
const TRUE_T = 57380
const FALSE_T = 57381
const GLOBAL_T = 57382

var GrowToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INT_LITERAL",
	"DOUBLE_LITERAL",
	"STRING_LITERAL",
	"IDENTIFIER",
	"FUNCTION",
	"IF",
	"ELSE",
	"ELSIF",
	"WHILE",
	"FOR",
	"RETURN_T",
	"BREAK",
	"CONTINUE",
	"NULL_T",
	"LP",
	"RP",
	"LC",
	"RC",
	"SEMICOLON",
	"COMMA",
	"ASSIGN",
	"LOGICAL_AND",
	"LOGICAL_OR",
	"EQ",
	"NE",
	"GT",
	"GE",
	"LT",
	"LE",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"TRUE_T",
	"FALSE_T",
	"GLOBAL_T",
}
var GrowStatenames = [...]string{}

const GrowEofCode = 1
const GrowErrCode = 2
const GrowInitialStackSize = 16

//line growbar.y:145

//line yacctab:1
var GrowExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GrowNprod = 33
const GrowPrivate = 57344

var GrowTokenNames []string
var GrowStates []string

const GrowLast = 37

var GrowAct = [...]int{

	17, 18, 32, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 2, 1, 33, 0, 0, 0, 0,
	10, 8, 9, 11, 12, 13, 14, 15, 16, 3,
	4, 5, 6, 7, 29, 30, 31,
}
var GrowPact = [...]int{

	-4, -4, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000,
}
var GrowPgo = [...]int{

	0, 14, 13,
}
var GrowR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2,
}
var GrowR2 = [...]int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1,
}
var GrowChk = [...]int{

	-1000, -1, -2, 33, 34, 35, 36, 37, 25, 26,
	24, 27, 28, 29, 30, 31, 32, 4, 5, 7,
	8, 9, 10, 11, 12, 13, 14, 15, 16, 38,
	39, 40, 6, -2,
}
var GrowDef = [...]int{

	0, -2, 1, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 2,
}
var GrowTok1 = [...]int{

	1,
}
var GrowTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40,
}
var GrowTok3 = [...]int{
	0,
}

var GrowErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	GrowDebug        = 0
	GrowErrorVerbose = false
)

type GrowLexer interface {
	Lex(lval *GrowSymType) int
	Error(s string)
}

type GrowParser interface {
	Parse(GrowLexer) int
	Lookahead() int
}

type GrowParserImpl struct {
	lval  GrowSymType
	stack [GrowInitialStackSize]GrowSymType
	char  int
}

func (p *GrowParserImpl) Lookahead() int {
	return p.char
}

func GrowNewParser() GrowParser {
	return &GrowParserImpl{}
}

const GrowFlag = -1000

func GrowTokname(c int) string {
	if c >= 1 && c-1 < len(GrowToknames) {
		if GrowToknames[c-1] != "" {
			return GrowToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func GrowStatname(s int) string {
	if s >= 0 && s < len(GrowStatenames) {
		if GrowStatenames[s] != "" {
			return GrowStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func GrowErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !GrowErrorVerbose {
		return "syntax error"
	}

	for _, e := range GrowErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + GrowTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := GrowPact[state]
	for tok := TOKSTART; tok-1 < len(GrowToknames); tok++ {
		if n := base + tok; n >= 0 && n < GrowLast && GrowChk[GrowAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if GrowDef[state] == -2 {
		i := 0
		for GrowExca[i] != -1 || GrowExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; GrowExca[i] >= 0; i += 2 {
			tok := GrowExca[i]
			if tok < TOKSTART || GrowExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if GrowExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += GrowTokname(tok)
	}
	return res
}

func Growlex1(lex GrowLexer, lval *GrowSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = GrowTok1[0]
		goto out
	}
	if char < len(GrowTok1) {
		token = GrowTok1[char]
		goto out
	}
	if char >= GrowPrivate {
		if char < GrowPrivate+len(GrowTok2) {
			token = GrowTok2[char-GrowPrivate]
			goto out
		}
	}
	for i := 0; i < len(GrowTok3); i += 2 {
		token = GrowTok3[i+0]
		if token == char {
			token = GrowTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = GrowTok2[1] /* unknown char */
	}
	if GrowDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", GrowTokname(token), uint(char))
	}
	return char, token
}

func GrowParse(Growlex GrowLexer) int {
	return GrowNewParser().Parse(Growlex)
}

func (Growrcvr *GrowParserImpl) Parse(Growlex GrowLexer) int {
	var Grown int
	var GrowVAL GrowSymType
	var GrowDollar []GrowSymType
	_ = GrowDollar // silence set and not used
	GrowS := Growrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Growstate := 0
	Growrcvr.char = -1
	Growtoken := -1 // Growrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Growstate = -1
		Growrcvr.char = -1
		Growtoken = -1
	}()
	Growp := -1
	goto Growstack

ret0:
	return 0

ret1:
	return 1

Growstack:
	/* put a state and value onto the stack */
	if GrowDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", GrowTokname(Growtoken), GrowStatname(Growstate))
	}

	Growp++
	if Growp >= len(GrowS) {
		nyys := make([]GrowSymType, len(GrowS)*2)
		copy(nyys, GrowS)
		GrowS = nyys
	}
	GrowS[Growp] = GrowVAL
	GrowS[Growp].yys = Growstate

Grownewstate:
	Grown = GrowPact[Growstate]
	if Grown <= GrowFlag {
		goto Growdefault /* simple state */
	}
	if Growrcvr.char < 0 {
		Growrcvr.char, Growtoken = Growlex1(Growlex, &Growrcvr.lval)
	}
	Grown += Growtoken
	if Grown < 0 || Grown >= GrowLast {
		goto Growdefault
	}
	Grown = GrowAct[Grown]
	if GrowChk[Grown] == Growtoken { /* valid shift */
		Growrcvr.char = -1
		Growtoken = -1
		GrowVAL = Growrcvr.lval
		Growstate = Grown
		if Errflag > 0 {
			Errflag--
		}
		goto Growstack
	}

Growdefault:
	/* default state action */
	Grown = GrowDef[Growstate]
	if Grown == -2 {
		if Growrcvr.char < 0 {
			Growrcvr.char, Growtoken = Growlex1(Growlex, &Growrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if GrowExca[xi+0] == -1 && GrowExca[xi+1] == Growstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Grown = GrowExca[xi+0]
			if Grown < 0 || Grown == Growtoken {
				break
			}
		}
		Grown = GrowExca[xi+1]
		if Grown < 0 {
			goto ret0
		}
	}
	if Grown == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Growlex.Error(GrowErrorMessage(Growstate, Growtoken))
			Nerrs++
			if GrowDebug >= 1 {
				__yyfmt__.Printf("%s", GrowStatname(Growstate))
				__yyfmt__.Printf(" saw %s\n", GrowTokname(Growtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Growp >= 0 {
				Grown = GrowPact[GrowS[Growp].yys] + GrowErrCode
				if Grown >= 0 && Grown < GrowLast {
					Growstate = GrowAct[Grown] /* simulate a shift of "error" */
					if GrowChk[Growstate] == GrowErrCode {
						goto Growstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if GrowDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", GrowS[Growp].yys)
				}
				Growp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if GrowDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", GrowTokname(Growtoken))
			}
			if Growtoken == GrowEofCode {
				goto ret1
			}
			Growrcvr.char = -1
			Growtoken = -1
			goto Grownewstate /* try again in the same state */
		}
	}

	/* reduction by production Grown */
	if GrowDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Grown, GrowStatname(Growstate))
	}

	Grownt := Grown
	Growpt := Growp
	_ = Growpt // guard against "declared and not used"

	Growp -= GrowR2[Grown]
	// Growp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Growp+1 >= len(GrowS) {
		nyys := make([]GrowSymType, len(GrowS)*2)
		copy(nyys, GrowS)
		GrowS = nyys
	}
	GrowVAL = GrowS[Growp+1]

	/* consult goto table to find next state */
	Grown = GrowR1[Grown]
	Growg := GrowPgo[Grown]
	Growj := Growg + GrowS[Growp].yys + 1

	if Growj >= GrowLast {
		Growstate = GrowAct[Growg]
	} else {
		Growstate = GrowAct[Growj]
		if GrowChk[Growstate] != -Grown {
			Growstate = GrowAct[Growg]
		}
	}
	// dummy call; replaced with literal code
	switch Grownt {

	case 3:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:26
		{
			fmt.Println("is ADD")
		}
	case 4:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:30
		{
			fmt.Println("is SUB")
		}
	case 5:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:34
		{
			fmt.Println("is MUL")
		}
	case 6:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:38
		{
			fmt.Println("is DIV")
		}
	case 7:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:42
		{
			fmt.Println("is MOD")
		}
	case 8:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:46
		{
			fmt.Println("is AND")
		}
	case 9:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:50
		{
			fmt.Println("is OR")
		}
	case 10:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:54
		{
			fmt.Println("is ASSIGN")
		}
	case 11:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:58
		{
			fmt.Println("is EQ")
		}
	case 12:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:62
		{
			fmt.Println("is NE")
		}
	case 13:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:66
		{
			fmt.Println("is GT")
		}
	case 14:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:70
		{
			fmt.Println("is GE")
		}
	case 15:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:74
		{
			fmt.Println("is LT")
		}
	case 16:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:78
		{
			fmt.Println("is LE")
		}
	case 17:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:82
		{
			fmt.Println(GrowDollar[1].expression.int_value)
		}
	case 18:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:86
		{
			fmt.Println(GrowDollar[1].expression.double_value)
		}
	case 19:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:90
		{
			fmt.Println(GrowDollar[1].identifier)
		}
	case 20:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:94
		{
			fmt.Println("is FUNCTION")
		}
	case 21:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:98
		{
			fmt.Println("is IF")
		}
	case 22:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:102
		{
			fmt.Println("is ELSE")
		}
	case 23:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:106
		{
			fmt.Println("is ELSIF")
		}
	case 24:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:110
		{
			fmt.Println("is WHILE")
		}
	case 25:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:114
		{
			fmt.Println("is FOR")
		}
	case 26:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:118
		{
			fmt.Println("is RETURN_T")
		}
	case 27:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:122
		{
			fmt.Println("is BREAK")
		}
	case 28:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:126
		{
			fmt.Println("is CONTINUE")
		}
	case 29:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:130
		{
			fmt.Println("is TRUE_T")
		}
	case 30:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:134
		{
			fmt.Println("is FALSE_T")
		}
	case 31:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:138
		{
			fmt.Println("is GLOBAL_T")
		}
	case 32:
		GrowDollar = GrowS[Growpt-1 : Growpt+1]
		//line growbar.y:142
		{
			fmt.Println(GrowDollar[1].expression.string_value)
		}
	}
	goto Growstack /* stack new state and value */
}
