#goScript

Golang like scripting language

Uses go/ast and reflect to allow arbitrary expression evaluation
in a given context.

#Highlights

* Simple and hackable code.

* Fully reflect function call including variable number of args.

* Automatic casting
  * Numeric casting  always to the bigger more generic representation.
  * True && 0 = false
  * Left operand governs the type 1 + "1" = 2,  "1" + 1 = "11"

* Map, slice and array indexer access

* Full control over evaluation context.
  *Generic map[string]interface{} or map[string]interface{} contexts
  *Any struct with its fields as variables
  *Custom Context for implentation/user defined context.

* Mostly tested ;)

#Roadmap

* 1) Rock solid expression evaluation.
* 2) Enrich evaluation context
* 3) Full script program implementation
* 4) Performance optimizations
