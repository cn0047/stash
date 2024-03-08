Payment
-

PAN - Primary Account Number (payment card number).

Pre-Authorization - to verify that credit card is valid and has sufficient funds.
Puts hold on funds for 7-10 days, after which the funds are released back to cardholder,
if they are not captured in that timeframe.

Capture - can be taken on its own or against a Pre-Authorization.
Can only be initiated for up to the Pre-Authorized amount or less.

# PCI DSS (Payment Card Industry Data Security Standard)

Can be stored, but must be adequately protected:
* Card holder name.
* Masked or hashed PAN.
* Service code and expiration data.
