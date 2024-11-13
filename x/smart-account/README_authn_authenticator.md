### AuthnVerification Authenticator

The `AuthnVerification` authenticator enables users to authenticate transactions using **WebAuthn passkeys**, providing a secure, passwordless authentication mechanism based on public key cryptography. This authenticator leverages the [WebAuthn standard](https://www.w3.org/TR/webauthn-2/), which is widely supported by modern browsers and operating systems. Users can authenticate using hardware security keys (like YubiKeys) or platform authenticators (such as Touch ID, Windows Hello, or Android biometrics).

#### How It Works

- **Registration**:

  - Users generate a new passkey (credential) using the WebAuthn API.
  - The public key and associated metadata (such as the credential ID and cryptographic algorithm) are stored on-chain as an `AuthnPubKey`.
  - This process is similar to registering a passkey on a web service but stores the necessary verification data on the blockchain.

- **Authentication**:
  - When submitting a transaction, users sign the transaction data using their passkey's private key.
    - The `challenge` for the webauthn assertion is the SHA-256 hash of the transaction data which is then again validated on chain.
  - The `AuthnVerification` authenticator retrieves the stored public key and verifies the signature against the transaction data.
  - If the signature is valid, the transaction is authenticated and can be processed by the blockchain.

#### Key Features

- **Passwordless Security**: Enhances security by eliminating passwords, reducing the risk of phishing and password-related attacks.
- **Public Key Cryptography**: Uses asymmetric cryptography, ensuring that private keys never leave the user's device.
- **Replay Protection**: Relies on the blockchain's sequence number (account nonce) to prevent replay attacks without needing additional challenges.
- **Simplified Verification**: Focuses on signature verification, omitting the need to verify rpId, origin, or authenticator flags in this blockchain context.

#### AuthnPubKey Data Structure

The `AuthnPubKey` is a data structure that holds the necessary information for verifying signatures:

- **key_id** (`string`): The credential ID generated during passkey creation.
- **cose_algorithm** (`int32`): The COSE algorithm identifier (e.g., `-7` for ES256).
- **key** (`bytes`): The public key in DER format.
- **relaying_party_id** (`bytes`): The hashed `rpId`.

```protobuf
message AuthnPubKey {
  string key_id = 1;
  int32 cose_algorithm = 2;
  bytes key = 3;
  bytes relaying_party_id = 4;
}
```

#### References to WebAuthn and Passkeys

- **WebAuthn Standard**: [Web Authentication: An API for accessing Public Key Credentials Level 2](https://www.w3.org/TR/webauthn-2/)
- **Passkeys Explained**: [Passkeys: The Future of Secure Authentication](https://webauthn.guide/)

#### Example Usage

##### **1. Registering a Passkey**

**Client-Side (JavaScript)**

```javascript
import { createSigningClient, ixo } from '@ixo/impactxclient-sdk';
import cbor from 'cbor';
import { Fido2Lib } from 'fido2-lib';

export const fido2 = new Fido2Lib({
	timeout: 60000,
	rpId: process.env.NEXT_PUBLIC_AUTHN_RP_ID!,
	rpName: 'Passkey Smart Accounts',
	challengeSize: 64,
	attestation: 'none',
	cryptoParams: [-7, -257],
	authenticatorUserVerification: 'preferred',
});

async function registerPasskey(wallet) {
	const accountAddress = wallet?.baseAccount?.address;
	if (!accountAddress) {
		throw new Error('No account found');
	}

	// Generate registration options
	const registrationOptions = await fido2.attestationOptions();

	// Set up user information
	const publicKeyCredentialCreationOptions = {
		...registrationOptions,
		user: {
			id: Uint8Array.from(accountAddress, c => c.charCodeAt(0)),
			name: accountAddress,
			displayName: accountAddress,
		},
	};

	// Create a new credential (passkey)
	const credential = await navigator.credentials.create({
		publicKey: publicKeyCredentialCreationOptions,
	});
  const rawPublicKey = new Uint8Array(credential.response?.getPublicKey());

	// Extract the attestation object
	const attestationObject = credential.response.attestationObject;
	const decodedAttestationObject = cbor.decodeAllSync(attestationObject)[0];
	const authData = decodedAttestationObject.authData;
	// Extract data from authData
	const { algorithm, rpIdHash } = extractDataFromAuthData(authData);

  // Prepare expected parameters for attestation result verification
	const expectedAttestationResult: ExpectedAttestationResult = {
		challenge: base64urlEncode(registrationOptions.challenge),
		origin: process.env.NEXT_PUBLIC_AUTHN_ORIGIN!,
		rpId: process.env.NEXT_PUBLIC_AUTHN_RP_ID,
		factor: 'either',
	};
  // Verify attestation response
	await fido2.attestationResult(credential, expectedAttestationResult);

	// Create AuthnPubKey object
	const authnPubKey = ixo.smartaccount.crypto.AuthnPubKey.encode(
		ixo.smartaccount.crypto.AuthnPubKey.fromPartial({
			keyId: credential.id,
			key: rawPublicKey,
			coseAlgorithm: algorithm,
      relayingPartyId: rpIdHash,
		}),
	).finish();

	// Create MsgAddAuthenticator message
	const message = {
		typeUrl: '/ixo.smartaccount.v1beta1.MsgAddAuthenticator',
		value: ixo.smartaccount.v1beta1.MsgAddAuthenticator.fromPartial({
			sender: accountAddress,
			authenticatorType: 'AuthnVerification',
			data: authnPubKey,
		}),
	};

	// Sign and broadcast the transaction
	const client = await createSigningClient(CHAIN_RPC_URL, wallet);
	const result = await client.signAndBroadcast(
		accountAddress,
		[message],
		{
			amount: [{ denom: 'uixo', amount: '2000' }],
			gas: '200000',
		},
		'Register passkey as authenticator',
	);

	return { result, credentialId: credential.id };
}

// Utility function to extract data from authData
export function extractDataFromAuthData(authDataBuffer: any) {
	let authData: DataView<any> | null = null;
	if (authDataBuffer?.constructor?.name === 'Uint8Array') {
		authData = new DataView(authDataBuffer.buffer, authDataBuffer.byteOffset, authDataBuffer.byteLength);
	} else if (authDataBuffer?.constructor?.name === 'Buffer') {
		const uint8Array = new Uint8Array(authDataBuffer);
		authData = new DataView(uint8Array.buffer, uint8Array.byteOffset, uint8Array.byteLength);
	} else if (authDataBuffer?.constructor?.name === 'ArrayBuffer') {
		authData = new DataView(authDataBuffer);
	} else {
		throw new Error('Invalid authDataBuffer type');
	}

	let offset = 0;
	// rpIdHash: 32 bytes
	const rpIdHash = authDataBuffer.slice(offset, offset + 32);
	// Skip rpIdHash (32 bytes)
	offset += 32;
	// Read flags (1 byte)
	const flags = authData.getUint8(offset);
	offset += 1;
	// Skip signCount (4 bytes)
	offset += 4;

	// Check if attestedCredentialData is present
	const attestedCredentialDataPresent = (flags & 0x40) !== 0;
	if (!attestedCredentialDataPresent) {
		throw new Error('Attested credential data not present in authData');
	}

	// AttestedCredentialData
	// - aaguid: 16 bytes
	// - credentialIdLength: 2 bytes
	// - credentialId: variable length
	// - credentialPublicKey: variable length (CBOR encoded)

	// Skip aaguid (16 bytes)
	offset += 16;
	// Get credentialIdLength (2 bytes)
	const credentialIdLength = authData.getUint16(offset);
	offset += 2;
	// Skip credentialId
	offset += credentialIdLength;

	// The rest is the credentialPublicKey (CBOR encoded)
	const publicKeyBytes = authDataBuffer.slice(offset);

	// Decode the public key from COSE_Key format
	const cosePublicKey = cbor.decodeAllSync(publicKeyBytes)[0];

	// Extract key parameters based on COSE_Key format for EC2 keys
	const keyType = cosePublicKey.get(1);
	const algorithm = cosePublicKey.get(3);
	const curve = cosePublicKey.get(-1);
	return { keyType, algorithm, curve, rpIdHash };
}
```

##### **2. Authenticating a Transaction**

**Client-Side (JavaScript)**

```javascript
import { makeSignDoc, makeAuthInfoBytes, makeSignBytes, encodePubkey } from '@cosmjs/proto-signing';
import { cosmos, createSigningClient, ixo } from '@ixo/impactxclient-sdk';
import { fromBase64 } from '@cosmjs/encoding';
import { sha256 } from '@cosmjs/crypto';

export async function signAndBroadcastWithPasskey({ wallet, messages, credentialId }) {
	// Get account number and sequence
	const client = await createSigningClient(CHAIN_RPC_URL, wallet);

	const address = wallet.baseAccount.address;
	const { accountNumber, sequence } = await client.getSequence(address);

	// query chain for correct auth id
	const sigAuthId = Long.fromString('1');

	// Create sign doc
	const txBodyBytes = client.registry.encodeTxBody({
		messages,
		memo: 'memo',
		nonCriticalExtensionOptions: [
			{
				typeUrl: '/ixo.smartaccount.v1beta1.TxExtension',
				value: ixo.smartaccount.v1beta1.TxExtension.encode({
					// add selectedAuthenticators per message
					selectedAuthenticators: [sigAuthId, sigAuthId],
				}).finish(),
			},
		],
	});

	const fee = {
		amount: [{ denom: 'uixo', amount: '2000' }],
		gasLimit: 200000,
		payer: '',
		granter: '',
	};
	const authInfoBytes = makeAuthInfoBytes(
		[
			// MUST HAVE, since sequence is still validated, but public can be
			// empty, since it not used for verification
			{
				pubkey: encodePubkey({
					type: 'tendermint/PubKeySecp256k1',
					value: '',
				}),
				sequence: sequence,
			},
		],
		fee.amount,
		fee.gasLimit,
		fee.granter,
		fee.payer,
	);

	const signDoc = makeSignDoc(txBodyBytes, authInfoBytes, CHAIN_ID, accountNumber);
	const signBytes = makeSignBytes(signDoc);

	// Sign the transaction using the passkey
	const signatureData = await signWithPasskey({ signBytes, credentialId });
	const signatureBytes = fromBase64(Buffer.from(JSON.stringify(signatureData)).toString('base64'));

	// Build the signed transaction
	const txRaw = cosmos.tx.v1beta1.TxRaw.fromPartial({
		bodyBytes: txBodyBytes,
		authInfoBytes: authInfoBytes,
		signatures: [signatureBytes],
	});
	const txRawBytes = cosmos.tx.v1beta1.TxRaw.encode(txRaw).finish();

	// Broadcast the transaction
	const result = await client.broadcastTx(txRawBytes);
	return result;
}

export async function signWithPasskey({ signBytes, credentialId }: Props) {
	// Compute the SHA-256 hash of the signBytes
	const challengeHash = sha256(signBytes); // (32 bytes)

	const assertionOptions = await fido2.assertionOptions();
	const publicKeyCredentialRequestOptions: PublicKeyCredentialRequestOptions = {
		...assertionOptions,
		challenge: challengeHash,
		allowCredentials: [
			{
				type: 'public-key',
				id: base64urlDecode(credentialId),
			},
		],
	};

	// Get assertion
	const assertion: any = await navigator.credentials.get({
		publicKey: publicKeyCredentialRequestOptions,
	});

	// Prepare signature data
	const signatureData = {
		authenticatorData: base64urlEncode(assertion.response.authenticatorData),
		clientDataJSON: base64urlEncode(assertion.response.clientDataJSON),
		signature: base64urlEncode(assertion.response.signature),
	};
	return signatureData;
}
```

#### Considerations

- **Replay Protection**: The blockchain's sequence number ensures that each transaction is unique and prevents replay attacks without additional challenges.
- **Privacy**: Since the origin and rpIdHash are not verified, users can authenticate from any context, which is suitable for a public blockchain environment.
- **Security**: By relying on the security of WebAuthn passkeys and the blockchain's cryptographic guarantees, this authenticator provides strong security without additional complexity.

#### Summary

The `AuthnVerification` authenticator provides a seamless and secure way for users to authenticate transactions using modern passkey technology. It leverages the strengths of WebAuthn while adapting to the blockchain context, focusing on signature verification and utilizing existing blockchain features for security.

By integrating passkeys into the blockchain authentication process, users benefit from enhanced security and convenience, paving the way for a more user-friendly and secure decentralized ecosystem.

---

**Note**: The examples above assume familiarity with the WebAuthn API and blockchain transaction structures. Ensure that you handle all cryptographic operations securely and follow best practices when integrating authentication mechanisms.

---
