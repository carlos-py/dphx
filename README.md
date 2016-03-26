# dphx

Hi.


## Config

Config is loaded from the process environment.

A sample settings file along with options description and example values is listed below.

```bash
# SSH server address.
export DPHX_SSH_ADDR='example.net'

# SSH username.
export DPHX_SSH_USER='user'

# SSH password (optional).
export DPHX_SSH_PASSWORD='password'

# SSH public key files list (optional).
export DPHX_SSH_PUBLIC_KEYS='/home/me/.ssh/id_rsa,/home/me/.ssh/other_id_rsa'

# SSH agent address, usually a unix socket.
# Defaults to $SSH_AUTH_SOCK value.
export DPHX_SSH_AGENT='/run/user/1337/keyring-blah/ssh'

# SOCKS5 server bind address (optional).
# Defaults to "127.0.0.1:1080".
export DPHX_SOCKS_ADDR='0.0.0.0:1080'
```
