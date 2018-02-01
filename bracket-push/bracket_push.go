package brackets

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	size := len(s)
	return s[:size-1], s[size-1]
}

// func (s stack) Top() rune {
// 	return s[len(s)-1]
// }

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

// Bracket evaluates given expression and if matches returns true.
func Bracket(expr string) (bool, error) {
	s := make(stack, 0)

	for _, cur := range expr {
		if cur == '{' || cur == '[' || cur == '(' {
			s = s.Push(cur)
		} else if cur == '}' || cur == ']' || cur == ')' {
			if !s.IsEmpty() {
				var top rune
				s, top = s.Pop()
				if !isMatch(cur, top) {
					return false, nil
				}
			} else {
				return false, nil
			}
		} else {
			continue
		}
	}

	return s.IsEmpty(), nil
}

func isMatch(cur rune, top rune) bool {
	switch cur {
	case '}':
		return top == '{'
	case ')':
		return top == '('
	case ']':
		return top == '['
	default:
		return false
	}
}

// go test -run ^NOTHING -bench 'BenchmarkBracket'
// goos: darwin
// goarch: amd64
// BenchmarkBracket-8   	 2000000	       972 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/bracket-push	2.941s
