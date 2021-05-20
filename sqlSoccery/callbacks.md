## Payment Callback
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
  "TransactionType": "Pay Bill",
  "TransID": "OI76BFE31M",
  "TransTime": "20200907092133",
  "BusinessShortCode": "4027891",
  "InvoiceNumber": "",
  "OrgAccountBalance": "30800.00",
  "ThirdPartyTransID": "",
  "MSISDN": "254797678252",
  "FirstName": "JOHn",
  "MiddleName": "DOE",
  "LastName": "",
  "TransAmount": "15000.00",
  "BillRefNumber": "R#GF4A"
}' http://localhost:1337/v2/successcall
```

## Payment Callback
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
  "TransactionType": "Pay Bill",
  "TransID": "OI76BFE31M",
  "TransTime": "20200907092133",
  "TransAmount": "6000.00",
  "BusinessShortCode": "4027891",
  "BillRefNumber": "DF#FF6",
  "InvoiceNumber": "",
  "OrgAccountBalance": "30800.00",
  "ThirdPartyTransID": "",
  "MSISDN": "254797678252",
  "FirstName": "JOHn",
  "MiddleName": "DOE",
  "LastName": ""
}' http://localhost:1337/v2/successcall
```

