package scheduler

import "fmt"

type PrintJob struct {
	Message string
}

func (p *PrintJob) Execute() {
	fmt.Println(p.Message)
}
