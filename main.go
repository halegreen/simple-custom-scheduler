package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	simple "github.com/halegreen/simple-custom-scheduler/pkg"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	command := app.NewSchedulerCommand(app.WithPlugin(simple.Name, simple.NewSimplePlugin))
	logs.InitLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
