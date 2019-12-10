module github.com/eliukblau/pixterm/cmd/pixterm

go 1.13

require (
	github.com/eliukblau/pixterm/pkg/ansimage v0.0.0
	github.com/lucasb-eyer/go-colorful v1.0.3
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
)

replace github.com/eliukblau/pixterm/pkg/ansimage => ../../pkg/ansimage
