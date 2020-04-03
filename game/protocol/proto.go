package protocol

import (
	"net/http"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/myhttp"
)

// GameRequest ...
type GameRequest struct {
	Token      string
	BetIndex   int64
	GameTypeID string
	PlayerID   int64
}

// InitData ...
func (c *GameRequest) InitData(r *http.Request) {
	postData := myhttp.PostData(r)
	c.Token = r.Header.Get("Authorization")
	c.BetIndex = foundation.InterfaceToInt64(postData["bet"])
	c.GameTypeID = foundation.InterfaceToString(postData["gametypeid"])
}

// // Respon ...
// type Respon struct {

// }

// // InitData ...
// func (c *Respon) InitData(r *http.Request) {
// 	postData := myhttp.PostData(r)
// 	c.Token = foundation.InterfaceToString(postData["token"])
// 	c.BetIndex = foundation.InterfaceToInt64(postData["bet"])
// 	c.GameTypeID = foundation.InterfaceToString(postData["gametypeid"])
// 	c.PlayerID = foundation.InterfaceToInt64(postData["playerid"])
// }
