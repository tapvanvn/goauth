# GoAuth 
General convenient authentication library.

**```In Construction```**

## Client
Client adapt to system to provide an account manager machanism.

## Client may
- Issue an unique id to each account.
- Provide a method to authorize the incoming user in login process.
- Also may or may not provice a method for issueing and verifying the signature of account.
- May be self managed type such as userpass or email/phone, or thirdparty such as google/apple/ethereum.

### Session begining
- If client is self managed such as usepass or email, check if a signing up/verifing account process is needed, or just release an error if account is not registed yet.

- If everything is ok, issue a SessionAdapt which contains the sessionID, accountID, clientType and also client's specific infomation that it need to process.
