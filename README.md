# qBittorrent Shell
- Make qBittorrent-nox easier to install, uninstall and reload
- Written in golang 

## Usage
```
Usage:
    qbtsh [Options] [Commands]

Options:
    -ee   : Choose qBittorrent-nox enhanced edition

Command:
    -I    : Install
    -R    : Uninstall
    -U    : Update
    -r    : Reload
    -h    : Show help
    -v    : Show version

Example:
    1) qbtsh -I      : Install default qBittorrent-nox
    2) qbtsh -ee -I  : Install qBittorrent-nox enhanced edition
    3) qbtsh -U      : Update default qBittorrent-nox
    4) qbtsh -ee -U  : Update qBittorrent-nox enhanced edition
    5) qbtsh -R      : Remove cache and uninstall qBittorrent-nox
    6) qbtsh -r      : Reload service
```

## Build
### Example
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/qbtsh.git

cd qbtsh

go build -v -trimpath -ldflags "-s -w"
```
