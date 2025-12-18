package program_data

type Structure map[string]*Namespace

// TODO: Avoid global state. Pass a context or symbol table object through the compilation phases.
var namespaces = make(Structure)
