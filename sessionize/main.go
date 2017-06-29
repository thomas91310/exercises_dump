package main

import "os"

//resolve https://github.com/MediaMath/DataScienceTeam/issues/41

func main() {
	app := newApp()
	app.Run(os.Args)
}
