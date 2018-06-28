VARIOUS TEST ABOUT PRINTING VARIABLES

Restrictions:
- Mostly applies to GUI/TUI clients

Instructions:
- debug main.go
- continue

then:

12.
- evaluate `longmap["iteration 2"].value` after a while you should get `3`.

13.
- view the value of longstr
- the debugger should signal that it wasn't completely loaded
- there should be a way to load more of the string

14.
- view the value of longslice
- the debugger should signal that it wasn't completely loaded
- there should be a way to load more values

15.
- view the value of longmap
- the debugger should signal that it wasn't completely loaded
- there should be a way to load more values

16.
- view the value of infptr
- open the ptr field
- keep opening the ptr field until you are bored

17.
- view the value of infiface
- open the iface field
- keep opening the iface field until you are bored

18.
- view the value of intmap
- output shouldn't be confusing
