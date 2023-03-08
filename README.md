# V2RAY with ECH support
This is an experimental project that adds ECH support to [v2ray project](https://github.com/v2fly/v2ray-core).

If your domain is behind a CDN and has been censored, you can use this project to bypass it.
Currently this project only supports Cloudflare CDN.

There are two steps that your domain send in plain text: In DNS query and TLS handshake.
For concealing the domain in DNS query, you can use DNS over HTTPS or hardcode the ip of your domain (which is the CDN IP) in /ets/hosts.
For concealing the domain in TLS handshake, you should use ECH.
ECH is an approach that can hide SNI or other critical info in TLS handshake. You can learn more about ECH [here](https://blog.cloudflare.com/handshake-encryption-endgame-an-ech-update).

# How to Build
You should use [cloudflare golang](https://github.com/cloudflare/cloudflare-go) for building the project.
Which it supports ECH for TLS handshake.

Just clone the project, build it and replace your `/usr/local/go` folder with that.
Something like scripts below:
```bash
git clone https://github.com/cloudflare/cloudflare-go go && cd go/src
./all.bash
```
And then copy this `go` directory to your `/usr/local/go`

After installing cloudflare golang, you can build this project. Somthing like below scripts:
```bash
git clone https://github.com/imannamdari/v2ray-core.git && cd v2ray-core
go build -o v2ray ./main/main.go
```

# How to Use
Build the project or download the binary. Append these below configs to your `tlsSettings` (in v4 v2ray config) or `securitySettings` (in v5 v2ray config):

v5 config:
```json
"securitySettings": {
    "serverName": "your-domain.com",
    "enableEch": true,
    "echSetting": {
        "dnsAddr": "1.1.1.1:53"
    }
}
```
v4 config:
```json
"tlsSettings": {
    "serverName": "your-domain.com",
    "enableEch": true,
    "echSetting": {
        "dnsAddr": "1.1.1.1:53"
    }
}
```

By default this project use `127.0.0.1:53` for dns lookup. (in the case you want to use DNS over HTTPS) This lookup is only used for getting ECH keys and will not be used for other dns lookups. You can change this dns addr with `echSetting` config.
