// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package report

import (
	"fmt"
	"strconv"
	"strings"
)

func (t tokens) expectBracketedCode() (rest tokens, err error, code string) {
	if len(t) == 0 {
		return t, fmt.Errorf("bracketedCode: want code: got eoi"), ""
	} else if t[0].value == "\n" {
		return t, fmt.Errorf("bracketedCode: want code: got eol"), ""
	}
	if !(strings.HasPrefix(t[0].value, `[`) && strings.HasSuffix(t[0].value, `]`)) {
		return t, fmt.Errorf("bracketedCode: want code: got %q", t[0].value), ""
	}
	return t[1:], nil, strings.Trim(t[0].value, "[]")
}

func (t tokens) expectBreakPoint() (rest tokens, err error, breakPoint int) {
	if rest, err = t.expectWords("Break", "point:"); err != nil {
		return rest, fmt.Errorf("breakPoint: %w", err), breakPoint
	}
	if rest, err, breakPoint = rest.expectPercentage(); err != nil {
		return rest, fmt.Errorf("breakPoint: %w", err), breakPoint
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("breakPoint: %w", err), breakPoint
	}
	return rest, nil, breakPoint
}

func (t tokens) expectCombat() (rest tokens, err error, combat *Combat) {
	combat = &Combat{}
	if rest, err = t.expectLiteral("Combat:"); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err = rest.expectLiteral("attack"); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err, combat.Attack = rest.expectNumber(); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err = rest.expectLiteral(","); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err = rest.expectLiteral("defense"); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err, combat.Defense = rest.expectNumber(); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err = rest.expectLiteral(","); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err = rest.expectLiteral("missile"); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err, combat.Missile = rest.expectNumber(); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	combat.Behind = &CombatBehind{}
	if rest, err = rest.expectLiteral("behind"); err != nil {
		return rest, fmt.Errorf("combat: behind: %w", err), combat
	}
	if rest, err, combat.Behind.ID = rest.expectNumber(); err != nil {
		return rest, fmt.Errorf("combat: behind: %w", err), combat
	}
	if rest, err = rest.expectWords("(front", "line", "in", "combat)"); err != nil {
		return rest, fmt.Errorf("combat: behind: %w", err), combat
	} else {
		combat.Behind.Notes = "front line in combat"
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("combat: %w", err), combat
	}
	return rest, nil, combat
}

func (t tokens) expectEndOfInput() (rest tokens, err error) {
	if len(t) == 0 {
		return t, nil
	}
	return t, fmt.Errorf("%d: want eoi: got %q", t[0].line, t[0].value)
}

func (t tokens) expectEndOfLine() (rest tokens, err error) {
	if len(t) == 0 {
		return t, fmt.Errorf("eol: want eol: got: end-of-input")
	} else if t[0].value != "\n" {
		return t, fmt.Errorf("%d: want eol: got %q", t[0].line, t[0].value)
	}
	return t[1:], nil
}

func (t tokens) expectFaction(id, name string) (rest tokens, err error, faction *Faction) {
	faction = &Faction{}
	if rest, err = t.expectLiteral(name); err != nil {
		return rest, fmt.Errorf("faction: %w", err), faction
	} else {
		faction.Name = name
	}
	if rest, err = rest.expectLiteral(fmt.Sprintf("[%s]", id)); err != nil {
		return rest, fmt.Errorf("faction: %w", err), faction
	} else {
		faction.ID = id
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("faction: %w", err), faction
	}
	if rest, err = rest.expectLiteral("------------------------------------------------------------------------"); err != nil {
		return rest, fmt.Errorf("faction: %w", err), faction
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("faction: %w", err), faction
	} else {
		for rest.acceptEndOfLine() { // skip newlines
			if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
				panic(err)
			}
		}
	}
	if rest, err, faction.UnclaimedItems = rest.expectUnclaimedItems(); err != nil {
		return rest, fmt.Errorf("faction: %w", err), faction
	}

	return rest, nil, faction
}

