package main

import (
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"math/rand"
)

func main() {
	rand.Seed()
	command := app.NewSchedulerCommand(app.WithPlugin())
}
