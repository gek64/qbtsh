module qbtsh

go 1.17

require (
	gek_app v0.0.0
	gek_toolbox v0.0.0
)

require (
	gek_downloader v0.0.0 // indirect
	gek_exec v0.0.0 // indirect
	gek_file v0.0.0 // indirect
	gek_github v0.0.0 // indirect
	gek_json v0.0.0 // indirect
	gek_math v0.0.0 // indirect
	gek_service_rc v0.0.0 // indirect
	gek_service_systemd v0.0.0 // indirect
)

replace (
	gek_app => ../gek/gek_app
	gek_downloader => ../gek/gek_downloader
	gek_exec => ../gek/gek_exec
	gek_file => ../gek/gek_file
	gek_github => ../gek/gek_github
	gek_json => ../gek/gek_json
	gek_math => ../gek/gek_math
	gek_service_rc => ../gek/gek_service_rc
	gek_service_systemd => ../gek/gek_service_systemd
	gek_toolbox => ../gek/gek_toolbox
)
