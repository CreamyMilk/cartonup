package tenant

import (
	"time"

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

func (t *Tenant) ClearPayment() error {
	if t.RentStatus == 0 {
		//Payoff the due for the stated DueID

		_, err := database.DB.Exec(`UPDATE
		tbl_add_fair
		SET bill_status=?,
		paid_date=? 
		WHERE f_id=?`, 1, time.Now().Format("02/01/2006"), t.DueID)
		if err != nil {
			return err
		}
	}

	//Early
	//TODOl
	if t.RentStatus == 1 {

		//Add a successfully Paid Record
		_, err := database.DB.Exec(`INSERT INTO icrib_house_db.tbl_add_fair (
     type , floor_no , unit_no , rid , month_id , xyear , rent , water_bill , electric_bill
    , gas_bill , security_bill , utility_bill , other_bill , total_rent , issue_date
    , issued_date , paid_date , branch_id , bill_status , added_date
) VALUES (
	"RENTED",
    , ? -- floor_no varchar
    , ? -- unit_no varchar
    , ? -- rid int
    , ? -- month_id int
    , ? -- xyear int NULLABLE
    , ? -- rent decimal
    , ? -- water_bill decimal
    , ? -- electric_bill decimal
    , ? -- gas_bill decimal
    , ? -- security_bill decimal
    , ? -- utility_bill decimal
    , ? -- other_bill decimal
    , ? -- total_rent decimal
    , ? -- issue_date varchar
    , ? -- issued_date date NULLABLE
    , ? -- paid_date varchar NULLABLE
    , ? -- branch_id int
    , ? -- bill_status tinyint
    , ? -- added_date timestamp
)
		`, "RENTED")
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
