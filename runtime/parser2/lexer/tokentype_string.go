// Code generated by "stringer -type=TokenType -trimprefix Token"; DO NOT EDIT.

package lexer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TokenError-0]
	_ = x[TokenEOF-1]
	_ = x[TokenSpace-2]
	_ = x[TokenNumber-3]
	_ = x[TokenIdentifier-4]
	_ = x[TokenString-5]
	_ = x[TokenPlus-6]
	_ = x[TokenMinus-7]
	_ = x[TokenStar-8]
	_ = x[TokenSlash-9]
	_ = x[TokenNilCoalesce-10]
	_ = x[TokenParenOpen-11]
	_ = x[TokenParenClose-12]
	_ = x[TokenBraceOpen-13]
	_ = x[TokenBraceClose-14]
	_ = x[TokenBracketOpen-15]
	_ = x[TokenBracketClose-16]
	_ = x[TokenQuestionMark-17]
	_ = x[TokenComma-18]
	_ = x[TokenColon-19]
	_ = x[TokenLeftArrow-20]
	_ = x[TokenLess-21]
	_ = x[TokenGreater-22]
	_ = x[TokenBlockCommentStart-23]
	_ = x[TokenBlockCommentContent-24]
	_ = x[TokenBlockCommentEnd-25]
}

const _TokenType_name = "ErrorEOFSpaceNumberIdentifierStringPlusMinusStarSlashNilCoalesceParenOpenParenCloseBraceOpenBraceCloseBracketOpenBracketCloseQuestionMarkCommaColonLeftArrowLessGreaterBlockCommentStartBlockCommentContentBlockCommentEnd"

var _TokenType_index = [...]uint8{0, 5, 8, 13, 19, 29, 35, 39, 44, 48, 53, 64, 73, 83, 92, 102, 113, 125, 137, 142, 147, 156, 160, 167, 184, 203, 218}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
