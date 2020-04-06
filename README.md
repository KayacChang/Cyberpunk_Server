[https://hackmd.io/jgCwoneZSsqh17mOaq_Q5w?both](https://hackmd.io/jgCwoneZSsqh17mOaq_Q5w?both)
---
# 賽博龐克 Cyberpunk
## Client API
### 
### POST /game/init

#### 遊戲初始化

##### Request
```HTTP
POST http://<gameserverIP>:8000/game/init HTTP/2.0
Authorization: <token>
```

##### Respont

```JSON
{
    "data": {
        "betrate":{
            "betrate":[]int,
            "betratedefaultindex":int,
            "betratelinkindex":[]int
        },
        "player":{
            "gametypeid":string,
            "id":string,
            "money":int
        },
    "reel":{
        "freereel":[][]int,
        "normalreel":[][]int
        }
    },
    "error":{
        "ErrorCode":int,
        "Msg":string
    }
}
```

| Name                | Type    |  Note               |
| ------------------- | ------- |:---------------------- |
| betrate             | []int   | 可下注金額             |
| betratedefaultindex | int     | 預設下注顯示(表演用)   |
| betratelinkindex    | []int   | 下注金額快捷鍵(表演用) |
| gametypeid          | string  | 遊戲類型ID             |
| id                  | string  | 玩家遊戲ID             |
| money               | int     | 玩家現有金額           |
| freereel            | [][]int | 免費遊戲輪帶           |
| normalreel          | [][]int | 一般遊戲輪帶           |
| ErrorCode           | int     | 錯誤代碼               |
| Msg                 | string  | 錯誤訊息               |

### POST /game/result
#### 取得遊戲結果
```HTTP
POST http://<gameserverIP>:8000/game/result HTTP/2.0
Authorization: <token>
```

##### Request
```JSON
{
    "gametypeid":string,
    "bet":int,
}
```


| Name       | Type   | Note           |
| ---------- | ------ |:-------------- |
| gametypeid | string | 遊戲類型ID     |
| bet        | int    | 下注金額索引值 |

##### Respont

```JSON
{
	"data":{
		"freeresult":[{
				"gameresult":[]InfoLine,
				"plate":[][]int,
				"plateindex":[][]int
			}
		],
		"freewildbonusrate":int,
		"isfreegame":int,
		"normalresult": {
			"gameresult":[]InfoLine,
			"plate":[][]int,
			"plateindex":[][]int,
			"randwild":[]int
		},
		"playermoney":int,
		"totalwinscore":int
	},
	"error": {
		"ErrorCode":int,
		"Msg":string
	}
}
```

| Name       | Type                   | Note             |
| ---------- | ---------------------- |:---------------- |
| freeresult | msp[string]interface{} | 免費遊戲結果     |
| gameresult | InfoLine               | 遊戲盤面連線資訊 |
| plate      | [][]int                | 遊戲盤面編號     |
| plateindex | [][]int                | 遊戲盤面索引值   |

#### InfoLine
##### 單線連線資訊
```JSON
{
    "ScotterPoint":[][]int
    "WildPoint":[][]int,
    "LineSymbolPoint":[][]int,
    "LineSymbolNum":[][]int,
    "LineWinIndex":int,
    "Score":int,
    "JackPotScore":int,
    "SpecialWinRate":int,
    "LineWinRate":int
}
```


| Name            | Type    | Note           |
| --------------- | ------- | -------------- |
| ScotterPoint    | [][]int | Scotter位置    |
| WildPoint       | [][]int | Wild位置       |
| LineSymbolPoint | [][]int | 連線Symbol位置 |
| LineSymbolNum   | [][]int | 連線Symbol編號 |
| LineWinIndex    | int     | 連線編號       |
| Score           | int     | 得分           |
| JackPotScore    | int     | 未使用         |
| SpecialWinRate  | int     | 未使用         |
| LineWinRate     | int     | 連線賠率       |
