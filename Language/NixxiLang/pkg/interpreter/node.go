package interpreter

type Node interface {
	getNodeName() string
}

type VarNode struct {
	Line         int
	VariableName []byte
	Type         VarType
	Value        []byte
}

func (varNode VarNode) getNodeName() string {
	return "VAR_NODE"
}

type MutNode struct {
	Line         int
	VariableName []byte
}

func (mutNode MutNode) getNodeName() string {
	return "MUT_NODE"
}

type IfNode struct {
	Line       int
	LeftValue  any
	RightValue any
	Operator   OperatorValue
	Children   []Node
}

func (ifNode IfNode) getNodeName() string {
	return "IF_NODE"
}

type ElseNode struct {
	Line     int
	Children []Node
}

func (ElseNode ElseNode) getNodeName() string {
	return "ELSE_NODE"
}

type ElseIfNode struct{}

type LoopNode struct{}

type ReturnNode struct {
	Line  int
	From  []byte
	Value any
}

func (returnNode ReturnNode) getNodeName() string {
	return "RETURN_NODE"
}
