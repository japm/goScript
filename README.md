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

#Roadmap
