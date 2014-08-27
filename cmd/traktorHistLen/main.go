package main

import (
	"encoding/xml"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/cryptix/goNITraktor"
)

func main() {
	app := cli.NewApp()
	app.Name = "History Analyzer"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "file,f"},
	}
	app.Action = func(ctx *cli.Context) {
		fname := ctx.String("file")
		if fname == "" {
			log.Fatal("no hist file specified")
		}
		input, err := os.Open(fname)
		check(err)
		defer input.Close()

		var histRoot goNITraktor.NmlRoot
		err = xml.NewDecoder(input).Decode(&histRoot)
		check(err)

		// log.Println("Head:", histRoot.Head)
		log.Println("Collection Entries:", histRoot.Collection.EntryCnt)

		pnodes := len(histRoot.Playlists.Nodes)
		if pnodes < 1 || histRoot.Playlists.Nodes[0].Name != "$ROOT" {
			log.Fatalln("No $ROOT Playlist Node")
		}

		subnodes := len(histRoot.Playlists.Nodes[0].Subnodes)
		if subnodes < 1 || histRoot.Playlists.Nodes[0].Subnodes[0].Name != "HISTORY" {
			log.Fatalln("No HISTORY Playlist Subnode")
		}

		hist := histRoot.Playlists.Nodes[0].Subnodes[0]
		log.Println("History Playlsit Length:", hist.Playlist.EntryCnt)

	}
	app.Run(os.Args)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
