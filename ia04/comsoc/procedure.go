package comsoc

import (
	"errors"
)

func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	count = make(Count)

	for i := 0; i < len(p); i++ {
		count[p[i][0]]++
	}

	return count, err
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, e := MajoritySWF(p)
	if e == nil {
		bestAlts = maxCount(count)
	} else {
		return nil, e
	}
	return bestAlts, nil
}

func BordaSWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}
	count = make(Count)

	for i := range p {
		vote := len(p[0]) - 1
		for j := range p[i] {
			count[p[i][j]] = count[p[i][j]] + vote
			vote--
		}
	}
	return count, err
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, e := BordaSWF(p)
	if e == nil {
		bestAlts = maxCount(count)
		return bestAlts, nil
	} else {
		return nil, e
	}
}

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	count = make(Count)
	for i := range p {
		for j := 0; j < thresholds[i]; j++ {
			count[p[i][j]]++
		}
	}
	return count, err
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, e := ApprovalSWF(p, thresholds)
	if e == nil {
		bestAlts = maxCount(count)
	} else {
		return nil, e
	}
	return bestAlts, nil
}

func TieBreakFactory(a []Alternative) (tiebreak func([]Alternative) (Alternative, error)) {
	cmp := a
	return func(a []Alternative) (alt Alternative, e error) {
		if len(a) == 0 {
			return -1, errors.New("alternative list est Null")
		}
		var note map[Alternative]int
		for i := range cmp {
			note[cmp[i]] = i
		}

		var ans Alternative
		for i := range a {
			if note[a[i]] < note[ans] {
				ans = a[i]
			}
		}

		return ans, nil
	}
}

func SWFFactory(s func(p Profile) (Count, error), t func([]Alternative) (Alternative, error)) (swf func(Profile) (Count, error)) {
	return func(p Profile) (Count, error) {
		temp, e := s(p)
		if e != nil {
			return nil, e
		}
		var new_count Count
		note := make(map[int][]Alternative, 10)
		for i, j := range temp {
			note[j] = append(note[j], i)
		}

		for i, j := range note {
			temp, err := t(j)
			if err != nil {
				return nil, err
			}
			new_count[temp] = i
		}

		return new_count, nil
	}
}

func SCFFactory(s func(p Profile) ([]Alternative, error), t func([]Alternative) (Alternative, error)) (scf func(Profile) (Alternative, error)) {
	return func(p Profile) (Alternative, error) {
		temp, e := s(p)
		if e != nil {
			return -1, e
		}
		a, err := t(temp)
		if err != nil {
			return -1, err
		}
		return a, nil
	}
}

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	count := make(map[Alternative]int)
	note := make([]Alternative, 0)
	for i := range p[0] {
		note = append(note, p[0][i])
	}

	for i := 0; i < len(note); i++ {
		for j := i + 1; j < len(note); j++ {
			/// 比较note[i]和note[j]
			a := 0
			b := 0
			for k := range p {
				index_1 := -1
				index_2 := -1
				for t := range p[k] {
					if p[k][t] == note[i] {
						index_1 = t
					}
					if p[k][t] == note[j] {
						index_2 = t
					}
				}
				if index_1 < index_2 {
					a++
				} else if index_1 > index_2 {
					b++
				}
			}
			if a > b {
				count[note[i]]++
			} else if b > a {
				count[note[j]]++
			}
		}
	}

	ans := make([]Alternative, 0)
	max_v := -1
	for _, j := range count {
		if j > max_v {
			max_v = j
		}
	}

	for i, j := range count {
		if j == max_v {
			ans = append(ans, i)
		}
	}

	if len(ans) > 1 {
		ans2 := make([]Alternative, 0)
		return ans2, nil
	} else {
		return ans, nil
	}
}

/// avance
func CopelandSWF(p Profile) (Count, error) {
	count := make(map[Alternative]int)
	note := make([]Alternative, 0)
	for i := range p[0] {
		note = append(note, p[0][i])
	}

	for i := 0; i < len(note); i++ {
		for j := i + 1; j < len(note); j++ {
			/// 比较note[i]和note[j]
			a := 0
			b := 0
			for k := range p {
				index_1 := -1
				index_2 := -1
				for t := range p[k] {
					if p[k][t] == note[i] {
						index_1 = t
					}
					if p[k][t] == note[j] {
						index_2 = t
					}
				}
				if index_1 < index_2 {
					a++
				} else if index_1 > index_2 {
					b++
				}
			}
			if a > b {
				count[note[i]]++
				/// 比起Condorcet 增加一个减的环节
				count[note[j]]--
			} else if b > a {
				count[note[j]]++
				count[note[i]]--
			}
		}
	}
	return count, nil
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	ans := make([]Alternative, 0)
	max_v := -1
	for _, j := range count {
		if j > max_v {
			max_v = j
		}
	}

	for i, j := range count {
		if j == max_v {
			ans = append(ans, i)
		}
	}
	return ans, nil
}