func (t tokens) expectFastStudyDays() (rest tokens, err error, fsd int) {
	if rest, err, fsd = t.expectInteger(); err != nil {
		return rest, fmt.Errorf("fastStudyDays: %w", err), 0
	}
	if rest, err = rest.expectWords("fast", "study", "days", "are", "left.", "\n"); err != nil {
		return rest, fmt.Errorf("fastStudyDays: %w", err), 0
	}
	return rest, err, fsd
}

func (t tokens) expectGame() (rest tokens, err error, g *Game) {
	g = &Game{}
	if rest, err = t.expectLiteral("Olympia"); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	if rest, err, g.ID = rest.expectID(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	if rest, err = rest.expectLiteral("turn"); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	if rest, err, g.Turn = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	if rest, err, g.Faction = rest.expectInitialPositionReportFor(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	if rest, err, g.Season = rest.expectSeason(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	for rest.acceptEndOfLine() { // skip newlines
		if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
			panic(err)
		}
	}
	if rest, err = rest.expectWelcomeMessage(g); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	for rest.acceptEndOfLine() { // skip newlines
		if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
			panic(err)
		}
	}
	if rest, err = rest.expectNextTurn(g); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	for rest.acceptEndOfLine() { // skip newlines
		if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
			panic(err)
		}
	}
	if rest, err, g.NoblePoints = rest.expectNoblePoints(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}
	for rest.acceptEndOfLine() { // skip newlines
		if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
			panic(err)
		}
	}
	if rest, err, g.FastStudyDays = rest.expectFastStudyDays(); err != nil {
		return rest, fmt.Errorf("game: %w", err), g
	}

	return rest, nil, g
}

func (t tokens) expectHealth() (rest tokens, err error, health int) {
	if rest, err = t.expectLiteral("Health:"); err != nil {
		return rest, fmt.Errorf("health: %w", err), health
	}
	if rest, err, health = rest.expectPercentage(); err != nil {
		return rest, fmt.Errorf("health: %w", err), health
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("health: %w", err), health
	}
	return rest, nil, health
}

func (t tokens) expectID() (rest tokens, err error, id string) {
	if len(t) == 0 {
		return t, fmt.Errorf("id: want id: got end-of-input"), ""
	}
	if t[0].value == "\n" {
		return t, fmt.Errorf("id: want id: got end-of-line"), ""
	}
	if strings.HasPrefix(t[0].value, "\"") {
		return t, fmt.Errorf("id: want id: got %q", t[0].value), ""
	}
	return t[1:], nil, t[0].value
}

func (t tokens) expectInitialPositionReportFor() (rest tokens, err error, faction *Faction) {
	faction = &Faction{}

	if rest, err = t.expectLiteral("Initial"); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	}
	if rest, err = rest.expectLiteral("Position"); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	}
	if rest, err = rest.expectLiteral("Report"); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	}
	if rest, err = rest.expectLiteral("for"); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	}
	if rest, err, faction.Name = rest.expectName(); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	}
	if rest, err, faction.ID = rest.expectID(); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	} else {
		// it should have brackets. let's remove them.
		faction.ID = strings.Trim(faction.ID, `[]`)
		// it may have quote marks. let's remove them.
		faction.ID = strings.Trim(faction.ID, `"`)
	}
	if rest, err = rest.expectWords(".", "\n"); err != nil {
		return rest, fmt.Errorf("initialPositionReportFor: %w", err), faction
	}
	return rest, nil, faction
}

func (t tokens) expectInteger() (rest tokens, err error, i int) {
	if len(t) == 0 {
		return t, fmt.Errorf("integer: want integer: got end-of-input"), 0
	}
	if t[0].value == "\n" {
		return t, fmt.Errorf("integer: want integer: got end-of-line"), 0
	}
	if i, err = strconv.Atoi(t[0].value); err != nil {
		return t, fmt.Errorf("integer: want integer: got %q", t[0].value), 0
	}
	return t[1:], nil, i
}

