// +build !js !wasm

package main

import (
	"flag"

	"code.rocketnine.space/tslocum/boxcars/game"
)

func parseFlags(g *game.Game) {
	flag.StringVar(&g.Username, "username", "", "Username")
	flag.StringVar(&g.Password, "password", "", "Password")
	flag.StringVar(&g.ServerAddress, "address", "fibs.com:4321", "Server address")
	flag.BoolVar(&g.Watch, "watch", false, "Watch random game")
	flag.BoolVar(&g.TV, "tv", false, "Watch random games continuously")
	flag.Parse()
}
