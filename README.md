# ovpnstatusd

## Usage

```shell
$ ovpnstatusd -help
Usage of ovpnstatusd:
  -interval int
        The update interval in milliseconds (default 1000)
  -port int
        The port ovpnstatusd listens on (default 8080)
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
and are labeled with `example_label`.

| Name          | Description   |
| ------------- | ------------- |
| example_gauge | Example Gauge |
