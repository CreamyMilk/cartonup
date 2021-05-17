package tenant

type Tenant struct {
	TenantID  int
	FullName  string
	HouseName string
	DueMonth  int
	DueYear   int
	AmountDue int64
}

func GetTenantByID(tid int) *Tenant {
	return nil
}
func (t *Tenant) GetDetails() error {
	return nil
}

func (t *Tenant) GetLatestDue() error {
	return nil
}

func (t *Tenant) GetWalletBalance() error {
	return nil
}

func (t *Tenant) PayRentViaWallet() error {
	return nil
}
