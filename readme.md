# bothan

Confirm an ip:port is hosting Empire. Eventual support for other C2s.

## Usage

```bash
# Direct single query
❯❯ bothan -v localhost:8080
DEBU[2020-01-19T20:25:02-05:00] Requesting...                                 host="localhost:8080"
INFO[2020-01-19T20:25:02-05:00] SUCCESS                                       host="http://localhost:8080" tool=empire

# Take a pre-existing list of host:port lines
❯❯ bothan -v -f hostslist.txt
DEBU[2020-01-19T20:25:10-05:00] Requesting...                                 host="localhost:8080"
INFO[2020-01-19T20:25:10-05:00] SUCCESS                                       host="http://localhost:8080" tool=empire

# Take Stdin
❯❯ cat masscan.oD.txt | jq -r '. | "\(.ip):\(.port)"' | bothan -f -
ERRO[2020-01-19T20:25:25-05:00] Get https://1.1.1.1:53: EOF                   host="1.1.1.1:53"
INFO[2020-01-19T20:25:25-05:00] SUCCESS                                       host="http://192.168.1.199:8080" tool=empire

# For masscan specifically, there's an option to parse its -oD json output format
❯❯ masscan 192.168.1.0/24 -p 8080 -oD - | bothan --masscan -f -
INFO[2020-01-19T20:25:31-05:00] SUCCESS                                       host="http://192.168.1.199:8080" tool=empire
```

Successes are written to Stdout, all other logs, Stderr.

## Install

```bash
go get github.com/audibleblink/bothan
```

## Build

1. Have `go`
2. Have `make`
3. Type `make`

```
bin
├── 386
│   ├── bothan.darwin
│   ├── bothan.linux
│   └── bothan.windows.exe
├── amd64
│   ├── bothan.darwin
│   ├── bothan.linux
│   └── bothan.windows.exe
├── arm
│   └── bothan.linux
└── arm64
    └── bothan.linux

4 directories, 8 files
```

## False Positives

For testing while developing this, I initially ran Empire in its default state, and then
customized. After that, I relied on servers listed by public threat intel feeds to test against;
around 100.

98 were identified as Empire. One just timed out and the other looked heavily modified.

That is to say, I didn't have what I'd consider a statistically relevant sample set to test
against, so please report any false {positives,negatives} you encounter.
