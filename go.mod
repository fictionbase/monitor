module github.com/fictionbase/monitor

go 1.12

require (
	github.com/fictionbase/fictionbase v0.0.0-20190510051858-251968169696
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	golang.org/x/sys v0.0.0-20190509141414-a5b02f93d862 // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace (
	github.com/fictionbase/agent => ../agent
	github.com/fictionbase/fictionbase => ../fictionbase
	github.com/fictionbase/router => ../router
)
