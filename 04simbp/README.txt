SIMULTANEOUS BREAKPOINTS SHOULD BE ALL NOTIFIED TO THE USER

Restrictions:
- must have a multi-CPU machine with GOMAXPROCS not set or set to something other than 1

Instructions:
- debug main.go
- set a breakpoint on main.go:9
- hit continue repeatedly
- eventually multiple breakpoints will be hit simultaneously, they should all be notified to the user
