To communicate with Sign in with Apple, you’ll use a private key to sign one or more developer tokens.

First enable the Sign in with Apple service on an iOS, tvOS, watchOS, or macOS App ID and classify as the primary App ID. Enable the service on related apps and associate using the grouping feature. Register a Services ID, verify your domain, and associate to an app for each website that uses Sign in with Apple.

Next create and download a private key with Sign in with Apple enabled and associate it with a primary App ID. You can associate two keys with each primary App ID. Then get the key identifier (kid) to create a JSON web token (JWT) that you’ll use to communicate with the capabilities you enabled.

If you suspect a private key is compromised, first create a new private key associated with the primary App ID. Then, after transitioning to the new key, revoke the old private key.

https://help.apple.com/developer-account/#/devcdfbb56a3
https://help.apple.com/developer-account/#/dev3a82eef1c?sub=deve86584d6a