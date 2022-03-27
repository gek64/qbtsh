# qBittorrent Shell
- Make qBittorrent-nox easier to install, uninstall and reload on linux
- static qBittorrent-nox use [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)

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
# for example (system is linux and arch is amd64)
curl -Lo /usr/local/bin/qbtsh https://github.com/gek64/qbtsh/releases/latest/download/qbtsh-linux-amd64
chmod +x /usr/local/bin/qbtsh
/usr/local/bin/qbtsh -h
```

## Compile
### How to compile if prebuilt binaries are not found
```sh
git clone https://github.com/gek64/gek.git
git clone https://github.com/gek64/qbtsh.git
cd qbtsh
go build -v -trimpath -ldflags "-s -w"
```

## QA
### Q: FreeBSD support?
A: [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases) only support linux, and it is also hard to build a full static qBittorrent-nox for FreeBSD, you can use my dynamic [freebsd build](https://github.com/gek64/qbittorrent-nox).

## License
- **GNU Lesser General Public License v2.1**
- See `LICENSE` for details
