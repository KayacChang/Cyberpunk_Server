package gamerule

import (
	"fmt"

	"github.com/YWJSonic/ServerUtility/foundation"
	"github.com/YWJSonic/ServerUtility/gameplate"
)

type result struct {
	Normalresult   map[string]interface{}
	Otherdata      map[string]interface{}
	Normaltotalwin int64
	Freeresult     []map[string]interface{}
	Freetotalwin   int64
}

// Result att 0: freecount
func (r *Rule) newlogicResult(betMoney int64) result {

	option := gameplate.PlateOption{
		Scotter: []int{r.Scotter1()},
		Wild:    []int{r.Wild1()},
	}

	normalresult, otherdata, normaltotalwin := r.outputGame(betMoney, option)
	fmt.Println("----normalresult----", normalresult)
	fmt.Println("----otherdata----", otherdata)
	fmt.Println("----normaltotalwin----", normaltotalwin)
	// result["normalresult"] = normalresult
	// result["isfreegame"] = 0
	// totalWin += normaltotalwin

	if iscotter, ok := otherdata["isfreegame"]; ok && iscotter.(int) == 1 {
		freeresult, freeotherdata, freetotalwin := r.outputFreeGame(betMoney, option)
		fmt.Println("----freeresult----", freeresult)
		fmt.Println("----freeotherdata----", freeotherdata)
		fmt.Println("----freetotalwin----", freetotalwin)
		// result["freeresult"] = freeresult
		otherdata["freewildbonusrate"] = freeotherdata["freewildbonusrate"]
		// result["isfreegame"] = 1
		// totalWin += freetotalwin
		return result{
			Normalresult:   normalresult,
			Otherdata:      otherdata,
			Normaltotalwin: normaltotalwin,
			Freeresult:     freeresult,
			Freetotalwin:   freetotalwin,
		}
	}

	// result["totalwinscore"] = totalWin
	return result{
		Normalresult:   normalresult,
		Otherdata:      otherdata,
		Normaltotalwin: normaltotalwin,
	}
}

// outputGame out put normal game result, mini game status, totalwin
func (r *Rule) outputGame(betMoney int64, option gameplate.PlateOption) (map[string]interface{}, map[string]interface{}, int64) {
	var totalScores int64
	normalResult := make(map[string]interface{})
	otherdata := make(map[string]interface{})

	randWild := r.randWild()
	normalResult, otherdata, totalScores = r.aRound(betMoney, r.normalReel(), randWild, option, 1)
	normalResult["randwild"] = randWild
	// normalResult["randwild"] = [][]int{}

	return normalResult, otherdata, totalScores
}

func (r *Rule) outputFreeGame(betMoney int64, option gameplate.PlateOption) ([]map[string]interface{}, map[string]interface{}, int64) {
	var totalScores int64
	var wildCount, bonusRate int
	otherdata := make(map[string]interface{})
	var freeResult []map[string]interface{}
	var lockWildarray = make([][]int, len(r.NormalReelSize))

	for i, imax := 0, r.FreeGameCount(); i < imax; i++ {
		tmpResult, _, tmpTotalScores := r.aRound(betMoney, r.freeReel(), lockWildarray, option, 2)
		totalScores += tmpTotalScores
		freeResult = append(freeResult, tmpResult)

		lockWildarray = r.lockWild(tmpResult["plate"].([][]int), lockWildarray, option)
	}
	for _, colArray := range lockWildarray {
		wildCount += len(colArray)
	}
	// freeWildCount[fmt.Sprintf("%v", wildCount)]++

	for limitIndex, limitCount := range r.WildBonusLimit {
		if wildCount >= limitCount {
			bonusRate = r.WildBonusRate[limitIndex]
		}
	}
	if bonusRate > 0 {
		totalScores *= int64(bonusRate)
		otherdata["freewildbonusrate"] = bonusRate
	} else {
		otherdata["freewildbonusrate"] = 0
	}
	return freeResult, otherdata, totalScores
}

