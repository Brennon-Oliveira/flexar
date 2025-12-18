# Improvement Suggestions for Flexar

This document outlines detailed technical solutions for the architectural and logic issues identified in `ANALYSIS.md`.

## 1. Removing Global State

**Problem:** The `program_data` package relies on a global `namespaces` map, making the compiler thread-unsafe and hard to test.

**Solution:** Introduce a `CompilerContext` or `SymbolTable` struct that holds the state for a single compilation run. Pass this context to all visitors and functions.

### Refactoring `program_data`

Instead of global variables, define a struct:

```go
// program_data/context.go

package program_data

type Context struct {
    Namespaces map[string]*Namespace
    Errors     []error // Centralized error handling (see Section 3)
}

func NewContext() *Context {
    return &Context{
        Namespaces: make(map[string]*Namespace),
        Errors:     make([]error, 0),
    }
}

func (c *Context) GetNamespace(name string) *Namespace {
    if ns, ok := c.Namespaces[name]; ok {
        return ns
    }
    ns := &Namespace{
        Name:      name,
        Files:     []string{},
        Functions: make(map[string]*Function),
        Classes:   make(map[string]*Class),
    }
    c.Namespaces[name] = ns
    return ns
}
```

### Updating the Visitor

Update `FirstPassVisitor` to accept this context:

```go
// first_semantic_pass/first_pass.go

type FirstPassVisitor struct {
    antlr.ParseTreeVisitor
    Context *program_data.Context // Inject context
}

func Visit(tree antlr.ParseTree, ctx *program_data.Context) {
    visitor := &FirstPassVisitor{Context: ctx}
    visitor.VisitProgram(tree.(*parser.ProgramContext))
}

// Usage in visitor methods:
func (v *FirstPassVisitor) VisitNamespace(ctx *parser.NamespaceContext) *program_data.Namespace {
    // Use v.Context instead of global function
    namespace := v.Context.GetNamespace(ctx.Namespace_name().GetText())
    // ...
    return namespace
}
```

### Updating `main.go`

```go
func main() {
    // ... setup ...
    context := program_data.NewContext() // Create context once

    for _, file := range files {
        utils.SetCurrentFile(file)
        compilation.Compile(file, context) // Pass context
    }
    // ...
}
```

---

## 2. Fixing the Namespace Bug

**Problem:** `Namespace.AddFile` re-initializes the `Functions` and `Classes` maps, wiping out data from previous files in the same namespace.

**Solution:** Check if the maps are nil before initializing them.

```go
// program_data/namespace.go

func (n *Namespace) AddFile(file string) {
    n.Files = append(n.Files, file)

    // Only initialize if they don't exist yet
    if n.Functions == nil {
        n.Functions = make(map[string]*Function)
    }
    if n.Classes == nil {
        n.Classes = make(map[string]*Class)
    }
}
```

*Note: If you use the `NewContext` approach from Section 1, you can ensure these are initialized at creation time, making this check redundant but still safe.*

---

## 3. Proper Error Handling

**Problem:** `utils.SemanticError` calls `os.Exit(0)`, which is bad practice for libraries and tools.

**Solution:** Accumulate errors in the `Context` and check them after passes.

### Update Context

```go
// program_data/context.go

func (c *Context) AddError(line int, message string) {
    err := fmt.Errorf("(%s) SemanticError on line %d: %s", utils.GetCurrentFile(), line, message)
    c.Errors = append(c.Errors, err)
}

func (c *Context) HasErrors() bool {
    return len(c.Errors) > 0
}
```

### Update Usage

```go
// first_semantic_pass/first_pass.go

func (v *FirstPassVisitor) VisitClass(ctx *parser.ClassContext, namespace *program_data.Namespace) interface{} {
    className := ctx.NAME().GetText()

    if namespace.ExistsClass(className) {
        // Don't exit, just record error
        v.Context.AddError(ctx.NAME().GetSymbol().GetLine(), "Class "+className+" already exists...")
        return nil // Stop processing this specific node if needed
    }
    // ...
}
```

### Check in `main.go`

```go
    // ... after compilation loop ...
    if context.HasErrors() {
        for _, err := range context.Errors {
            fmt.Println(err)
        }
        os.Exit(1)
    }
```

---

## 4. Grammar Improvements

### Fixing Negative Numbers

**Problem:** `INT_NUM : [-]? [0-9]+` causes ambiguity with subtraction (e.g., `1-1`).

**Solution:** Handle unary minus in the parser, not the lexer.

**Lexer Change:**
```antlr
// Flexar.g4

INT_NUM
    : [0-9]+  // Remove [-]?
    ;
```

**Parser Change:**
Add a unary minus rule to your math expressions.

```antlr
unary_math
    : before_unary 
    | after_unary 
    | math_value 
    | parenthesis_expression
    | MINUS unary_math  // Add this line for recursive unary minus (e.g. -5, --5)
    ;
```

### Unifying Expressions

**Problem:** Separating `expression` and `expression_math` prevents natural composition (e.g., `(a > b) + 5`).

**Solution:** Use a single `expression` rule with precedence defined by the order of alternatives.

```antlr
expression
    : OPEN_PAREN expression CLOSE_PAREN      # ParenExpr
    | NOT expression                         # NotExpr
    | expression EXP expression              # PowerExpr
    | expression (STAR|DIV|MODULE) expression # MultExpr
    | expression (PLUS|MINUS) expression     # AddExpr
    | expression (SHL|SHR) expression        # ShiftExpr
    | expression (LESS|GREATER|...) expression # RelationalExpr
    | expression (EQUAL|NOT_EQUAL) expression  # EqualityExpr
    | expression BIT_AND expression          # BitAndExpr
    | expression BIT_XOR expression          # BitXorExpr
    | expression BIT_OR expression           # BitOrExpr
    | expression AND expression              # LogicAndExpr
    | expression OR expression               # LogicOrExpr
    | <assoc=right> expression ASSIGN expression # AssignExpr
    | atom                                   # AtomExpr
    ;

atom
    : INT_NUM
    | FLOAT_NUM
    | STRING
    | ID
    | func_call
    ;
```

*Note: You will need to label the alternatives (e.g., `# AddExpr`) to make visitor implementation easier.*

---

## 5. CLI Improvements

**Problem:** Manual `os.Args` parsing is brittle.

**Solution:** Use a library like `github.com/spf13/cobra` or the standard `flag` package.

**Example using `flag` (Standard Lib):**

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define flags
    outputDir := flag.String("o", "./build", "Output directory")
    verbose := flag.Bool("v", false, "Verbose output")
    
    flag.Parse()

    // Get positional arguments (files/dirs)
    args := flag.Args()
    if len(args) < 1 {
        fmt.Println("Usage: flexar [options] <file.fl|directory>")
        flag.PrintDefaults()
        os.Exit(1)
    }

    projectPath := args[0]
    // ... proceed with projectPath ...
}
```
