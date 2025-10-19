package papers

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type outlineParser struct{}

func NewOutlineParser() parser.BlockParser {
	return &outlineParser{}
}

func outlineLevelTerminator(level int) byte {
	switch level {
	case 0,1,2:
		return '.'
	case 3,4,5,6:
		return ')'
	default:
		return 0
	}
}

// Trigger returns a list of characters that triggers Parse method of
// this parser.
// If Trigger returns a nil, Open will be called with any lines.
func (o *outlineParser) Trigger() []byte {
	return []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z', '(',
	}
}

// Open parses the current line and returns a result of parsing.
//
// Open must not parse beyond the current line.
// If Open has been able to parse the current line, Open must advance a reader
// position by consumed byte length.
//
// If Open has not been able to parse the current line, Open should returns
// (nil, NoChildren). If Open has been able to parse the current line, Open
// should returns a new Block node and returns HasChildren or NoChildren.
func (o *outlineParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	panic("implement me")
}

// Continue parses the current line and returns a result of parsing.
//
// Continue must not parse beyond the current line.
// If Continue has been able to parse the current line, Continue must advance
// a reader position by consumed byte length.
//
// If Continue has not been able to parse the current line, Continue should
// returns Close. If Continue has been able to parse the current line,
// Continue should returns (Continue | NoChildren) or
// (Continue | HasChildren)
func (o *outlineParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	panic("implement me")
}

// Close will be called when the parser returns Close.
func (o *outlineParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	panic("implement me")
}

// CanInterruptParagraph returns true if the parser can interrupt paragraphs,
// otherwise false.
func (o *outlineParser) CanInterruptParagraph() bool {
	return false
}

// CanAcceptIndentedLine returns true if the parser can open new node when
// the given line is being indented more than 3 spaces.
func (o *outlineParser) CanAcceptIndentedLine() bool {
	return true
}
