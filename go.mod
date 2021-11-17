module qbtsh

go 1.17

require (
	gek_exec v0.0.0
	gek_file v0.0.0
	gek_service v0.0.0
	gek_toolbox v0.0.0
)

require (
	gek_github v0.0.0 // indirect
	gek_json v0.0.0 // indirect
)

replace (
	gek_exec => ../gek/gek_exec
	gek_file => ../gek/gek_file
	gek_github => ../gek/gek_github
	gek_json => ../gek/gek_json
	gek_service => ../gek/gek_service
	gek_toolbox => ../gek/gek_toolbox
)
