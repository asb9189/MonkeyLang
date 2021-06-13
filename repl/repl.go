package repl

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "MonkeyLang/lexer"
  "MonkeyLang/parser"
  "MonkeyLang/evaluator"
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Fprintf(out, PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    if line == "exit" { os.Exit(0) }

    l := lexer.New(line)
    p := parser.New(l)

    program := p.ParseProgram()

    if len(p.Errors()) != 0 {
      printParserErrors(out, p.Errors())
      continue
    }

    evaluated := evaluator.Eval(program)
    if evaluated != nil {
      io.WriteString(out, evaluated.Inspect())
      io.WriteString(out, "\n")
    }

  }
}

func printParserErrors(out io.Writer, errors []string) {
  io.WriteString(out, MONKEY_FACE)
  io.WriteString(out, "🙊🙈 Looks like somebody made an oopsie 😳\n")
  io.WriteString(out, "parser errors:\n")
  for _, msg := range errors {
    io.WriteString(out, "\t" + msg + "\n")
  }
}
