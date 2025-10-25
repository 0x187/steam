package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    var configFile string
    flag.StringVar(&configFile, "config", "config.json", "Path to configuration file")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    flag.Parse()

}
