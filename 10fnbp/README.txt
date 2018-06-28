SET BREAKPOINTS ON FUNCTIONS

- debug main.go
- set a breakpoint on main.main (if the debugger has a command to set breakpoints on function names use that)
- breakpoint should be set after the prologue
	- on linux + go1.10.3 this is address 0x49b358, on other architectures/compiler versions use the disassembler to check

