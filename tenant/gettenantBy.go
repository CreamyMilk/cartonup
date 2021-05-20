package tenant

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/database"
)

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
	if err != nil {
		fmt.Println(err)
	}
	return t
}
func GetTenantByID(tid int) *Tenant {
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

		WHERE fair.rid = ?
		order by RentStatus ASC,
		DueYear ASC,
		DueMonth ASC
`, tid).Scan(
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
	if err != nil {
		fmt.Println(err)
	}
	return t
}