func (r *Rule) aRound(betMoney int64, scorll [][]int, randWild [][]int, option gameplate.PlateOption, gameType int) (map[string]interface{}, map[string]interface{}, int64) {

	var totalScores int64
	winLineInfo := []interface{}{}
	otherdata := make(map[string]interface{})
	result := make(map[string]interface{})

	plateIndex, plateSymbol := gameplate.NewPlate2D(r.NormalReelSize, scorll)

	// set random wild
	plateSymbolInsertWild := r.setRandomWild(plateSymbol, randWild)
	plateLineMap := gameplate.PlateToLinePlate(plateSymbolInsertWild, r.LineMap)

	for lineIndex, plateLine := range plateLineMap {
		newLine := gameplate.CutSymbolLink(plateLine, option) // cut line to win line point
		for _, payLine := range r.ItemResults[len(newLine)] { // win line result group
			if r.isWin(newLine, payLine, option) { // win result check
				// if gameType == 1 {
				// 	normalPayLineCount[fmt.Sprintf("%v", payLine)]++
				// } else {
				// 	freePayLineCount[fmt.Sprintf("%v", payLine)]++
				// }

				infoLine := gameplate.NewInfoLine()

				for i, max := 0, len(payLine)-1; i < max; i++ {
					infoLine.AddNewPoint(newLine[i], r.LineMap[lineIndex][i], option)
				}
				infoLine.LineWinIndex = lineIndex
				infoLine.LineWinRate = payLine[len(payLine)-1]
				infoLine.Score = int64(infoLine.LineWinRate) * (betMoney / int64(r.BetLine))
				totalScores += infoLine.Score
				winLineInfo = append(winLineInfo, infoLine)
			}
		}
	}

	plateSymbolCollectResult := gameplate.PlateSymbolCollect(r.Scotter1(), plateSymbolInsertWild, option, map[string]interface{}{
		"isincludewild":   false,
		"isseachallplate": true,
	})
	scotterCount := foundation.InterfaceToInt(plateSymbolCollectResult["targetsymbolcount"])
	scotterLineSymbol := plateSymbolCollectResult["symbolnumcollation"].([][]int)
	scotterLinePoint := plateSymbolCollectResult["symbolpointcollation"].([][]int)

	if scotterCount >= r.Scotter1GameLimit() {
		infoLine := gameplate.NewInfoLine()

		for i, max := 0, len(scotterLineSymbol); i < max; i++ {
			if len(scotterLineSymbol[i]) > 0 {
				infoLine.AddNewLine(scotterLineSymbol[i], scotterLinePoint[i], option)
			} else {
				infoLine.AddEmptyPoint()
			}
		}

		winLineInfo = append(winLineInfo, infoLine)
		otherdata["freegamecount"] = r.FreeGameCount
		otherdata["isfreegame"] = 1

	} else {
		otherdata["isfreegame"] = 0
	}

	result["scores"] = totalScores
	result["gameresult"] = winLineInfo
	if gameType == 1 {
		result = gameplate.ResultMapLine(plateIndex, plateSymbol, winLineInfo)
	} else {
		result = gameplate.ResultMapLine(plateIndex, plateSymbolInsertWild, winLineInfo)
	}
	return result, otherdata, totalScores
}

func (r *Rule) setRandomWild(plateSymbol [][]int, randomWildPoint [][]int) [][]int {

	if len(randomWildPoint) <= 0 {
		return plateSymbol
	}

	var result = make([][]int, len(plateSymbol))

	for cIndex, colSymols := range plateSymbol {
		result[cIndex] = foundation.CopyArray(colSymols)
		for j, jmax := 0, len(randomWildPoint[cIndex]); j < jmax; j++ {
			result[cIndex][randomWildPoint[cIndex][j]] = r.Wild1()
		}
	}
	return result
}

// plateToLinePlate ...
func (r *Rule) plateToLinePlate(plate [][]int, lineMap [][]int) [][]int {
	var plateLineMap [][]int
	var plateline []int

	for _, linePoint := range lineMap {
		plateline = []int{}
		for lineIndex, point := range linePoint {
			plateline = append(plateline, plate[lineIndex][point])
		}
		plateLineMap = append(plateLineMap, plateline)
	}

	return plateLineMap
}

// CutSymbolLink get line link array
func (r *Rule) cutSymbolLink(symbolLine []int, option gameplate.PlateOption) []int {
	var newSymbolLine []int
	mainSymbol := symbolLine[0]

	for _, symbol := range symbolLine {
		if isWild, _ := option.IsWild(symbol); isWild {

		} else if isWild, _ := option.IsWild(mainSymbol); isWild {
			mainSymbol = symbol
		} else if symbol != mainSymbol {
			break
		}

		newSymbolLine = append(newSymbolLine, symbol)
	}

	return newSymbolLine
}

// isWin symbol line compar parline is win
func (r *Rule) isWin(lineSymbol []int, payLineSymbol []int, option gameplate.PlateOption) bool {
	targetSymbol := 0
	isWin := true
	EmptyNum := option.EmptyNum()
	mainSymbol := EmptyNum

	for lineIndex, max := 0, len(payLineSymbol)-1; lineIndex < max; lineIndex++ {
		targetSymbol = lineSymbol[lineIndex]

		if isWild, _ := option.IsWild(targetSymbol); isWild {
			if mainSymbol == EmptyNum {
				mainSymbol = targetSymbol
			}
			continue
		}

		switch payLineSymbol[lineIndex] {
		case targetSymbol:
			mainSymbol = targetSymbol
		default:
			isWin = false
			return isWin
		}
	}

	if mainSymbol != payLineSymbol[0] {
		return false
	}

	return isWin
}

func (r *Rule) lockWild(plater [][]int, lockWild [][]int, option gameplate.PlateOption) [][]int {

	for colIndex, colarray := range plater {
		for rowIndex, row := range colarray {
			if isWild, _ := option.IsWild(row); isWild && !foundation.IsInclude(rowIndex, lockWild[colIndex]) {
				lockWild[colIndex] = append(lockWild[colIndex], rowIndex)
			}
		}
	}

	return lockWild
}
