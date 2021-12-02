# qBittorrent Shell
- Make qBittorrent-nox easier to install, uninstall and reload on linux
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
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/qbtsh.git

cd qbtsh

go build -v -trimpath -ldflags "-s -w"
```

## QA

### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`
A: This report occurred after `Windows 10 21h2`. This application does not contain any malware, backdoors, and advertisements, all released files are build by github actions. This problem can be solved by adding this application to the whitelist or compiling by yourself.

### Q: Why should I clone `https://github.com/gek64/gek.git` before building
A: I don’t want the project to depend on a certain cloud service provider, and this is also a good way to avoid network problems.


## License

**GNU Lesser General Public License v2.1**

See `LICENSE` for details
