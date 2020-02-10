package chess_state

import (
	"log"
	"sync"
)

// Observer should use game.Game as read only
type Observer interface {
	Update(*Game, *Move) error
}

type Observers struct {
	observers []Observer
	mutex     sync.Mutex
}

func (o *Observers) Add(observer Observer) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.observers = append(o.observers, observer)
}

// Update calls update on all observers, removing observers that return error.
func (o *Observers) Update(game *Game, move *Move) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	errChan := make(chan error)
	for _, observer := range o.observers {
		observer := observer
		go func() {
			errChan <- observer.Update(game, move)
		}()
	}

	// i := 0
	for range o.observers {
		if err := <-errChan; err != nil {
			log.Println(err)
			continue
		}
		// TODO: figure out if we need to remove observer on error
		// o.observers[i] = observer
		// i++
	}
	// o.observers = o.observers[:i]
}

type LoggingObserver struct{}

func (o *LoggingObserver) Update(game *Game, move *Move) error {
	log.Println(game.Game.Position().Board().Draw())
	return nil
}
