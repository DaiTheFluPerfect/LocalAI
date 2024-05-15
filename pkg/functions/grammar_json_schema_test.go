package functions_test

import (
	"strings"

	. "github.com/go-skynet/LocalAI/pkg/functions"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testFunctions = []ItemFunction{
	{
		Type: "object",
		Properties: FunctionProperties{
			Function: FunctionName{
				Const: "create_event",
			},
			Arguments: Argument{ // this is OpenAI's parameter
				Type: "object",
				Properties: map[string]interface{}{
					"title": map[string]string{"type": "string"},
					"date":  map[string]string{"type": "string"},
					"time":  map[string]string{"type": "string"},
				},
			},
		},
	},
	{
		Type: "object",
		Properties: FunctionProperties{
			Function: FunctionName{
				Const: "search",
			},
			Arguments: Argument{
				Type: "object",
				Properties: map[string]interface{}{
					"query": map[string]string{"type": "string"},
				},
			},
		},
	},
}

var testFunctionsName = []ItemName{
	{
		Type: "object",
		Properties: NameProperties{
			Function: FunctionName{
				Const: "create_event",
			},
			Arguments: Argument{ // this is OpenAI's parameter
				Type: "object",
				Properties: map[string]interface{}{
					"title": map[string]string{"type": "string"},
					"date":  map[string]string{"type": "string"},
					"time":  map[string]string{"type": "string"},
				},
			},
		},
	},
	{
		Type: "object",
		Properties: NameProperties{
			Function: FunctionName{
				Const: "search",
			},
			Arguments: Argument{
				Type: "object",
				Properties: map[string]interface{}{
					"query": map[string]string{"type": "string"},
				},
			},
		},
	},
}

func rootResult(s string) string {
	return `root-0-name ::= "\"create_event\""
freestring ::= (
		[^"\\] |
		"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
  )* space
root-0 ::= "{" space "\"arguments\"" space ":" space root-0-arguments "," space "\"name\"" space ":" space root-0-name "}" space
root-1-arguments ::= "{" space "\"query\"" space ":" space string "}" space
realvalue ::= root-0 | root-1
root ::= ` + s + `
space ::= " "?
root-0-arguments ::= "{" space "\"date\"" space ":" space string "," space "\"time\"" space ":" space string "," space "\"title\"" space ":" space string "}" space
root-1 ::= "{" space "\"arguments\"" space ":" space root-1-arguments "," space "\"name\"" space ":" space root-1-name "}" space
string ::= "\"" (
[^"\\] |
"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
)* "\"" space
arr  ::=
"[\n"  (
	realvalue
(",\n"  realvalue)*
)? "]"
root-1-name ::= "\"search\""`
}

const (
	testInput1 = `
	{
		"oneOf": [
			{
				"type": "object",
				"properties": {
					"function": {"const": "create_event"},
					"arguments": {
						"type": "object",
						"properties": {
							"title": {"type": "string"},
							"date": {"type": "string"},
							"time": {"type": "string"}
						}
					}
				}
			},
			{
				"type": "object",
				"properties": {
					"function": {"const": "search"},
					"arguments": {
						"type": "object",
						"properties": {
							"query": {"type": "string"}
						}
					}
				}
			}
		]
	}`

	inputResult1 = `root-0-function ::= "\"create_event\""
freestring ::= (
		[^"\\] |
		"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
  )* space
root-0 ::= "{" space "\"arguments\"" space ":" space root-0-arguments "," space "\"function\"" space ":" space root-0-function "}" space
root-1-arguments ::= "{" space "\"query\"" space ":" space string "}" space
root ::= root-0 | root-1
space ::= " "?
root-0-arguments ::= "{" space "\"date\"" space ":" space string "," space "\"time\"" space ":" space string "," space "\"title\"" space ":" space string "}" space
root-1 ::= "{" space "\"arguments\"" space ":" space root-1-arguments "," space "\"function\"" space ":" space root-1-function "}" space
string ::= "\"" (
	[^"\\] |
	"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
)* "\"" space
root-1-function ::= "\"search\""`

	inputResult2 = `root-0-function ::= "\"create_event\""
freestring ::= (
		[^"\\] |
		"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
  )* space
root-0 ::= "{" space "\"arguments\"" space ":" space root-0-arguments "," space "\"function\"" space ":" space root-0-function "}" space
root-1-arguments ::= "{" space "\"query\"" space ":" space string "}" space
realvalue ::= root-0 | root-1
root ::= arr | realvalue
space ::= " "?
root-0-arguments ::= "{" space "\"date\"" space ":" space string "," space "\"time\"" space ":" space string "," space "\"title\"" space ":" space string "}" space
root-1 ::= "{" space "\"arguments\"" space ":" space root-1-arguments "," space "\"function\"" space ":" space root-1-function "}" space
string ::= "\"" (
	[^"\\] |
	"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
)* "\"" space
arr  ::=
  "[\n"  (
		realvalue
    (",\n"  realvalue)*
  )? "]"
root-1-function ::= "\"search\""`

	testInput2 = `
{
	"oneOf": [
		{
			"type": "object",
			"properties": {
				"name": {"const": "create_event"},
				"arguments": {
					"type": "object",
					"properties": {
						"title": {"type": "string"},
						"date": {"type": "string"},
						"time": {"type": "string"}
					}
				}
			}
		},
		{
			"type": "object",
			"properties": {
				"name": {"const": "search"},
				"arguments": {
					"type": "object",
					"properties": {
						"query": {"type": "string"}
					}
				}
			}
		}
	]
}`

	inputResult3 = `root-0-name ::= "\"create_event\""
freestring ::= (
		[^"\\] |
		"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
  )* space
root-0 ::= "{" space "\"arguments\"" space ":" space root-0-arguments "," space "\"name\"" space ":" space root-0-name "}" space
root-1-arguments ::= "{" space "\"query\"" space ":" space string "}" space
root ::= root-0 | root-1
space ::= " "?
root-0-arguments ::= "{" space "\"date\"" space ":" space string "," space "\"time\"" space ":" space string "," space "\"title\"" space ":" space string "}" space
root-1 ::= "{" space "\"arguments\"" space ":" space root-1-arguments "," space "\"name\"" space ":" space root-1-name "}" space
string ::= "\"" (
[^"\\] |
"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
)* "\"" space
root-1-name ::= "\"search\""`

	inputResult4 = `root-0-name ::= "\"create_event\""
freestring ::= (
		[^"\\] |
		"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
  )* space
root-0 ::= "{" space "\"arguments\"" space ":" space root-0-arguments "," space "\"name\"" space ":" space root-0-name "}" space
root-1-arguments ::= "{" space "\"query\"" space ":" space string "}" space
realvalue ::= root-0 | root-1
root ::= arr | realvalue
space ::= " "?
root-0-arguments ::= "{" space "\"date\"" space ":" space string "," space "\"time\"" space ":" space string "," space "\"title\"" space ":" space string "}" space
root-1 ::= "{" space "\"arguments\"" space ":" space root-1-arguments "," space "\"name\"" space ":" space root-1-name "}" space
string ::= "\"" (
[^"\\] |
"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
)* "\"" space
arr  ::=
"[\n"  (
	realvalue
(",\n"  realvalue)*
)? "]"
root-1-name ::= "\"search\""`
)

var _ = Describe("JSON schema grammar tests", func() {
	Context("JSON", func() {
		It("generates a valid grammar from JSON schema", func() {
			grammar := NewJSONSchemaConverter("").GrammarFromBytes("", []byte(testInput1), false, false)
			results := strings.Split(inputResult1, "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))))
		})
		It("generates a valid grammar from JSON schema", func() {
			grammar := NewJSONSchemaConverter("").GrammarFromBytes("", []byte(testInput2), false, false)
			results := strings.Split(inputResult3, "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))))
		})
		It("generates a valid grammar from JSON Objects", func() {

			structuredGrammar := JSONFunctionStructureFunction{
				OneOf: testFunctions}

			grammar := structuredGrammar.Grammar("", "", false, false)
			results := strings.Split(inputResult1, "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))))
		})

		It("generates a valid grammar from JSON Objects for multiple function return", func() {
			structuredGrammar := JSONFunctionStructureFunction{
				OneOf: testFunctions}

			grammar := structuredGrammar.Grammar("", "", true, false)
			results := strings.Split(inputResult2, "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})

		It("generates a valid grammar from JSON Objects for multiple function return", func() {
			structuredGrammar := JSONFunctionStructureName{
				OneOf: testFunctionsName}

			grammar := structuredGrammar.Grammar("", "", true, false)
			results := strings.Split(inputResult4, "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})

		It("generates a valid grammar from JSON Objects for multiple function return with a suffix and array", func() {
			structuredGrammar := JSONFunctionStructureName{
				OneOf: testFunctionsName}

			grammar := structuredGrammar.Grammar("suffix", "", true, false)
			results := strings.Split(rootResult(`"suffix" arr | realvalue`), "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})
		It("generates a valid grammar from JSON Objects with a suffix", func() {
			structuredGrammar := JSONFunctionStructureName{
				OneOf: testFunctionsName}

			grammar := structuredGrammar.Grammar("suffix", "", false, false)
			results := strings.Split(rootResult(`"suffix" realvalue`), "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})
		It("generates a valid grammar from JSON Objects with a suffix and could return string", func() {
			structuredGrammar := JSONFunctionStructureName{
				OneOf: testFunctionsName}

			grammar := structuredGrammar.Grammar("suffix", "", false, true)
			results := strings.Split(rootResult(`( "suffix" realvalue | freestring )`), "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})
		It("generates a valid grammar from JSON Objects with a suffix that could return text or an array of tools", func() {
			structuredGrammar := JSONFunctionStructureName{
				OneOf: testFunctionsName}

			grammar := structuredGrammar.Grammar("suffix", "", true, true)
			results := strings.Split(rootResult(`( "suffix" (arr | realvalue) | freestring )`), "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})

		It("generates a valid grammar from JSON Objects without a suffix that could return text or an array of tools or just string", func() {
			structuredGrammar := JSONFunctionStructureName{
				OneOf: testFunctionsName}

			grammar := structuredGrammar.Grammar("", "", true, true)
			results := strings.Split(rootResult(`freestring | arr | realvalue`), "\n")
			for _, r := range results {
				if r != "" {
					Expect(grammar).To(ContainSubstring(r))
				}
			}
			Expect(len(results)).To(Equal(len(strings.Split(grammar, "\n"))), grammar)
		})
	})
})
