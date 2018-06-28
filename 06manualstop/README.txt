IT SHOULD BE POSSIBLE TO PAUSE A PROCESS MANUALLY

- debug main.go
- continue
- after a while hit the pause command (delve equivalent: press Ctrl-C)
- check the goroutine list, goroutine 1 should be calling time.Sleep (or printing, depending on your timing)
