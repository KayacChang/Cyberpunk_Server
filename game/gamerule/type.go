package gamerule

import (
	"github.com/YWJSonic/ServerUtility/foundation"
	"github.com/YWJSonic/ServerUtility/igame"
)

// Rule ...
type Rule struct {
	GameTypeID          string    `json:"GameTypeID"`
	BetLine             int       `json:"BetLine"`
	BetRate             []int64   `json:"BetRate"`
	BetRateDefaultIndex int64     `json:"BetRateDefaultIndex"`
	BetRateLinkIndex    []int64   `json:"BetRateLinkIndex"`
	FreeGameCountArray  []int     `json:"FreeGameCountArray"`
	FreeReel            [][]int   `json:"FreeReel"`
	GameIndex           int       `json:"GameIndex"`
	IsAttachInit        bool      `json:"IsAttachInit"`
	ItemResults         [][][]int `json:"ItemResults"`
	Items               []int     `json:"Items"`
	LineMap             [][]int   `json:"LineMap"`
	NormalReel          [][][]int `json:"NormalReel"`
	NormalReelSize      []int     `json:"NormalReelSize"`
	RandWildCount       []int     `json:"RandWildCount"`
	RandWildWeightings  []int     `json:"RandWildWeightings"`
	RTPSetting          []int     `json:"RTPSetting"`
	ScotterGameLimit    []int     `json:"ScotterGameLimit"`
	ScotterItemIndex    []int     `json:"ScotterItemIndex"`
	Version             string    `json:"Version"`
	WildBonusLimit      []int     `json:"WildBonusLimit"`
	WildBonusRate       []int     `json:"WildBonusRate"`
	WildsItemIndex      []int     `json:"WildsItemIndex"`
	WinBetRateLimit     int64     `json:"WinBetRateLimit"`
	WinScoreLimit       int64     `json:"WinScoreLimit"`
}

// GetGameTypeID ...
func (r *Rule) GetGameTypeID() string {
	return r.GameTypeID
}

// GetBetMoney ...
func (r *Rule) GetBetMoney(index int64) int64 {
	return r.BetRate[index]
}

// GetReel ...
func (r *Rule) GetReel() map[string][][]int {
	scrollmap := map[string][][]int{
		"normalreel": r.normalReel(),
		"freereel":   r.freeReel(),
	}
	return scrollmap
}

// GetBetSetting ...
func (r *Rule) GetBetSetting() map[string]interface{} {
	tmp := make(map[string]interface{})
	tmp["betrate"] = r.BetRate
	tmp["betratelinkindex"] = r.BetRateLinkIndex
	tmp["betratedefaultindex"] = r.BetRateDefaultIndex
	return tmp
}

// CheckGameType ...
func (r *Rule) CheckGameType(gameTypeID string) bool {
	if r.GameTypeID != gameTypeID {
		return false
	}
	return true
}

func (r *Rule) normalReel() [][]int {
	return r.NormalReel[r.normalRTP()]
}

func (r *Rule) normalRTP() int {
	return r.RTPSetting[0]
}

func (r *Rule) freeReel() [][]int {
	return r.FreeReel
}

// Wild1 ...
func (r *Rule) Wild1() int {
	return r.WildsItemIndex[0]
}

func (r *Rule) randWild() [][]int {
	var randwild = make([][]int, len(r.normalReel()))
	var randpoint []int
	var pointArray = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	wildCount := r.RandWildCount[foundation.RangeRandom(r.RandWildWeightings)]

	if wildCount <= 0 {
		return [][]int{}
	}

	randpoint = foundation.RandomMutily(pointArray, wildCount)

	var col = 0
	for _, value := range randpoint {
		col = value / 3
		randwild[col] = append(randwild[col], value%3)
	}

	return randwild
}

// Scotter1 ...
func (r *Rule) Scotter1() int {
	return r.ScotterItemIndex[0]
}

// Scotter1GameLimit ...
func (r *Rule) Scotter1GameLimit() int {
	return r.ScotterGameLimit[0]
}

// FreeGameCount ...
func (r *Rule) FreeGameCount() int {
	return r.FreeGameCountArray[0]
}

// GameRequest ...
func (r *Rule) GameRequest(config *igame.RuleRequest) *igame.RuleRespond {
	betMoney := r.GetBetMoney(config.BetIndex)
	result := make(map[string]interface{})
	var totalWin int64

	gameResult := r.newlogicResult(betMoney)

	result["freeresult"] = make([]map[string]interface{}, 0, 0)
	result["isfreegame"] = 0
	result["freewildbonusrate"] = 0

	result["normalresult"] = gameResult.Normalresult
	totalWin += gameResult.Normaltotalwin

	if gameResult.Freeresult != nil {
		result["freeresult"] = gameResult.Freeresult
		result["isfreegame"] = 1
		result["freewildbonusrate"] = gameResult.Otherdata["freewildbonusrate"]
		totalWin += gameResult.Freetotalwin
	}

	result["totalwinscore"] = totalWin

	return &igame.RuleRespond{
		BetMoney:      betMoney,
		Totalwinscore: totalWin,
		GameResult:    result,
	}
}

// // Result att 0: freecount
// func (r *Rule) logicResult(betMoney int64, JP *jackPart) map[string]interface{} {
// 	var result = make(map[string]interface{})
// 	var totalWin int64
// 	normalresult, otherdata, normaltotalwin := r.outputGame(betMoney, JP)
// 	result = foundation.AppendMap(result, otherdata)
// 	result["normalresult"] = normalresult
// 	totalWin += normaltotalwin
// 	if otherdata["isrespin"].(int) == 1 {
// 		respinresult, respintotalwin := r.outRespin(betMoney, JP)
// 		totalWin += respintotalwin
// 		result["respin"] = respinresult
// 		result["isrespin"] = 1
// 	}
// 	result["totalwinscore"] = totalWin
// 	return result
// }
