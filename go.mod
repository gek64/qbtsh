module qbtsh

go 1.17

require (
	gek_downloader v0.0.0
	gek_exec v0.0.0
	gek_file v0.0.0
	gek_github v0.0.0
	gek_service v0.0.0
	gek_toolbox v0.0.0
)

replace (
	gek_downloader => ../gek/gek_downloader
	gek_exec => ../gek/gek_exec
	gek_file => ../gek/gek_file
	gek_github => ../gek/gek_github
	gek_json => ../gek/gek_json
	gek_service => ../gek/gek_service
	gek_toolbox => ../gek/gek_toolbox
)
