## Request For Payment 
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
  "TransactionType": "Pay Bill",
  "TransID": "OI76BFE31M",
  "TransTime": "20200907092133",
  "TransAmount": "6000.00",
  "BusinessShortCode": "4027891",
  "BillRefNumber": "ALICE FF6",
  "InvoiceNumber": "",
  "OrgAccountBalance": "30800.00",
  "ThirdPartyTransID": "",
  "MSISDN": "254723713074",
  "FirstName": "ALICE",
  "MiddleName": "MATHENGE",
  "LastName": ""
}' http://localhost:1337/v2/successcall
```
