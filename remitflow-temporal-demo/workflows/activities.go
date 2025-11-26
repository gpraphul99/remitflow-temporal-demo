package workflows

import "fmt"

func (a *Activities) LockFXRate(amount int64) error {
    fmt.Printf("FX rate locked for ¥%d\n", amount)
    return nil
}

func (a *Activities) KYCCheck(customerID string) error {
    if customerID == "bad-kyc-001" {
        return fmt.Errorf("KYC failed")
    }
    fmt.Println("KYC passed")
    return nil
}

func (a *Activities) ComplianceCheck(name string) error {
    fmt.Println("Compliance check passed")
    return nil
}

func (a *Activities) ExecutePayout(req RemittanceRequest) error {
    fmt.Printf("Paid %d JPY → %s to %s\n", req.AmountJPY, req.ToCurrency, req.RecipientName)
    return nil
}

func (a *Activities) Refund(amount int64) error {
    fmt.Printf("Refunded ¥%d\n", amount)
    return nil
}
