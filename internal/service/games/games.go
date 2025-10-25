package games

import (
    "context"
    "log"
    "strconv"
    "sync"
    "time"
)

// SteamClient defines the interface for interacting with the Steam client.
type SteamClient interface {
    PlayGames(appIDs []uint32)
}

// GamesService manages the list of games being played.
type GamesService struct {
    steamClient         SteamClient
    heartbeatInterval   time.Duration
    maxGames            int
    currentAppIDs       []uint32
    mutex               sync.RWMutex
    cancelHeartbeat     context.CancelFunc
    logger              *log.Logger
}

// NewGamesService creates a new GamesService.
func NewGamesService(steamClient SteamClient, heartbeatInterval time.Duration, maxGames int, logger *log.Logger) *GamesService {
    return &GamesService{
        steamClient:       steamClient,
        heartbeatInterval: heartbeatInterval,
        maxGames:          maxGames,
        logger:            logger,
    }
}

// Start begins the games service, including the heartbeat mechanism.
func (s *GamesService) Start(ctx context.Context) {
    s.logger.Println("Games service started")
    go func() {
        <-ctx.Done()
        s.Stop()
    }()
}

// Stop gracefully shuts down the games service.
func (s *GamesService) Stop() {
    s.logger.Println("Stopping games service")
    if s.cancelHeartbeat != nil {
        s.cancelHeartbeat()
    }

    s.mutex.Lock()
    s.logger.Println("Clearing played games on shutdown.")
    s.currentAppIDs = []uint32{}
    s.steamClient.PlayGames(s.currentAppIDs)
    s.mutex.Unlock()

    s.logger.Println("Games service stopped")
}

// SetGames sets the list of games to be played.
func (s *GamesService) SetGames(appIDs []string) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    s.logger.Printf("Received AppIDs: %v", appIDs)

    processedAppIDs := s.processAppIDs(appIDs)

    s.currentAppIDs = processedAppIDs
    s.logger.Printf("Updated AppIDs: %v", s.currentAppIDs)

    s.steamClient.PlayGames(s.currentAppIDs)
    s.resetHeartbeat()
}

// processAppIDs validates, cleans, and truncates the AppID list.
func (s *GamesService) processAppIDs(appIDs []string) []uint32 {
    if len(appIDs) == 0 {
        return []uint32{}
    }

    uniqueAppIDs := make(map[uint32]struct{})
    var validAppIDs []uint32

    for _, appIDStr := range appIDs {
        appID, err := s.parseAppID(appIDStr)
        if err != nil {
            s.logger.Printf("Invalid AppID '%s': %v", appIDStr, err)
            continue
        }
        if _, exists := uniqueAppIDs[appID]; !exists {
            uniqueAppIDs[appID] = struct{}{}
            validAppIDs = append(validAppIDs, appID)
        }
    }

    if len(validAppIDs) > s.maxGames {
        s.logger.Printf("Truncating AppID list from %d to %d games", len(validAppIDs), s.maxGames)
        return validAppIDs[:s.maxGames]
    }

    return validAppIDs
}

// parseAppID converts a string to a uint32.
func (s *GamesService) parseAppID(appIDStr string) (uint32, error) {
    appID, err := strconv.ParseUint(appIDStr, 10, 32)
    if err != nil {
        return 0, err
    }
    return uint32(appID), nil
}

// resetHeartbeat cancels the existing heartbeat and starts a new one.
func (s *GamesService) resetHeartbeat() {
    if s.cancelHeartbeat != nil {
        s.cancelHeartbeat()
    }

    ctx, cancel := context.WithCancel(context.Background())
    s.cancelHeartbeat = cancel

    go s.heartbeat(ctx)
}

// heartbeat periodically sends the current game list to the Steam client.
func (s *GamesService) heartbeat(ctx context.Context) {
    if s.heartbeatInterval == 0 {
        return // No heartbeat
    }

    ticker := time.NewTicker(s.heartbeatInterval)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            s.mutex.RLock()
            s.logger.Println("Sending heartbeat...")
            s.steamClient.PlayGames(s.currentAppIDs)
            s.mutex.RUnlock()
        case <-ctx.Done():
            s.logger.Println("Heartbeat stopped.")
            return
        }
    }
}