func STV_SWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}
	count = make(Count)

	/// delete un candidate dans P
	del := func(data *Profile, a Alternative) {
		for i := 0; i < len(*data); i++ {
			j := 0
			for ; (*data)[i][j] != a; j++ {
			}
			k := j + 1
			for ; k < len((*data)[i]); k++ {
				(*data)[i][k-1] = (*data)[i][k]
			}
			(*data)[i][k-1] = -1
		}
	}
	size := len(p[0])

	for i := range p[0] {
		count[p[0][i]] = 1
	}

	for i := 0; i < size-1; i++ {
		var a Alternative = p[0][0]
		m := make(map[Alternative]int)
		for j := 0; j < len(p); j++ {
			m[p[j][0]]++
		}
		for i, j := range m {
			if j < m[a] {
				a = i
			}
		}
		count[a] = -1
		del(&p, a)
	}

	return count, err
}

func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	c, err := STV_SWF(p)
	if err != nil {
		return nil, err
	}
	for i, j := range c {
		if j == 1 {
			bestAlts = append(bestAlts, i)
		}
	}
	return bestAlts, nil
}

/// TD6

func Distance_edit(a1 []Alternative, a2 []Alternative) (ans int, e error) {
	if len(a1) != len(a2) {
		return -1, errors.New("taille de deux slice n'est pas meme")
	}
	ans = 0
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			i++
		}
	}
	return ans, nil
}

func Distance_edit_somme(a1 []Alternative, p Profile) (ans int, e error) {
	ans = 0
	for j := range p {
		if len(a1) != len(p[j]) {
			return -1, errors.New("taille de deux slice n'est pas meme")
		}
		a, _ := Distance_edit(a1, p[j])
		ans += a
	}
	return ans, nil
}

func permute(nums []Alternative) [][]Alternative {
	var ans [][]Alternative
	var dfs func(l []Alternative, temp []Alternative)
	dfs = func(l []Alternative, temp []Alternative) {
		if len(l) == 0 {
			ans = append(ans, temp)
		}
		for i := 0; i < len(l); i++ {
			n := append([]Alternative{}, l...)
			dfs(append(n[:i], n[i+1:]...), append(temp, l[i]))
		}
	}
	dfs(nums, []Alternative{})
	return ans
}

func Kemeny(p Profile) (ans []Alternative, e error) {
	e = checkProfile(p)
	if e != nil {
		return nil, e
	}
	// all alternative
	all := make([]Alternative, 0)
	for i := range p[0] {
		all = append(all, p[0][i])
	}

	ans = nil
	min_distance := int(^uint(0) >> 1)
	all_combination := permute(all)

	for _, j := range all_combination {
		d, err := Distance_edit_somme(j, p)
		if err != nil {
			e = err
			return nil, e
		}
		if d < min_distance {
			ans = j
			min_distance = d
		}
	}

	return ans, nil
}

func possibleWinners_STV(p Profile) map[Alternative] []Alternative {
	ans := make(map[Alternative] []Alternative)
	a := make([]Alternative,0)
	p = append(p,a)
	// all alternative
	all := make([]Alternative, 0)
	for i := range p[0] {
		all = append(all, p[0][i])
	}

	all_combination := permute(all)

	for _, j := range all_combination {
		p[len(p)-1] = j
		re,_ := STV_SCF(p)
		ans[re[0]] = j
	}
	return ans
}

func IsPossibleWinner_STV(P Profile, c Alternative) bool {
	ans := possibleWinners_STV(P)
	if ans[c] != nil {
		return true
	}
	return false
}

func isNecessaryWinner(P Profile, c Alternative) bool {
	ans := possibleWinners_STV(P)
	for i,_ := range ans {
		if i != c {
			return false
		}
	}
	return true
}

