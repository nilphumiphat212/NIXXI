package interpreter

func ParserMultiLine(tokens *[][]Token) []Node {
	var nodes []Node
	for line, tokenInLine := range *tokens {
		for _, token := range tokenInLine {
			nodes = append(nodes, VarNode{
				Line:         line,
				VariableName: []byte(token.Type),
				Type:         VARTYPE_NULL,
				Value:        []byte(token.Value),
			})
		}
	}
	return nodes
}
