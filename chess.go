package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// Race .
type Race struct{}

// Type Race
type RaceOrc Race     // 兽人
type RaceBeast Race   // 野兽
type RaceOrge Race    // 食人魔
type RaceDruid Race   // 德鲁伊
type RaceUndead Race  // 亡灵
type RaceGnome Race   // 地精
type RaceTroll Race   // 巨魔
type RaceElf Race     // 精灵
type RaceHuman Race   // 人类
type RaceNaga Race    // 娜迦
type RaceDemon Race   // 恶魔
type RaceElement Race // 元素
type RaceDwarf Race   // 矮人
type RaceDragon Race  // 龙

// RaceType .
var RaceType = []reflect.Type{
	reflect.TypeOf(RaceOrc{}),     // 兽人
	reflect.TypeOf(RaceBeast{}),   // 野兽
	reflect.TypeOf(RaceOrge{}),    // 食人魔
	reflect.TypeOf(RaceDruid{}),   // 德鲁伊
	reflect.TypeOf(RaceUndead{}),  // 亡灵
	reflect.TypeOf(RaceGnome{}),   // 地精
	reflect.TypeOf(RaceTroll{}),   // 巨魔
	reflect.TypeOf(RaceElf{}),     // 精灵
	reflect.TypeOf(RaceHuman{}),   // 人类
	reflect.TypeOf(RaceNaga{}),    // 娜迦
	reflect.TypeOf(RaceDemon{}),   // 恶魔
	reflect.TypeOf(RaceElement{}), // 元素
	reflect.TypeOf(RaceDwarf{}),   // 矮人
	reflect.TypeOf(RaceDragon{}),  // 龙
}

// RaceMap .
var RaceMap = make(map[string]reflect.Type, len(RaceType))

// Career .
type Career struct{}
type CareerSoldier Career      // 战士
type CareerDruid Career        // 德鲁伊
type CareerMage Career         // 法师
type CareerHunter Career       // 猎人
type CareerAssassinator Career // 刺客
type CareerCraftsman Career    // 工匠
type CareerShaman Career       // 萨满祭司
type CareerKnight Career       // 骑士
type CareerDemonHunter Career  // 恶魔猎手
type CareerWizard Career       // 术士

// CareerType .
var CareerType = []reflect.Type{
	reflect.TypeOf(CareerSoldier{}),
	reflect.TypeOf(CareerDruid{}),
	reflect.TypeOf(CareerMage{}),
	reflect.TypeOf(CareerHunter{}),
	reflect.TypeOf(CareerAssassinator{}),
	reflect.TypeOf(CareerCraftsman{}),
	reflect.TypeOf(CareerShaman{}),
	reflect.TypeOf(CareerKnight{}),
	reflect.TypeOf(CareerDemonHunter{}),
	reflect.TypeOf(CareerWizard{}),
}

// CareerMap .
var CareerMap = make(map[string]reflect.Type, len(CareerType))

// Color .
type Color struct{}
type ColorWhite Color  // 白卡
type ColorCygn Color   // 青卡
type ColorBlue Color   // 蓝卡
type ColorPurple Color // 紫卡
type ColorOrange Color // 橙卡

var ColorType = []reflect.Type{
	reflect.TypeOf(ColorWhite{}),
	reflect.TypeOf(ColorCygn{}),
	reflect.TypeOf(ColorBlue{}),
	reflect.TypeOf(ColorPurple{}),
	reflect.TypeOf(ColorOrange{}),
}
var ColorMap = make(map[string]reflect.Type, len(ColorType))

var ColorAmount = map[reflect.Type]int{
	reflect.TypeOf(ColorWhite{}):  45,
	reflect.TypeOf(ColorCygn{}):   30,
	reflect.TypeOf(ColorBlue{}):   25,
	reflect.TypeOf(ColorPurple{}): 15,
	reflect.TypeOf(ColorOrange{}): 10,
}

// Chess .
type Chess struct {
	name   []string
	career interface{}
	race   []interface{}
	color  interface{}
}

var chesses []Chess

func mapinit(arr []reflect.Type, m map[string]reflect.Type) {
	for _, typ := range arr {
		m[typ.Name()] = typ
	}
}

// ChessesInit .
func ChessesInit(jsonfile string) error {
	mapinit(ColorType, ColorMap)
	mapinit(CareerType, CareerMap)
	mapinit(RaceType, RaceMap)
	file, err := os.Open(jsonfile)
	if err != nil {
		return err
	}
	jdec := json.NewDecoder(file)
	jsonChesses := []struct {
		Name   []string
		Career string
		Race   []string
		Color  string
	}{}
	if err := jdec.Decode(&jsonChesses); err != nil {
		return err
	}
	chesses = make([]Chess, 0, len(jsonChesses))
loop:
	for i, jsonChess := range jsonChesses {
		chess := Chess{}
		career, ok := CareerMap[jsonChess.Career]
		if !ok {
			fmt.Println(i, "unknow career:", career)
			continue loop
		}
		chess.career = reflect.New(career).Elem().Interface()
		chess.race = make([]interface{}, 0, len(jsonChess.Race))
		for _, jsonRace := range jsonChess.Race {
			race, ok := RaceMap[jsonRace]
			if !ok {
				fmt.Println(i, "unknow race:", jsonRace)
				continue loop
			}
			chess.race = append(chess.race, reflect.New(race).Elem().Interface())
		}
		color, ok := ColorMap[jsonChess.Color]
		if !ok {
			fmt.Println(i, "unknow color:", color)
			continue loop
		}
		chess.color = reflect.New(color).Elem().Interface()
		chess.name = jsonChess.Name
		chesses = append(chesses, chess)
	}
	return nil
}
