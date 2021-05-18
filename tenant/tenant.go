package tenant

import "github.com/CreamyMilk/cartonup/database"

type Tenant struct {
	TenantID   int
	DueID      int
	FullName   string
	HouseName  string
	DueMonth   int
	DueYear    int
	RentStatus int
	AmountDue  int64
}

func GetTenantByHouseNo(houseNumber string) *Tenant {
	return nil

}
func GetTenantByID(tid int) *Tenant {
	return nil
}
func (t *Tenant) ClearPayment() error {
	if t.RentStatus == 0 {
		//Payoff the due for the stated DueID
		err := databse.DB.Exec(`UPDATE tbl_add_fair
		SET rent_status=?
		WHERE rid=?`, 1, c.DueID)
		if err != nil {
			return err
		}
	}
	//EarlyPayment
	if t.RentStatus == 1 {
		//Add a successgfully Paid Record
		err := database.DB.Exec(`INSERT INTO tbl_add_fail
		userId,y.............
		`)
		if err != nil {
			return err
		}
	}
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
