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
- 使 `qBittorrent-nox` 在 `Linux` 内核系统上更容易安装、卸载及重载。
- 采用我自己编写的一套应用安装框架 [gek_app](https://github.com/gek64/gek/tree/main/gek_app)，编写应用时依照系统文件与第三方软件文件分离的思想，应用安装在 `/usr/local/bin/qbittorrent-nox`，应用数据存储在 `/usr/local/etc/qBittorrent/`，避免将来系统变动带来的问题。
- 静态编译的 `qBittorrent-nox` 使用 [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)。

## 使用说明
```
使用说明:
    qbtsh [指令]
指令:
    -install    : 安装
    -uninstall  : 卸载
    -update     : 升级
    -reload     : 重载
    -h          : 显示帮助
    -v          : 显示版本
实例:
    1) qbtsh -install     : 安装 qBittorrent-nox
    3) qbtsh -update      : 升级 qBittorrent-nox
    5) qbtsh -uninstall   : 卸载 qBittorrent-nox，并移除应用配置文件及应用缓存
    6) qbtsh -reload      : 重载服务
```

## 安装说明
```sh
# 例如宿主机是 amd64 架构的基于 linux 内核的系统
# 下载应用
curl -Lo /usr/local/bin/qbtsh https://github.com/gek64/qbtsh/releases/latest/download/qbtsh-linux-amd64
# 给应用执行权限
chmod +x /usr/local/bin/qbtsh
# 查看帮助
/usr/local/bin/qbtsh -h
```

## 编译说明
```sh
# 下载应用源代码
git clone https://github.com/gek64/qbtsh.git
# 编译源代码
cd qbtsh
go build -v -trimpath -ldflags "-s -w"
```

## 常见问答
### 本工具是否支持 FreeBSD?
- 不支持，[userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases) 仅支持 `linux` 内核的系统, 并且在 `FreeBSD` 上编译一个完全静态的 `qBittorrent-nox` 是一件相当困难的事情。

### 有什么办法能在FreeBSD上使用吗?
- 你可以使用我为FreeBSD系统编写的 `qBittorrent-nox` 编译脚本来尝试自行编译动态链接的 `qBittorrent-nox` [freebsd build](https://github.com/gek64/qbittorrent-nox)，我已经为最新的FreeBSD系统编译好了 `amd64` 架构的二进制文件，你也可以直接使用。

### 运行以后该如何访问?
- 服务成功运行后访问 `SERVER_IP:8080` 即可，默认用户名是 `admin`，默认密码是 `adminadmin` ，虽然此时你已经可以正常使用，但是我还是建议你及时更改密码，并在开启了 `https` 的 `apache` 或者 `nginx` 等网页服务器下配置反向代理，并将 `qBittorrent-nox` 的监听地址从 `0.0.0.0` 更变为 `127.0.0.1` 以提高安全性。

## 许可证
- **GPL-3.0 License**
- 查看 `LICENSE` 获取详细内容

## 致谢
- [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
