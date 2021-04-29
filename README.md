# ovpnstatusd

Downloads *openvpn-status.log* from an [OpenVPN](https://openvpn.net/) server via
[SSH](https://en.wikipedia.org/wiki/Secure_Shell_Protocol) and provides the number of connected clients using a certain
common name scrapable by [Prometheus](https://prometheus.io/).

## Usage

```shell
$ ovpnstatusd -help
Usage of ./ovpnstatusd:
  -destination string
        The hostname/IP address and port of the destination, separated by colon.
  -interval int
        The update interval in milliseconds (default 60000)
  -password string
        The password
  -port int
        The port ovpnstatusd listens on (default 8080)
  -remote-path string
        The path to openvpn-status.log at the destination (default "/etc/openvpn/openvpn-status.log")
  -timeout uint
        The timeout in seconds (default 5)
  -user string
        The usernam
```

When running, ovpnstatusd listens to the specified port and provides the metrics at the
`/metrics` endpoint.

### via Docker

```shell
docker run toolcreator/ovpnstatusd
```

The arguments shown above can also be passed, e.g.:

```shell
docker run toolcreator/ovpnstatusd -interval 10000 -port 12345
```

Using the `-port` option may be particularly useful when running the container with `--network=host`
(i.e., when port mapping is not available).
When available, port mapping can of course be used as well, e.g.:

```shell
docker run -p 12345:8080 toolcreator/ovpnstatusd
```

#### docker-compose

```yml
ovpnstatusd:
  image: toolcreator/ovpnstatusd
  command:
    - "-interval=10000"
  ports:
    - '12345:8080'
```

Or, with `network_mode: "host"`:

```yml
ovpnstatusd:
  image: toolcreator/ovpnstatusd
  network_mode: "host"
  command:
    - "-interval=10000"
    - "-port=12345"
```

### Without Docker

1. Download the source code: `git clone https://github.com/toolcreator/ovpnstatusd.git`
2. Enter the root directory: `cd ovpnstatusd`
3. Compile: `go build`
4. Install: `go install`
5. Run: `ovpnstatusd`

You may also skip step 4 and use `./ovpnstatusd` to run the program instead.

## Metrics

All metrics are of type [gauge](https://prometheus.io/docs/concepts/metric_types/#gauge)
and are labeled with `common_name`.

| Name          | Description   |
| ------------- | ------------- |
| ovpnstatusd_client_count | The number of clients connected using the common name |
