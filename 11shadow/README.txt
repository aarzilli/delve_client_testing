SHADOWED VARIABLES SHOULD BE DISTINGUISHABLE

Restrictions:
- GUI/TUI only

Instructions:
- debug main.go
- set a breakpoint on main.go:9
- list the local variables (delve equivalent `locals`)
- there should be two variables named i
- the one with value 42 should be marked, somehow, as the only one visible
