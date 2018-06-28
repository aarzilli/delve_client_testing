CORE FILES SHOULD BE DEBUGGABLE

Restrictions:
- only applies to Linux debuggers

Instructions:
- run makecore.sh
- open the core file with the debugger (delve equivalent: `dlv core panic panic-core`)
- switch to goroutine 1 (delve equivalent: `goroutine 1`)
- print the stack (delve equivalent: `stack`)
- the stack trace should contain `main.main`
