package main

import (
	"flag"
	"fmt"
	"os"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/c"
)

var commentedFuncQuery string = `
(
    (comment) @comment .
    (function_definition
        ([(primitive_type) (type_identifier)]) @returns
        (function_declarator
            (identifier) @func_name
            (parameter_list) @func_params
        )
    )
)    
`

func main() {

	fileFlag := flag.String("file", "", "File to load and parse")
	flag.Parse()

	if fileFlag == nil || len(*fileFlag) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cLang := c.GetLanguage()
	bytes, err := os.ReadFile(*fileFlag)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, World!")

	parser := sitter.NewParser()
	parser.SetLanguage(cLang)

	tree := parser.Parse(nil, []byte(bytes))

	root := tree.RootNode()
	//	fmt.Println(root)

	q, _ := sitter.NewQuery([]byte(commentedFuncQuery), cLang)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, root)

	match := 0
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		m = qc.FilterPredicates(m, bytes)
		for i, c := range m.Captures {
			fmt.Println(c.Node.Type())
			fmt.Println(match, i, c.Node.Content(bytes))
		}
		match += 1
	}
}
