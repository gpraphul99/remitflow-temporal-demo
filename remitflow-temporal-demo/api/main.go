package main

import (
    "encoding/json"
    "log"
    "net/http"
    "remitflow-temporal-demo/workflows"
    "github.com/google/uuid"
    "github.com/gorilla/chi/v5"
    "go.temporal.io/sdk/client"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil { log.Fatalln(err) }
    defer c.Close()

    r := chi.NewRouter()
    r.Post("/remit", func(w http.ResponseWriter, r *http.Request) {
        var req workflows.RemittanceRequest
        json.NewDecoder(r.Body).Decode(&req)
        req.WorkflowID = "remit-" + uuid.New().String()

        opts := client.StartWorkflowOptions{
            ID:        req.WorkflowID,
            TaskQueue: "remitflow-task-queue",
        }
        _, err := c.ExecuteWorkflow(r.Context(), opts, workflows.RemittanceWorkflow, req)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        json.NewEncoder(w).Encode(map[string]string{"status": "started", "id": req.WorkflowID})
    })

    log.Println("API listening on :8080")
    http.ListenAndServe(":8080", r)
}