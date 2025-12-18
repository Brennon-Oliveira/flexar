# Flexar Project Analysis

## 1. File Structure and Project Architecture

The project follows a flat structure with packages separated by functionality (`compilation`, `grammar`, `program_data`, `first_semantic_pass`).

*   **`grammar/`**: Contains the ANTLR4 grammar and generated Go code. This is standard.
*   **`program_data/`**: Acts as the symbol table / AST storage. It relies heavily on global state.
*   **`compilation/`**: The driver for the compilation process.
*   **`first_semantic_pass/`**: Implements the first pass of semantic analysis (symbol collection).

**Critique:**
The structure is simple but lacks separation of concerns regarding state management. The `program_data` package effectively acts as a global singleton, which couples the entire compiler to a single instance of execution.

## 2. Code Patterns and Technical Decisions

### Global State
The most critical issue is the use of global variables in `program_data/program_data.go`:
```go
var namespaces = make(Structure)
```
This makes the compiler:
*   **Not Thread-Safe**: You cannot compile multiple projects or files in parallel.
*   **Hard to Test**: Tests that modify this state will interfere with each other unless carefully reset.
*   **Rigid**: It's hard to pass context (like configuration or cancellation) through the pipeline.

### Error Handling
The project uses `os.Exit(0)` deep within the call stack (e.g., `utils.SemanticError`, `compilation.Compile`).
*   **Problem**: This prevents the compiler from reporting multiple errors or being used as a library (e.g., by a language server).
*   **Recommendation**: Functions should return `error` or a list of errors. The `main` function should decide when to exit.

### Visitor Pattern
The `FirstPassVisitor` in `first_semantic_pass` manually dispatches calls (e.g., `v.VisitNamespace(...)`) instead of using the `Accept` method or the generated `BaseVisitor` interface fully. While functional, it's less idiomatic for ANTLR.

### Namespace Logic (Critical Bug)
In `program_data/namespace.go`, the `AddFile` method resets the `Functions` and `Classes` maps every time a file is added:
```go
func (n *Namespace) AddFile(file string) {
    n.Files = append(n.Files, file)
    n.Functions = make(map[string]*Function) // Wipes existing data!
    n.Classes = make(map[string]*Class)      // Wipes existing data!
}
```
This means a namespace spanning multiple files will only retain symbols from the last file processed.

## 3. Grammar and Language Design

### Lexer Issues
The `INT_NUM` rule includes the negative sign:
```antlr
INT_NUM : [-]? [0-9]+ ;
```
This causes ambiguity with subtraction. `1-1` (without spaces) is tokenized as `INT_NUM(1)` followed by `INT_NUM(-1)`, causing a syntax error because the parser expects an operator. Negative numbers should be handled in the parser via a unary minus rule.

### Expression Rules
The grammar splits `expression` and `expression_math`.
*   **Problem**: This separation limits composability. For example, you cannot easily use a comparison result inside a math expression (e.g., `(a == b) + 1`) because `expression_math` does not recurse back to the full `expression` rule (except via specific paths that might be blocked).
*   **Recommendation**: Unify into a single `expression` rule with proper precedence levels.

## 4. Recommendations

1.  **Remove Global State**: Create a `CompilerContext` or `SymbolTable` struct and pass it to visitors and functions.
2.  **Fix `AddFile` Bug**: Only initialize maps if they are nil.
3.  **Refactor Error Handling**: Accumulate errors in a list and return them. Remove `os.Exit` from library code.
4.  **Fix Grammar**:
    *   Remove `[-]?` from `INT_NUM`. Add `unary_math : MINUS term_math` (or similar).
    *   Unify `expression` and `expression_math` to allow full recursion and correct precedence.
5.  **CLI Improvements**: Use a library like `cobra` or `urfave/cli` for robust argument parsing.

## 5. Summary
The project is a good start but suffers from fundamental architectural issues (global state) and critical logic bugs (namespace wiping, lexer ambiguity) that need to be addressed before it can be reliable or scalable.
