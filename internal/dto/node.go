package dto

type NodeType string

const (
	NodeTypeInitial NodeType = "initial"
	NodeTypeStart   NodeType = "start"
	NodeTypeEnd     NodeType = "end"
	NodeTypeAction  NodeType = "action"
)
