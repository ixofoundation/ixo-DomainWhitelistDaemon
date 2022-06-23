# IXOWhitelistDaemon

IXOWhitelistDaemon is a simple golang application to serve domain whitelists.

## Installation

Requires a local golang development setup then make server can be used to start a local copy.
The local .env file needs to be set with a local server secret to be shared with the partner organization for domain verification purposes.



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