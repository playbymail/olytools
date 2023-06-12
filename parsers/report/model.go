// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package report

import (
	"fmt"
)

type Report struct {
	Game    *Game
	Faction *Faction
	Season  *Season
	Units   Units
}

func (rpt *Report) String() string {
	if rpt == nil {
		return "{}"
	}
	return fmt.Sprintf("{game: %s, faction: %s}", rpt.Game, rpt.Faction)
}

type Game struct {
	ID            string
	Turn          int
	Faction       *Faction
	Season        *Season
	NoblePoints   *NoblePoints
	FastStudyDays int
}

func (g *Game) String() string {
	if g == nil {
		return "{}"
	}
	return fmt.Sprintf("{id: %q, turn: %d, faction:%s, season:%s, noblePoints: %s, fastStudyDays: %d}", g.ID, g.Turn, g.Faction, g.Season, g.NoblePoints, g.FastStudyDays)
}

type Faction struct {
	ID             string
	Name           string
	UnclaimedItems Inventory
}

func (f *Faction) String() string {
	if f == nil {
		return "{}"
	}
	return fmt.Sprintf("{id: %q, name: %q}", f.ID, f.Name)
}

type Season struct {
	Year   int
	Month  int
	Season string
}

func (s *Season) String() string {
	return fmt.Sprintf("{year: %d, month: %d, season: %q}", s.Year, s.Month, s.Season)
}

type NoblePoints struct {
	Points     int
	Gained     int
	Spent      int
	NextTurn   int
	NextNobles NoblesList
}

type NoblesList []int

func (nl NoblesList) String() string {
	s, sep := "[", ""
	for _, np := range nl {
		s = s + sep + fmt.Sprintf("%d", np)
		sep = ", "
	}
	return s
}

func (np *NoblePoints) String() string {
	if np == nil {
		return "{}"
	}
	return fmt.Sprintf("{points: %d, gained: %d, spent: %d, nextTurn: %d, nextNobles: %s}", np.Points, np.Gained, np.Spent, np.NextTurn, np.NextNobles)
}

type Inventory []*InventoryItem

func (i Inventory) String() string {
	s, sep := "[", ""
	for _, item := range i {
		s = s + sep + item.String()
		sep = ", "
	}
	return s + "]"
}

type InventoryItem struct {
	Code   string
	Name   string
	Qty    int
	Weight int
	Ride   int // this is wrong
}

func (item *InventoryItem) String() string {
	if item == nil {
		return "{}"
	}
	return fmt.Sprintf("{code: %q, name: %q, qty: %d, weight: %d}", item.Code, item.Name, item.Qty, item.Weight)
}

type Units []*Unit

func (u Units) String() string {
	s, sep := "[", ""
	for _, unit := range u {
		s += sep + unit.String()
		sep = ", "
	}
	return s + "]"
}

type Unit struct {
	ID          string
	Name        string
	Location    *Location
	Loyalty     string
	Health      int
	Combat      *Combat
	BreakPoint  int
	SkillsKnown Skills
}

func (u *Unit) String() string {
	if u == nil {
		return "{}"
	}
	return fmt.Sprintf("{id: %q, name: %q, location: %s, loyalty: %q, health: %d, combat: %s, breakPoint: %d}", u.ID, u.Name, u.Location, u.Loyalty, u.Health, u.Combat, u.BreakPoint)
}

type Location struct {
	ID   string
	Name string
	In   *Location
}

func (l *Location) String() string {
	if l == nil {
		return "{}"
	} else if l.In == nil {
		return fmt.Sprintf("{id: %q, name: %q}", l.ID, l.Name)
	}
	return fmt.Sprintf("{id: %q, name: %q, in: %s}", l.ID, l.Name, l.In)
}

type Combat struct {
	Attack  int
	Defense int
	Missile int
	Behind  *CombatBehind
}

func (c *Combat) String() string {
	if c == nil {
		return "{}"
	}
	return fmt.Sprintf("{attack: %d, defense: %d, missile: %d, behind: %s}", c.Attack, c.Defense, c.Missile, c.Behind)
}

type CombatBehind struct {
	ID    int
	Notes string
}

func (cb *CombatBehind) String() string {
	if cb == nil {
		return "{}"
	} else if cb.Notes == "" {
		return fmt.Sprintf("{id: %d}", cb.ID)
	}
	return fmt.Sprintf("{id: %d, notes: %q}", cb.ID, cb.Notes)
}

type Skills []*Skill

func (skills Skills) String() string {
	s, sep := "[", ""
	for _, skill := range skills {
		s += sep + skill.String()
		sep = ", "
	}
	return s + "]"
}

type Skill struct {
	Name  string
	Level int
}

func (s *Skill) String() string {
	if s == nil {
		return "{}"
	}
	return fmt.Sprintf("{name: %q, level: %d}", s.Name, s.Level)
}
