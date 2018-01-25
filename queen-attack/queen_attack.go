package queenattack

import (
	"errors"
	"math"
	"strconv"
)

type position struct {
	a rune
	p int
}

func isValidPos(pos string) (position, bool) {
	alpha := rune(pos[0])
	p, err := strconv.Atoi(string(pos[1]))
	if err != nil {
		return position{}, false
	}

	if alpha < 'a' || alpha > 'h' || p < 1 || p > 8 {
		return position{}, false
	}

	return position{a: alpha, p: p}, true
}

// CanQueenAttack validates two positions on chess board on queen attack.
func CanQueenAttack(w string, b string) (bool, error) {

	if len(w) != 2 || len(w) != len(b) || w == b {
		return false, errors.New("Invalid Input")
	}

	wpos, aok := isValidPos(w)
	bpos, bok := isValidPos(b)
	if !aok || !bok {
		return false, errors.New("Invalid Input")
	}

	if wpos.a == bpos.a || wpos.p == bpos.p {
		return true, nil
	}

	adiff := math.Abs(float64(wpos.a - bpos.a))
	pdiff := math.Abs(float64(wpos.p - bpos.p))
	if adiff != pdiff {
		return false, nil
	}
	return true, nil
}
