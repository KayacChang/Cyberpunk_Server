package gamerule

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation/fileload"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/igame"
)

func TestNew(t *testing.T) {
	for i := 0; i < 200; i++ {

		gamejsStr := fileload.Load("../../file/gameconfig.json")
		var gameRule = &Rule{}
		if err := json.Unmarshal([]byte(gamejsStr), &gameRule); err != nil {
			panic(err)
		}

		fmt.Println(gameRule.newlogicResult(0))
	}
}
func TestGameRequest(t *testing.T) {
	gamejsStr := fileload.Load("../../file/gameconfig.json")
	var gameRule = &Rule{}
	if err := json.Unmarshal([]byte(gamejsStr), &gameRule); err != nil {
		panic(err)
	}

	result := gameRule.GameRequest(&igame.RuleRequest{BetIndex: 0})
	fmt.Println(foundation.JSONToString(result.GameResult))

}
