package main

import (
	"context"
	"fmt"

	"github.com/zombor/terraform-operator/pkg/terraform"
	"go.temporal.io/sdk/client"
)

func main() {
	var (
		result string
	)

	c, err := client.Dial(client.Options{})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: "TASK_QUEUE",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, terraform.PlanWorkflow, []string{`provider "aws" {}`, `
resource "aws_s3_bucket" "a" {
	bucket = "test-bucket"
}`})
	if err != nil {
		panic(err)
	}

	err = we.Get(context.Background(), &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
