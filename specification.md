

## Cryptographic Package Standard
The RSA PSS package being used implements the following standard: RSASSA-PSS signature scheme according to RFC 8017

## Package explanation
It is a Golang service api that stores a list of whitelisted domain urls each with their own signed RSA PSS signature. The RSA signature shares a public and private key between two parties, where the server holds the private key and the client holds the public key. In this case its Opera and IXO. The key of the mechanism is Opera can validate that the response from the whitelist server is cryptographically correct by comparing the RSA PSS signature provided with their own signature generated using the urls in the response. They must match or there is likely a man in the middle attack occuring between the opera client and the ixo server. The server address is to be confirmed once a set of RSA keys is agreed upon.

## Package Requirements
It defines a client server relationship where IXO (The server) shares a public RSA key with client and the client uses the key to verify the integrity of the messages being sent. The requirement on the Client side is to have a RSA PSS verify function in line with the RSA standard defined above for comparison purposes.

## API Example Response Object
```json
   {
        data: {

            ID: 1,
            CreatedAt: "",
            UpdatedAt: "",
            DeletedAt: null,  // This is for soft deletes in terms of domains in the process of leaving the whitelist
            name: "examplename",
            url: "exampleurl",
            signature: "examplepksssignature"

        }
        message: "success"
        success: true

    }
```

## API Relevent Endpoint
### Gets all the domain whitelist items
```

/api/getwhitelist
```
