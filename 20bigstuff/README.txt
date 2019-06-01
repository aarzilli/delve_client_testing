Instructions:
- debug main.go
- when the execution stop, check the value of 's' it should contain one hundred million NUL characters, the debugger should show the first few ones
- use step over 4 times, the debugger should not freeze or be exceedingly slow

Possible causes of slowness include the aforementioned large variable 's', the fact that the stack is 10 million frames deep, or that there's one million goroutines sleeping.

