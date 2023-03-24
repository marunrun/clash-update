module clash-update

go 1.18

replace CloudflareSpeedTest => ./lib/CloudflareSpeedTest

require CloudflareSpeedTest v0.0.0-00010101000000-000000000000

require (
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/cheggaaa/pb/v3 v3.0.4 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	golang.org/x/sys v0.0.0-20191128015809-6d18c012aee9 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