func (t tokens) expectInventoryItem() (rest tokens, err error, item *InventoryItem) {
	item = &InventoryItem{}
	if rest, err, item.Qty = t.expectNumber(); err != nil {
		return rest, fmt.Errorf("inventoryItem: qty: %w", err), item
	}
	if rest, err, item.Name = rest.expectText(); err != nil {
		return rest, fmt.Errorf("inventoryItem: name: %w", err), item
	} else if item.Name == "riding" { // ugh. item names are not quoted
		if rest, err = rest.expectLiteral("horses"); err != nil {
			return rest, fmt.Errorf("inventory: name: riding horses: %w", err), item
		}
		item.Name = "riding horses"
	}
	if rest, err, item.Code = rest.expectBracketedCode(); err != nil {
		return rest, fmt.Errorf("inventoryItem: code: %w", err), item
	}
	if rest, err, item.Weight = rest.expectNumber(); err != nil {
		return rest, fmt.Errorf("inventoryItem: weight: %w", err), item
	}
	if rest.acceptLiteral("ride") {
		if rest, err = rest.expectLiteral("ride"); err != nil { // should never happen
			panic(err)
		}
		if rest, err, item.Ride = rest.expectInteger(); err != nil {
			return rest, fmt.Errorf("inventoryItem: ride: %w", err), item
		}
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("inventoryItem: eol: %w", err), item
	}
	return rest, nil, item
}

func (t tokens) expectLiteral(literal string) (tokens, error) {
	if len(t) == 0 {
		return t, fmt.Errorf("literal: want %q: got: end-of-input", literal)
	} else if t[0].value != literal {
		return t, fmt.Errorf("literal: %d: want %q: got %q", t[0].line, literal, t[0].value)
	}
	return t[1:], nil
}

func (t tokens) expectLocation() (rest tokens, err error, location *Location) {
	location = &Location{}
	if rest, err = t.expectLiteral("Location:"); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	if rest, err, location.Name = rest.expectName(); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	if rest, err, location.ID = rest.expectID(); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	if rest, err = rest.expectWords(",", "in", "province"); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	location.In = &Location{}
	if rest, err, location.In.Name = rest.expectName(); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	if rest, err, location.In.ID = rest.expectID(); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	if rest, err = rest.expectWords(",", "in"); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	location.In.In = &Location{}
	if rest, err, location.In.In.Name = rest.expectName(); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("location: %w", err), location
	}
	return rest, nil, location
}

func (t tokens) expectLoyalty() (rest tokens, err error, loyalty string) {
	if rest, err = t.expectLiteral("Loyalty:"); err != nil {
		return rest, fmt.Errorf("loyalty: %w", err), loyalty
	}
	if rest, err, loyalty = rest.expectText(); err != nil {
		return rest, fmt.Errorf("loyalty: %w", err), loyalty
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("loyalty: %w", err), loyalty
	}
	return rest, nil, loyalty
}

func (t tokens) expectName() (rest tokens, err error, name string) {
	if len(t) == 0 {
		return t, fmt.Errorf("name: want name: got end-of-input"), ""
	} else if t[0].value == "\n" {
		return t, fmt.Errorf("name: %d: want name: got end-of-line", t[0].line), ""
	}
	return t[1:], nil, t[0].value
}

func (t tokens) expectNextTurn(game *Game) (rest tokens, err error) {
	if rest, err = t.expectWords("The", "next", "turn", "will", "be", "turn"); err != nil {
		return t, fmt.Errorf("nextTurn: %w", err)
	}
	var nextTurn int
	if rest, err, nextTurn = rest.expectInteger(); err != nil {
		return t, fmt.Errorf("nextTurn: %w", err)
	} else if nextTurn != game.Turn+1 {
		return t, fmt.Errorf("nextTurn: %w", fmt.Errorf("expect turn %d: got %d", game.Turn+1, nextTurn))
	}
	if rest, err = rest.expectLiteral("."); err != nil {
		return t, fmt.Errorf("nextTurn: %w", err)
	}
	return rest, nil
}

