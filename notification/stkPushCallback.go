package notification

import (
	"fmt"
)

type StkPushCallback struct {
	ID   string `json:"_id"`
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

func (call *StkPushCallback) Classify() error {
	s := GetsktStoreByCheckID(call.Body.StkCallback.CheckoutRequestID)
	//We should do look ups here to notify our users
	fmt.Println(s)
	switch call.Body.StkCallback.ResultCode {
	case SUCESSFUL_PAYMENT:
		fmt.Println("So This Payment was successful")
	case INSUFFIECENT_BALANCE:
	case USER_IS_NONEEXISTANT:
	case USER_REJECTED_PAYMENT:
	case USER_STILL_REJECTED:
	case USER_COULD_NOT_BE_LOCATED:
	default:
		fmt.Println("So every thing failed to process this stuff")
	}
	return nil
}
