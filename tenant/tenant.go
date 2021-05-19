package tenant

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/database"
)

type Tenant struct {
	TenantID   int
	DueID      int
	FullName   string
	HouseName  string
	BranchID   string
	DueMonth   int
	DueYear    int
	RentStatus int
	AmountDue  int
}

func GetTenantByHouseNo(houseNumber string) *Tenant {
	t := new(Tenant)
	err := database.DB.QueryRow(`SELECT 
		fair.rid as TenantID,
		f_id as DueID,
		r_name as FullName,
		u.unit_no as HouseName,
		u.branch_id as BranchID,
		month_id as DueMonth,
		xyear as DueYear,
		bill_status as RentStatus,
		round(r_rent_pm) as AmountDssue

		FROM tbl_add_fair fair

		inner join tbl_add_floor fl 
		on fl.fid = fair.floor_no 

		inner join tbl_add_unit u
		on u.uid = fair.unit_no 

		inner join tbl_add_rent ar 
		on ar.rid = fair.rid 

		WHERE u.unit_no = ?
		order by RentStatus ASC,
		DueYear ASC,
		DueMonth ASC
`, houseNumber).Scan(
		&t.TenantID,
		&t.DueID,
		&t.FullName,
		&t.HouseName,
		&t.BranchID,
		&t.DueMonth,
		&t.DueYear,
		&t.RentStatus,
		&t.AmountDue,
	)
	fmt.Printf("%+v", t)
	fmt.Println(err)
	return t
}
func GetTenantByID(tid int) *Tenant {
	return nil
}
func (t *Tenant) ClearPayment() error {
	if t.RentStatus == 0 {
		//Payoff the due for the stated DueID
		_, err := database.DB.Exec(`UPDATE tbl_add_fair
		SET rent_zzzzz=?
		WHERE rid=?`, 1, t.DueID)
		if err != nil {
			return err
		}
	}
	//EarlyPayment
	if t.RentStatus == 1 {
		//Add a successgfully Paid Record
		_, err := database.DB.Exec(`INSERT INTO tbl_add_fail
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
