```
          ▄▄                            ▄▄
          ██         ██                 ███
          ██         ██                 ██
 ▄██▀███  ██▄████▄ ██████   ▄██▀███   ███████▄
▄█▀   ██  ██    ▀██  ██     ██        ██   ██
██    ██  ██     ██  ██     ▀██████▄  ██   ██
██▄   ██  ██▄   ▄██  ██     █▄   ███  ██   ██
 ▀██████  █▀█████▀   ▀████  ██████▀   ██  ████▄
      ██
   ▄████▄
```

[中文说明](https://github.com/gek64/qbtsh/blob/main/README_chs.md)

- Make `qBittorrent-nox` easier to install, uninstall and reload on `Linux` kernel system.
- Using an application installation framework [gek_app](https://github.com/gek64/gek/tree/main/gek_app) written by
  myself, when writing this application, according to the idea of separating system files from third-party software
  files, the application is installed in `/usr/local/bin/qbittorrent-nox`, and the application data is stored
  in `/usr/local/etc/qBittorrent/` to avoid problems caused by future system changes.
- Statically compiled `qBittorrent-nox`
  using [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)

## Usage

```
Usage:
    qbtsh [Commands]
Command:
    -install    : Install
    -uninstall  : Uninstall
    -update     : Update
    -reload     : Reload
    -h          : Show help
    -v          : Show version
Example:
    1) qbtsh -install     : Install qBittorrent-nox
    3) qbtsh -update      : Update qBittorrent-nox
    5) qbtsh -uninstall   : Remove config,cache and uninstall qBittorrent-nox
    6) qbtsh -reload      : Reload service
```

## Install

```sh
# For example, the host is a linux kernel based system with amd64 architecture
# Download the app
curl -Lo /usr/local/bin/qbtsh https://github.com/gek64/qbtsh/releases/latest/download/qbtsh-linux-amd64
# Give the app execute permission
chmod +x /usr/local/bin/qbtsh
# Show help
/usr/local/bin/qbtsh -h
```

## Compile

```sh
# Download application source code
git clone https://github.com/gek64/qbtsh.git
# Compile the source code
cd qbtsh
go build -v -trimpath -ldflags "-s -w"
```

## FAQ

### Does this tool support FreeBSD?

- Not supported, [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases) only
  supports `linux` kernel systems, and builds a fully static `qBittorrent-nox` on `FreeBSD` is quite a difficult thing.

### Is there any way to get this working on FreeBSD?

- You can try compiling a dynamically linked `qBittorrent-nox` yourself using the `qBittorrent-nox` build script I wrote
  for `FreeBSD` [freebsd build](https://github.com/gek64/qbittorrent-nox), also I have already compiled binaries for the
  latest `FreeBSD`/`amd64`, you can also use them directly.

### How to access after running?

- After the service runs successfully, you can access it at `SERVER_IP:8080`. The default user name is `admin`, and the
  default password is `adminadmin`. Although you can use it normally at this time, I still recommend you to change the
  password in time and use reverse proxy under the web server such as `apache` or `nginx` with `https` enabled, then
  change the listening address of `qBittorrent-nox` from `0.0.0.0` to `127.0.0.1` to improve security.

## License

- **GPL-3.0 License**
- See `LICENSE` for details

## Credits

- [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
