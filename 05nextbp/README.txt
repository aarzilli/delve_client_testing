BREAKPOINTS HIT DURING NEXT SHOULD BE DEALT WITH CORRECTLY

Instructions:
- debug main.go
- set a breakpoint on main.go:23
- continue once
- set a breakpoint on main.go:12
- repeatedly hit next
- one of two things should happen:
	- the program stops at a different breakpoint and the 'next' command is cancelled with CancelNext
	- the program skips some breakpoints and completes the 'next' command
- either way it should be possible to see the program terminate while just hitting 'next' repeatedly, without errors