func (t tokens) expectNoblePoints() (rest tokens, err error, np *NoblePoints) {
	np = &NoblePoints{}
	if rest, err = t.expectWords("Noble", "points:"); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err, np.Points = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err = rest.expectLiteral("("); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err, np.Gained = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err = rest.expectLiteral("gained,"); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err, np.Spent = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err = rest.expectWords("spent)", "\n",
		"The", "next", "NP", "will", "be", "received", "at", "the", "end", "of", "turn"); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err, np.NextTurn = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	if rest, err = rest.expectWords(".", "\n", "\n", "The", "next", "five", "nobles", "formed", "will", "be:"); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}
	var nextNobleNo int // one or more noble numbers
	if rest, err, nextNobleNo = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	} else {
		np.NextNobles = append(np.NextNobles, nextNobleNo)
		for rest.acceptInteger() { // zero or more integers
			if rest, err, nextNobleNo = rest.expectInteger(); err != nil { // this should never happen, so please panic!
				panic(err)
			}
			np.NextNobles = append(np.NextNobles, nextNobleNo)
		}
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("noblePoints: %w", err), np
	}

	return rest, nil, np
}

// number is an integer, optionally with commas
func (t tokens) expectNumber() (rest tokens, err error, i int) {
	if len(t) == 0 {
		return t, fmt.Errorf("number: want number: got end-of-input"), 0
	}
	if t[0].value == "\n" {
		return t, fmt.Errorf("number: want number: got end-of-line"), 0
	}
	if i, err = strconv.Atoi(strings.ReplaceAll(t[0].value, ",", "")); err != nil {
		return t, fmt.Errorf("number: want number: got %q", t[0].value), 0
	}
	return t[1:], nil, i
}

func (t tokens) expectPercentage() (rest tokens, err error, percentage int) {
	if rest, err, percentage = t.expectInteger(); err != nil {
		return rest, fmt.Errorf("percentage: %w", err), percentage
	}
	if rest, err = rest.expectLiteral("%"); err != nil {
		return rest, fmt.Errorf("percentage: %w", err), percentage
	}
	return rest, nil, percentage
}

func (t tokens) expectQuotedText() (rest tokens, err error, text string) {
	if len(t) == 0 {
		return t, fmt.Errorf(`quotedText: want "...": got end-of-input`), ""
	}
	if t[0].value == "\n" {
		return t, fmt.Errorf(`"quotedText: want "...": got end-of-line`), ""
	}
	if !strings.HasPrefix(t[0].value, `"`) || !strings.HasSuffix(t[0].value, `"`) {
		return t, fmt.Errorf(`quotedText: want "...": got %q`, t[0].value), t[0].value
	}
	// let's remove the quote marks from the text
	text = strings.Trim(t[0].value, `""`)
	return t[1:], nil, text
}

func (t tokens) expectReport() (rest tokens, err error, rpt *Report) {
	rpt = &Report{}

	if rest, err, rpt.Game = t.expectGame(); err != nil {
		return rest, fmt.Errorf("report: %w", err), rpt
	} else { // skip newlines at the end of the section
		for rest.acceptEndOfLine() {
			if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
				panic(err)
			}
		}
	}
	if rest, err, rpt.Faction = rest.expectFaction(rpt.Game.Faction.ID, rpt.Game.Faction.Name); err != nil {
		return rest, fmt.Errorf("report: %w", err), rpt
	} else { // skip newlines at the end of the section
		for rest.acceptEndOfLine() {
			if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
				panic(err)
			}
		}
	}
	if rest, err, rpt.Units = rest.expectUnits(); err != nil {
		return rest, fmt.Errorf("report: %w", err), rpt
	} else { // skip newlines at the end of the section
		for rest.acceptEndOfLine() {
			if rest, err = rest.expectEndOfLine(); err != nil { // this should never happen, so please panic!
				panic(err)
			}
		}
	}
	if rest, err = rest.expectEndOfInput(); err != nil {
		return rest, fmt.Errorf("report: %w", err), rpt
	}

	return rest, nil, rpt
}

func (t tokens) expectSeason() (rest tokens, err error, season *Season) {
	season = &Season{}
	if rest, err = t.expectLiteral("Season"); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err, season.Season = rest.expectQuotedText(); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral(","); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral("month"); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err, season.Month = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral(","); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral("in"); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral("the"); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral("year"); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err, season.Year = rest.expectInteger(); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectLiteral("."); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("season: %w", err), season
	}
	return rest, nil, season
}

