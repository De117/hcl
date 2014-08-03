package ast

// ValueType is an enum represnting the type of a value in
// a LiteralNode.
type ValueType byte

const (
	ValueTypeUnknown ValueType = iota
	ValueTypeInt
	ValueTypeString
	ValueTypeBool
	ValueTypeNil
)

// Node is implemented by all AST nodes for HCL.
type Node interface {
	Accept(Visitor)
}

// KeyedNode is a node that has a key associated with it.
type KeyedNode interface {
	Node

	Key() string
}

// Visitor is the interface that must be implemented by any
// structures who want to be visited as part of the visitor pattern
// on the AST.
type Visitor interface {
	Visit(Node)
}

// ObjectNode represents an object that has multiple elements.
// An object's elements may repeat (keys). This is expected to
// be validated/removed at a semantic check, rather than at a
// syntax level.
type ObjectNode struct {
	K    string
	Elem []KeyedNode
}

// AssignmentNode represents a direct assignment with an equals
// sign.
type AssignmentNode struct {
	K     string
	Value Node
}

// ListNode represents a list or array of items.
type ListNode struct {
	Elem []Node
}

// LiteralNode is a direct value.
type LiteralNode struct {
	Type  ValueType
	Value interface{}
}

func (n ObjectNode) Accept(v Visitor) {
	v.Visit(n)

	for _, e := range n.Elem {
		e.Accept(v)
	}
}

func (n ObjectNode) Key() string {
	return n.K
}

func (n AssignmentNode) Accept(v Visitor) {
	v.Visit(n)
	n.Value.Accept(v)
}

func (n AssignmentNode) Key() string {
	return n.K
}

func (n ListNode) Accept(v Visitor) {
	v.Visit(n)

	for _, e := range n.Elem {
		e.Accept(v)
	}
}

func (n LiteralNode) Accept(v Visitor) {
	v.Visit(n)
}