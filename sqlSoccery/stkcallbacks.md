## Success Callback
```json
curl --header "Content-Type: application/json"   --request POST   --data '{

"_id": "5f479ae1185f270004ddcd01",
"Body": {
"stkCallback": {
"MerchantRequestID": "20296-80517985-1",
"CheckoutRequestID": "ws_CO_010620212348254686",
"ResultDesc": "Request cancelled by user",
"ResultCode": 0}}
}' http://localhost:1337/v2/stkpushcall
```

