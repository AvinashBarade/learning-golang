1. Memory managment in GO?
2. Concurancy in go 
3. What do you understand by Shadowing in Go?
==> n Go, shadowing refers to the practice where a variable declared within a certain scope (typically in a nested block) has the same name as a variable declared in an outer scope. When this happens, the inner variable "shadows" the outer variable, meaning that within the inner scope, the outer variable is temporarily hidden and inaccessible. Instead, the inner variable is used.
4. what is variadic functions in Go?
==> functions that can take a variable number of arguments of the same type. They provide a flexible way to handle multiple inputs without having to define a specific number of parameters. func sum(numbers ...int) int {}
5. byte and rune data types?
==> The byte represents ASCII characters whereas the rune represents a single Unicode character which is UTF-8 encoded by default.
byte: Used for raw binary data or ASCII characters. It's useful when you need to handle data at the byte level.
rune: Used for Unicode characters, ensuring correct representation of characters from any language or symbol set.
6. 