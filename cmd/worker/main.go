package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/zombor/terraform-operator/pkg/os"
	"github.com/zombor/terraform-operator/pkg/terraform"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}

	defer c.Close()

	w := worker.New(c, "TASK_QUEUE", worker.Options{})
	w.RegisterWorkflow(terraform.PlanWorkflow)
	w.RegisterWorkflow(terraform.ApplyWorkflow)
	w.RegisterActivity(terraform.Plan)
	w.RegisterActivity(terraform.Apply)
	w.RegisterActivity(os.Write)
	w.RegisterActivity(os.Delete)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		panic(err)
	}
}
