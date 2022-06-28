# IXOWhitelistDaemon

IXOWhitelistDaemon is a simple RSA PKSS based golang application to serve  verifiable domain whitelists.

## Installation

Requires a local golang development setup then make server can be used to start a local copy.
Keys are generated at runtime unless a private.txt and public.txt is present in the local binary directory.



## Usage

```bash


# Starts a local development build
make server

# Builds the binary
make build

# docker-compose --build up
make d.up.build

# docker-compose u
make d.up

# docker-compose down
make d.down

```

## Features

```bash
#Gets all the domain whitelist items
/api/getwhitelist
#Creates a new domain whitelist item
/api/createwhitelistitem
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)