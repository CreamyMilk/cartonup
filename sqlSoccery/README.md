> We just want to able to complete payments using wallets

### Case 1
1. User pays using the app but opts to pay directly with MPESA
2. We use their userID and the provided PhoneNumber to send the payment request
3. On stk callback just send back a fcm notification(TODO)
4. On payment callback just classify the payment and perfrom settlement


