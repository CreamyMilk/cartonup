package notification

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/database"
)

type StkStore struct {
	storeid           int
	HouseName         string
	Branch            string
	PhoneNumber       string
	CheckoutRequestID string
	timeStamp         int64
}

func GetsktStoreByCheckID(check string) *StkStore {
	s := new(StkStore)
	err := database.DB.QueryRow(`SELECT 
	ID,
	houseName,
	branch,
	phoneNo,
	checkoutRequestID,
	UNIX_TIMESTAMP(createdAt) as timestamp FROM stkCalls WHERE checkoutRequestID=?`, check).Scan(&s.storeid, &s.HouseName, &s.Branch, &s.PhoneNumber, &s.CheckoutRequestID, &s.timeStamp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return s
}

func (s *StkStore) Store() error {
	_, err := database.DB.Exec(`
	INSERT INTO stkCalls (houseName,branch,phoneNo,checkoutRequestID) VALUES
	(?,?,?,?)
	`, s.HouseName, s.Branch, s.PhoneNumber, s.CheckoutRequestID)
	if err != nil {
		return err
	}
	return nil
}

// func (s *StkStore) String() string {
// 	return s.CheckoutRequestID
// }
