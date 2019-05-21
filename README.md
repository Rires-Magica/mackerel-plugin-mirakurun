# mackerel-plugin-mirakurun
A Mackerel plugin for monitoring tuners that managed by mirakurun.
## Usage
```shell
./mackerel-plugin-mirakurun [-mirakurun-port=<port>] [-metric-key-prefix=<prefix>]
```
All of these options are optional.
## Install
```shell
mkr plugin install Rires-Magica/mackerel-plugin-mirakurun@(version)
```
```toml:/etc/mackerel-agent/mackerel-agent.conf
[plugin.metrics.mirakurun]
command = ["/path/to/mackerel-plugin-mirakurun", "-mirakurun-port=<port>", "-metric-key-prefix=<prefix>"]
```
## License
MIT.
