package master

import (
	"github.com/lonng/nano/session"
	"github.com/nano/gameserver/db/model"
)

type User struct {
	session  *session.Session
	data     *model.User
	Uid      int64
	heroData *model.Hero
}
