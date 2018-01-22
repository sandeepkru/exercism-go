package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Table is simple struct to track scores table.
type Table struct {
	name           string
	mp, w, d, l, p int
}

const (
	win  = "win"
	loss = "loss"
	draw = "draw"
)

// Sort interface methods.
type records []*Table

func (r records) Len() int {
	return len(r)
}

func (r records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r records) Less(i, j int) bool {
	if r[i].p == r[j].p {
		return r[i].name < r[j].name
	}
	return r[i].p > r[j].p
}

// Tally function returns table of scores.
func Tally(input io.Reader, output io.Writer) error {
	var lines []string
	var teamMap = make(map[string]*Table)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	split := func(c rune) bool {
		return c == ';'
	}

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") { // ignore invalid lines
			continue
		}
		words := strings.FieldsFunc(line, split)
		if len(words) != 3 || (words[2] != win && words[2] != loss && words[2] != draw) {
			return errors.New("Invalid input")
		}

		team := words[0]
		opponent := words[1]
		result := words[2]
		_, ok := teamMap[team]
		if !ok {
			teamMap[team] = &Table{name: team}
		}

		_, ok = teamMap[opponent]
		if !ok {
			teamMap[opponent] = &Table{name: opponent}
		}

		team1, _ := teamMap[team]
		team2, _ := teamMap[opponent]

		team1.mp++
		team2.mp++

		switch result {
		case win:
			{
				team1.w++
				team1.p += 3
				team2.l++
			}
		case loss:
			{
				team2.w++
				team2.p += 3
				team1.l++
			}
		case draw:
			{
				team1.d++
				team2.d++
				team1.p++
				team2.p++
			}
		}

		teamMap[team] = team1
		teamMap[opponent] = team2
	}

	l := make(records, 0, len(teamMap))

	for _, rec := range teamMap {
		l = append(l, rec)
	}

	sort.Sort(l)

	io.WriteString(output, fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, r := range l {
		io.WriteString(output, fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", r.name, r.mp, r.w, r.d, r.l, r.p))
	}

	return nil
}
