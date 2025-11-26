package main

import (
    "log"
    "remitflow-temporal-demo/workflows"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil { log.Fatalln(err) }
    defer c.Close()

    w := worker.New(c, "remitflow-task-queue", worker.Options{})
    w.RegisterWorkflow(workflows.RemittanceWorkflow)
    w.RegisterActivity(workflows.Activities)

    log.Println("Worker started")
    err = w.Run(worker.InterruptCh())
    if err != nil { log.Fatalln(err) }
}