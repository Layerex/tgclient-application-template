# tgclient-application-template

Template for building go applications with [tgclient](https://github.com/3bl3gamer/tgclient) library.

## Template customization

```sh
sed * -i -e 's/tgclient-application-template/<application-name>/g' -e 's/Layerex/<your GitHub username>/g'
```

Don't forget to update libraries before proceeding to development:

```sh
go get -u && go mod tidy
```

## Running

```sh
go build
./tgclient-application-template
```

## Usage

```text
usage: ./tgclient-application-template [-h] [--dont-save-session] [--app-id APP_ID] [--app-hash APP_HASH]

<description>

options:
  -h, --help            Show this help message and exit
  --dont-save-session   Don't save session file (and don't use already saved one)
  --app-id APP_ID       Test credentials are used by default
  --app-hash APP_HASH   Test credentials are used by default

Session file is saved to /home/user/.local/share/tgclient-application-template/tg.session

<epilog>
```
