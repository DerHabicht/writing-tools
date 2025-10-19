package papers

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type Point interface {
	// PTT returns the point type (bullet or outline point).
	PTT() PTT
	
	// Level returns the nesting level of this point in the tree.
	Level() int
	
	// RenderLaTeX renders this point as a LaTeX fragment.
	// This implements the LaTeXRenderer interface.
	RenderLaTeX() []byte

	// Subpoints returns all children nodes as points.
	Subpoints() []Point
	
	// Type returns a type of this node.
	Type() ast.NodeType

	// Kind returns a kind of this node.
	Kind() ast.NodeKind

	// NextSibling returns a next sibling node of this node.
	NextSibling() ast.Node

	// PreviousSibling returns a previous sibling node of this node.
	PreviousSibling() ast.Node

	// Parent returns a parent node of this node.
	Parent() ast.Node

	// SetParent sets a parent node to this node.
	SetParent(ast.Node)

	// SetPreviousSibling sets a previous sibling node to this node.
	SetPreviousSibling(ast.Node)

	// SetNextSibling sets a next sibling node to this node.
	SetNextSibling(ast.Node)

	// HasChildren returns true if this node has any children, otherwise false.
	HasChildren() bool

	// ChildCount returns a total number of children.
	ChildCount() int

	// FirstChild returns a first child of this node.
	FirstChild() ast.Node

	// LastChild returns a last child of this node.
	LastChild() ast.Node

	// AppendChild append a node child to the tail of the children.
	AppendChild(self, child ast.Node)

	// RemoveChild removes a node child from this node.
	// If a node child is not children of this node, RemoveChild nothing to do.
	RemoveChild(self, child ast.Node)

	// RemoveChildren removes all children from this node.
	RemoveChildren(self ast.Node)

	// SortChildren sorts children by comparator.
	SortChildren(comparator func(n1, n2 ast.Node) int)

	// ReplaceChild replace a node v1 with a node insertee.
	// If v1 is not children of this node, ReplaceChild append a insetee to the
	// tail of the children.
	ReplaceChild(self, v1, insertee ast.Node)

	// InsertBefore inserts a node insertee before a node v1.
	// If v1 is not children of this node, InsertBefore append a insetee to the
	// tail of the children.
	InsertBefore(self, v1, insertee ast.Node)

	// InsertAfterinserts a node insertee after a node v1.
	// If v1 is not children of this node, InsertBefore append a insetee to the
	// tail of the children.
	InsertAfter(self, v1, insertee ast.Node)

	// OwnerDocument returns this node's owner document.
	// If this node is not a child of the Document node, OwnerDocument
	// returns nil.
	OwnerDocument() *ast.Document

	// Dump dumps an AST tree structure to stdout.
	// This function completely aimed for debugging.
	// level is a indent level. Implementer should indent informations with
	// 2 * level spaces.
	Dump(source []byte, level int)

	// Text returns text values of this node.
	// This method is valid only for some inline nodes.
	// If this node is a block node, Text returns a text value as reasonable as possible.
	// Notice that there are no 'correct' text values for the block nodes.
	// Result for the block nodes may be different from your expectation.
	//
	// Deprecated: Use other properties of the node to get the text value(i.e. Pragraph.Lines, Text.Value).
	Text(source []byte) []byte

	// HasBlankPreviousLines returns true if the row before this node is blank,
	// otherwise false.
	// This method is valid only for block nodes.
	HasBlankPreviousLines() bool

	// SetBlankPreviousLines sets whether the row before this node is blank.
	// This method is valid only for block nodes.
	SetBlankPreviousLines(v bool)

	// Lines returns text segments that hold positions in a source.
	// This method is valid only for block nodes.
	Lines() *text.Segments

	// SetLines sets text segments that hold positions in a source.
	// This method is valid only for block nodes.
	SetLines(*text.Segments)

	// IsRaw returns true if contents should be rendered as 'raw' contents.
	IsRaw() bool

	// SetAttribute sets the given value to the attributes.
	SetAttribute(name []byte, value interface{})

	// SetAttributeString sets the given value to the attributes.
	SetAttributeString(name string, value interface{})

	// Attribute returns a (attribute value, true) if an attribute
	// associated with the given name is found, otherwise
	// (nil, false)
	Attribute(name []byte) (interface{}, bool)

	// AttributeString returns a (attribute value, true) if an attribute
	// associated with the given name is found, otherwise
	// (nil, false)
	AttributeString(name string) (interface{}, bool)

	// Attributes returns a list of attributes.
	// This may be a nil if there are no attributes.
	Attributes() []ast.Attribute

	// RemoveAttributes removes all attributes from this node.
	RemoveAttributes()
}
