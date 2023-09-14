//go:build !js || !wasm
// +build !js !wasm

package main

import (
	"flag"

	"code.rocket9labs.com/tslocum/boxcars/game"
)

func parseFlags(g *game.Game) {
	flag.StringVar(&g.Username, "username", "", "Username")
	flag.StringVar(&g.Password, "password", "", "Password")
	flag.StringVar(&g.ServerAddress, "address", game.DefaultServerAddress, "Server address")
	flag.BoolVar(&g.Watch, "watch", false, "Watch random game")
	flag.BoolVar(&g.TV, "tv", false, "Watch random games continuously")
	flag.IntVar(&g.Debug, "debug", 0, "Print debug information")
	flag.Parse()
}
