# CaddyBanIP

## A Caddy 2 module that allows you to ban IPs from visiting a part or the entirety of your website.

[![Go Report Card](https://goreportcard.com/badge/github.com/DrivetDevelopment/CaddyBanIP)](https://goreportcard.com/report/github.com/DrivetDevelopment/CaddyBanIP)

## Install
Prerequisites:
- [xcaddy](https://github.com/caddyserver/xcaddy)
- a server

Run this command in your terminal to build a custom version of Caddy that includes CaddyBanIP.

``xcaddy build --with github.com/DrivetDevelopment/CaddyBanIP``

I recommend you to remove any other versions of Caddy you have installed, like the stock version of Caddy.

Linux? You might already have a systemd service for Caddy. Let's replace the binary:

```bash
mv caddy /usr/bin/caddy
chmod +x /usr/bin/caddy

systemctl restart caddy # if you have a systemd service
```

## Setup
Add this to your Caddyfile:

```
{
    order caddybanip before file_server
}
```

This is, unfortunately, required in Caddy. It tells Caddy to load the CaddyBanIP module before the file_server module.

Now, add the following to your Caddyfile:

```
(bannedIPs) {
        caddybanip {
            banned_ips "127.0.0.1|192.168.10.164" # IPs you do not want to visit this path. Note: Regex
            message "You are banned from visiting this website!" # Optional, but recommended if you think the default message is an advertisement or something.
        }
}

yourwebsite.com {
    import bannedIPs
}
```

You can also ban specific IPs from accessing paths, with the [handle_path](https://caddyserver.com/docs/caddyfile/directives/handle_path) directive.

## Users of CaddyBanIP
Do you use CaddyBanIP? Let me know by opening an issue or a pull request to add yourself into this list!

- [Odyssey346](https://odyssey346.dev)