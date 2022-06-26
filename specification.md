

## Cryptographic Package Standard
The HMAC package being used implements the following standard: Keyed-Hash Message Authentication Code (HMAC) as defined in U.S. Federal Information Processing Standards Publication 198.

## Package explanation
It is a Golang service api that stores a list of whitelisted domain urls each with their own signed hmac signature. The hmac signature uses a shared secret between two parties. In this case its Opera and IXO. The key of the mechanism is Opera can validate that the response from the whitelist server is cryptographically correct by comparing the hmac signature provided with their own signature generated using the urls in the response. They must match or there is likely a man in the middle attack occuring between the opera client and the ixo server. The server address is to be confirmed once a shared secret is agreed upon.

## Package Requirements
It defines a client server relationship where IXO (The server) shares a private secret key with Opera (The client) and Opera uses the key to verify the integrity of the messages being sent. The requirement on the Opera side is to have a hmac comparison function in line with the hmac standard defined above for comparison purposes.

## API Example Response Object
```json
   {
        data: {

            ID: 1,
            CreatedAt: "",
            UpdatedAt: "",
            DeletedAt: null,
            name: "examplename",
            url: "exampleurl",
            hash: "ExampleHmachash"

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
