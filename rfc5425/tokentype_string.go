// Code generated by "stringer -type TokenType rfc5425/tokens.go"; DO NOT EDIT.

package rfc5425

import "strconv"

const _TokenType_name = "ILLEGALEOFWSMSGLENSYSLOGMSG"

var _TokenType_index = [...]uint8{0, 7, 10, 12, 18, 27}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}