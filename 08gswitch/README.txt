GOROUTINE SWITCHING

- debug main.go
- continue
- once the program pauses, view the goroutine listing (delve equivalent: `goroutine`, there should be ~100), switch to a different goroutine (delve equivalent `goroutine <n>`)
- view local variables
- hit next
- execution should stay in the same goroutine