func (t tokens) expectSkillsKnown() (rest tokens, err error, skills Skills) {
	if rest, err = t.expectWords("Skills", "known:", "\n"); err != nil {
		return rest, fmt.Errorf("skills: %w", err), skills
	}
	if rest, err = rest.expectLiteral("none"); err != nil {
		return rest, fmt.Errorf("skills: %w", err), skills
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("skills: %w", err), skills
	}
	return rest, nil, skills
}

func (t tokens) expectText() (rest tokens, err error, text string) {
	if len(t) == 0 {
		return t, fmt.Errorf("text: want text: got end-of-input"), ""
	} else if t[0].value == "\n" {
		return t, fmt.Errorf("text: want text: got end-of-line"), ""
	}
	if strings.HasPrefix(t[0].value, `"`) {
		return t, fmt.Errorf("text: want text: got %q", t[0].value), t[0].value
	}
	return t[1:], nil, t[0].value
}

func (t tokens) expectUnclaimedItems() (rest tokens, err error, unclaimedItems []*InventoryItem) {
	if rest, err = t.expectWords("Unclaimed", "items:", "\n", "\n", "qty", "name", "weight", "\n", "---", "----", "------", "\n"); err != nil {
		return rest, fmt.Errorf("unclaimedItems: %w", err), unclaimedItems
	}
	// zero or more inventory lines ended by eof or eoi.
	for {
		if rest.acceptEndOfLine() || rest.acceptEndOfInput() {
			break
		}
		var item *InventoryItem
		if rest, err, item = rest.expectInventoryItem(); err != nil {
			return rest, fmt.Errorf("unclaimedItems: %w", err), unclaimedItems
		}
		unclaimedItems = append(unclaimedItems, item)
	}
	return rest, nil, unclaimedItems
}

func (t tokens) expectUnit() (rest tokens, err error, unit *Unit) {
	unit = &Unit{}
	if rest, err, unit.Name = t.expectName(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.ID = rest.expectID(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err = rest.expectWords("\n", "------------------------------------------------------------------------", "\n"); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.Location = rest.expectLocation(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.Loyalty = rest.expectLoyalty(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.Health = rest.expectHealth(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.Combat = rest.expectCombat(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.BreakPoint = rest.expectBreakPoint(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err, unit.SkillsKnown = rest.expectSkillsKnown(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	// inventory
	// capacity
	// lcoations
	// routes
	// inner locations
	// weather
	// seen where
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("unit: %w", err), unit
	}
	return rest, nil, unit
}

func (t tokens) expectUnits() (rest tokens, err error, units Units) {
	// zero or more unit sections ended by eof or eoi
	for rest = t; len(rest) != 0; {
		if rest.acceptEndOfLine() || rest.acceptEndOfInput() {
			fmt.Println("units: found eol or eoi")
			break
		}
		var unit *Unit
		if rest, err, unit = rest.expectUnit(); err != nil {
			return rest, fmt.Errorf("units: %w", err), units
		}
		fmt.Println(unit.String())
		units = append(units, unit)
	}
	return rest, fmt.Errorf("units: !implemented"), units
}

func (t tokens) expectWelcomeMessage(game *Game) (rest tokens, err error) {
	if rest, err = t.expectLiteral("Welcome"); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectLiteral("to"); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectLiteral("Olympia"); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectLiteral(game.ID + "!"); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectWords("This", "is", "an", "initial", "position", "report", "for", "your", "new", "faction.", "\n",
		"You", "are", "player"); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectLiteral(game.Faction.ID + ","); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectLiteral(`"` + game.Faction.Name + `"`); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectLiteral("."); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	if rest, err = rest.expectEndOfLine(); err != nil {
		return rest, fmt.Errorf("welcomeMessage: %w", err)
	}
	return rest, nil
}

func (t tokens) expectWords(words ...string) (rest tokens, err error) {
	rest = t
	for _, word := range words {
		if word == "\n" {
			if rest, err = rest.expectEndOfLine(); err != nil {
				return rest, fmt.Errorf("words: %d: %w", rest.LineNo(), err)
			}
		} else if rest, err = rest.expectLiteral(word); err != nil {
			return rest, fmt.Errorf("words: %d: %w", rest.LineNo(), err)
		}
	}
	return rest, nil
}
