// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package reports

type Model struct {
	Game              Game          `json:"game"`
	Banner            []string      `json:"banner"`
	NextTurn          int           `json:"next_turn"`
	Noble             *Noble        `json:"noble,omitempty"`
	FastStudyDaysLeft int           `json:"fast_study_days_left"`
	Faction           *Faction      `json:"faction,omitempty"`
	NewPlayers        []*PlayerInfo `json:"new_players,omitempty"`
	OrderTemplate     OrderTemplate `json:"order_template,omitempty"`
}

type Capacity struct {
	NumberTop    int     `json:"number_top"`
	NumberBottom int     `json:"number_bottom"`
	Name         string  `json:"name"`
	Percentage   float64 `json:"percentage"`
}
type City struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	In   *In    `json:"in,omitempty"`
}
type CombatFactors struct {
	Attack     int    `json:"attack"`
	Defense    int    `json:"defense"`
	Missile    int    `json:"missile"`
	Behind     int    `json:"behind"`
	BehindNote string `json:"behind_note"`
}
type ControlledBy struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Kind string `json:"kind"`
	In   *In    `json:"in,omitempty"`
}
type Faction struct {
	Name           string           `json:"name"`
	Id             string           `json:"id"`
	UnclaimedItems []*UnclaimedItem `json:"unclaimed_items"`
	Nobles         []*FactionNoble  `json:"nobles"`
	Locations      []*Location      `json:"locations"`
}
type Game struct {
	Name          string   `json:"name"`
	Turn          int      `json:"turn"`
	InitialReport bool     `json:"initial_report"`
	For           *Faction `json:"for"` // TODO: should out just name and id
	Season        string   `json:"season"`
	Month         int      `json:"month"`
	Year          int      `json:"year"`
}
type In struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Kind string `json:"kind,omitempty"`
}
type InnerLocation struct {
	Name           string           `json:"name"`
	Id             string           `json:"id"`
	Descr          string           `json:"descr,omitempty"`
	SafeHaven      bool             `json:"safe_haven,omitempty"`
	TravelTime     *TravelTime      `json:"travel_time,omitempty"`
	Kind           string           `json:"kind,omitempty"`
	InnerLocations []*InnerLocation `json:"inner_locations,omitempty"` // todo: should be name, id, kind, defense
	Defense        int              `json:"defense,omitempty"`
}
type Inventory struct {
	Qty    int               `json:"qty"`
	Name   string            `json:"name"`
	Code   string            `json:"code"`
	Weight int               `json:"weight"`
	Cap    InventoryCapacity `json:"cap,omitempty"`
}
type InventoryCapacity struct {
	Amount int    `json:"amount"`
	Note   string `json:"note"`
}
type Location struct {
	Name                    string           `json:"name"`
	Id                      string           `json:"id"`
	Kind                    string           `json:"kind"`
	In                      *In              `json:"in,omitempty"`
	SafeHaven               bool             `json:"safe_haven"`
	Rating                  string           `json:"rating,omitempty"`
	Province                *Province        `json:"province,omitempty"`
	RuledBy                 *Ruler           `json:"ruled_by,omitempty"`
	Routes                  *Route           `json:"routes,omitempty"`
	InnerLocations          []*InnerLocation `json:"inner_locations"`
	Weather                 Weather          `json:"weather,omitempty"`
	SeenHere                []*Seen          `json:"seen_here,omitempty"`
	CitiesRumoredToBeNearby []*City          `json:"cities_rumored_to_be_nearby,omitempty"`
	SkillsTaughtHere        []*Skill         `json:"skills_taught_here,omitempty"`
	MarketReport            []*MarketReport  `json:"market_report,omitempty"`
}
type NobleLocation struct {
	City     *NobleCity     `json:"city"`
	Province *NobleProvince `json:"province,omitempty"`
	Region   string         `json:"region"`
}
type NobleCity struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
type NobleProvince struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
type MarketReport struct {
	Trade     string `json:"trade"`
	Who       string `json:"who"`
	Price     int    `json:"price"`
	Qty       int    `json:"qty"`
	WeightPer int    `json:"weight_per"`
	Name      string `json:"name"`
	Code      string `json:"code"`
}
type Noble struct {
	Points         int      `json:"points"`
	Gained         int      `json:"gained"`
	Spent          int      `json:"spent"`
	Notes          []string `json:"notes"`
	IncomingNobles []int    `json:"incoming_nobles"`
}
type FactionNoble struct {
	Name       string         `json:"name"`
	Id         string         `json:"id"`
	Location   *NobleLocation `json:"location"`
	Loyalty    string         `json:"loyalty"`
	Health     float64        `json:"health"`
	Combat     *CombatFactors `json:"combat,omitempty"`
	BreakPoint float64        `json:"break_point"`
	Skills     []interface{}  `json:"skills"`
	Inventory  []*Inventory   `json:"inventory"`
	Capacity   *Capacity      `json:"capacity"`
}
type OrderTemplate []string
type PlayerInfo struct {
	Id      string `json:"id"`
	Faction string `json:"faction"`
	Email   string `json:"email"`
}
type Province struct {
	ControlledBy *ControlledBy `json:"controlled_by"`
}
type Route struct {
	Leaving []*RouteLeaving `json:"leaving,omitempty"`
}
type RouteLeaving struct {
	Direction string   `json:"direction"`
	To        *RouteTo `json:"to,omitempty"`
	Kind      string   `json:"kind,omitempty"`
}
type RouteTo struct {
	Name       string      `json:"name"`
	Id         string      `json:"id"`
	OtherName  string      `json:"other-name,omitempty"`
	Impassable bool        `json:"impassable,omitempty"`
	TravelTime *TravelTime `json:"travel_time,omitempty"`
}
type Ruler struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Kind string `json:"kind"`
}
type Seen struct {
	Name   string  `json:"name"`
	Id     string  `json:"id"`
	Kind   string  `json:"kind,omitempty"`
	On     string  `json:"on,omitempty"`
	With   []*With `json:"with,omitempty"`
	Splat  bool    `json:"splat,omitempty"`
	Undead bool    `json:"undead,omitempty"`
	Note   string  `json:"note,omitempty"`
}
type Skill struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
type TravelTime struct {
	Amount int    `json:"amount"`
	Units  string `json:"units"`
}
type UnclaimedItem struct {
	Qty    int    `json:"qty"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Weight int    `json:"weight"`
	Note   string `json:"note,omitempty"`
}
type Weather string
type With struct {
	Amount int    `json:"amount"`
	Units  string `json:"units,omitempty"`
	Unit   string `json:"unit,omitempty"`
}
