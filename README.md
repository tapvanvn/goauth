# GoAuth 
General convenient authentication library.

**```In Construction```**

## Client
Client adapt to system to provide an account manager machanism.

## Client may
- Issue an unique id to each account.
- Provides a method to authorize the incoming user in login process.
- Also may or may not provice a method for issueing and verifying the signature of account.
- May be self managed type such as userpass or email/phone, or thirdparty such as google/apple/ethereum.

### Session begining
- If client is self managed such as usepass or email, It can check if a process of register or verifing account is needed, or just release an error if account is not registed yet.

- If everything is ok, issue a SessionAdapt which contains the sessionID, accountID, clientType and also client's specific infomation if needed.
