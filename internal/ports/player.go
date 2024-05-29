package ports

import "github.com/jairogloz/go-l/internal/domain"

type PlayerService interface {
	Create(player domain.Player) (id interface{}, err error)
	Get(id domain.Player) (player domain.Player, err error)
	Delete(id domain.Player) (err error)
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
	Get(id domain.Player) (player domain.Player, err error)
	Delete(id domain.Player) (err error)
}
