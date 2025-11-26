package workflows

import (
    "time"
    "go.temporal.io/sdk/workflow"
)

type RemittanceRequest struct {
    WorkflowID       string
    CustomerID       string
    AmountJPY        int64
    ToCurrency       string
    RecipientName    string
    RecipientAccount string
}

type Activities struct{}

func RemittanceWorkflow(ctx workflow.Context, req RemittanceRequest) (string, error) {
    ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Second}
    ctx = workflow.WithActivityOptions(ctx, ao)

    var a *Activities

    // 1. Lock FX rate
    err := workflow.ExecuteActivity(ctx, a.LockFXRate, req.AmountJPY).Get(ctx, nil)
    if err != nil { return "", err }

    // 2. KYC check
    err = workflow.ExecuteActivity(ctx, a.KYCCheck, req.CustomerID).Get(ctx, nil)
    if err != nil {
        workflow.ExecuteActivity(ctx, a.Refund, req.AmountJPY).Get(ctx, nil)
        return "failed: kyc", err
    }

    // 3. Compliance
    err = workflow.ExecuteActivity(ctx, a.ComplianceCheck, req.RecipientName).Get(ctx, nil)
    if err != nil {
        workflow.ExecuteActivity(ctx, a.Refund, req.AmountJPY).Get(ctx, nil)
        return "failed: compliance", err
    }

    // 4. Payout
    err = workflow.ExecuteActivity(ctx, a.ExecutePayout, req).Get(ctx, nil)
    if err != nil {
        workflow.ExecuteActivity(ctx, a.Refund, req.AmountJPY).Get(ctx, nil)
        return "failed: payout", err
   )

    return "success", nil
}