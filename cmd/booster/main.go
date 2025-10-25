package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    username := flag.String("username", "", "Steam account username (overrides STEAM_USERNAME)")
    password := flag.String("password", "", "Steam account password (overrides STEAM_PASSWORD)")
    appids := flag.String("appids", "", "Comma-separated Steam app IDs to idle (overrides STEAM_APPIDS)")
    maxConcurrent := flag.Int("max-concurrent", 0, "Maximum number of games to idle simultaneously (overrides STEAM_MAX_CONCURRENT)")
    dataDir := flag.String("data-dir", "", "Directory used to store session/sentry data (overrides STEAM_DATA_DIR)")
    socks5 := flag.String("socks5", "", "SOCKS5 proxy URL (overrides STEAM_SOCKS5)")
    heartbeat := flag.Int("heartbeat", 0, "Heartbeat interval in seconds (overrides STEAM_HEARTBEAT_SEC)")
    backoffMin := flag.Int("backoff-min", 0, "Minimum retry backoff in seconds (overrides STEAM_BACKOFF_MIN_SEC)")
    backoffMax := flag.Int("backoff-max", 0, "Maximum retry backoff in seconds (overrides STEAM_BACKOFF_MAX_SEC)")
    logLevel := flag.String("log-level", "", "Log level to use (overrides STEAM_LOG_LEVEL)")
    configFile := flag.String("config", "./config.yaml", "Path to configuration file")
    noRemember := flag.Bool("no-remember", false, "Do not persist session data between runs")
    forgetSession := flag.Bool("forget-session", false, "Forget any stored session before logging in")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    flag.Parse()

    _ = []interface{}{
        username, password, appids, maxConcurrent, dataDir,
        socks5, heartbeat, backoffMin, backoffMax, logLevel,
        configFile, noRemember, forgetSession,
    }
}
