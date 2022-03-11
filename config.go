package main

var (
	// 工具链
	toolbox = []string{"service", "systemd"}

	// 应用相关
	repo     = "userdocs/qbittorrent-nox-static"
	repoList = map[string]string{
		"linux_amd64": "x86_64-qbittorrent-nox",
		"linux_arm":   "armv7-qbittorrent-nox",
		"linux_arm64": "aarch64-qbittorrent-nox",
	}
	bins                             = []string{"qbittorrent-nox"}
	binsLocation                     = "/usr/local/bin"
	binNeedExtract                   = false
	binUninstallDeleteLocationFolder = false

	// 配置文件相关
	configName                          = ""
	configContent                       = ""
	configLocation                      = "/usr/local/etc/qBittorrent"
	configUninstallDeleteLocationFolder = true

	// 服务相关
	serviceName    = "qbittorrent.service"
	serviceContent = `# nano /etc/systemd/system/qbittorrent.service
# systemctl enable qbittorrent && systemctl start qbittorrent

[Unit]
Description=qBittorrent Daemon Service
Wants=network-online.target
After=network-online.target nss-lookup.target

[Service]
Type=exec
User=root
ExecStart=/usr/local/bin/qbittorrent-nox --profile="/usr/local/etc/"
Restart=on-failure
SyslogIdentifier=qbittorrent

[Install]
WantedBy=multi-user.target`
)
