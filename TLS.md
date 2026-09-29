https://github.com/mikepound/tls-exercises
[C:\Users\KANDERS3\Downloads\experttransportlayersecuritytls1786989294089.pdf]
https://github.com/mikepound/tls-exercises/blob/master/glossary.pdf


SSL is deprecated, use TLS -> TLS1.3

phase:
- handshake
- data

Record Protocol:
- Header - clear, unencrypted
	- Type
		- subprotocol of (20,21,22,23)
			- 20 ChangeCipherSpec
				- switching to a new encryption suite
			- 21 Alert
				- Alert mechanism
				- has Level and description
					- level:
						- 1- warning
						- 2 - fatal
					- Description: - only broad error type, no actual description
						- e.g. close_notify(0), bad_certificate(42)
				- Structure:
					- record - nonce - c - tag
			- 22 Handshake
			- 23 Application Data
				- Opaque application data - encrypted
	- Version
		- major and minor: 03 01 for TLS 1.0
	- Length
		- max 16kb
- Data
	- 