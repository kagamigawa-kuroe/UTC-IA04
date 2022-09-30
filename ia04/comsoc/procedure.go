package comsoc

import "errors"

func MajoritySWF(p Profile) (count Count, err error){
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	count = make(Count)

	for i:=0; i < len(p); i++ {
		count[p[i][0]]++
	}

	return count, err
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error){
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

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error){
	count = make(Count)
	for i := range p {
		for j := 0; j < thresholds[i]; j++ {
			count[p[i][j]]++
		}
	}
	return count, err
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error){
	count, e :=  ApprovalSWF(p,thresholds)
	if e == nil {
		bestAlts = maxCount(count)
	} else {
		return nil, e
	}
	return bestAlts, nil
}

func TieBreakFactory(a []Alternative) (tiebreak func ([]Alternative) (Alternative, error)) {
	cmp := a
	return func (a []Alternative) (alt Alternative, e error) {
		if len(a) == 0 {
			return -1 ,errors.New("alternative list est Null")
		}
		var note map[Alternative]int
		for i := range cmp{
			note[cmp[i]] = i
		}

		var ans Alternative
		for i := range a{
			if note[a[i]] < note[ans] {
				ans = a[i]
			}
		}

		return ans,nil
	}
}

func SWFFactory(s func (p Profile) (Count, error), t func ([]Alternative)(Alternative, error) ) (swf func (Profile) (Count, error)) {
	return func (p Profile) (Count, error) {
		temp,e:= s(p)
		if e != nil {
			return nil, e
		}
		var new_count Count
		note := make(map[int][]Alternative,10)
		for i,j := range temp {
			note[j] = append(note[j], i)
		}

		for i,j := range note {
			temp,err := t(j)
			if err != nil {
				return nil,err
			}
			new_count[temp] = i
		}

		return new_count,nil
	}
}

func SCFFactory(s func (p Profile) ([]Alternative, error), t func ([]Alternative) (Alternative, error)) (scf func(Profile) (Alternative, error)){
	return func(p Profile) (Alternative, error) {
		temp,e:= s(p)
		if e != nil {
			return -1, e
		}
		a,err := t(temp)
		if err != nil {
			return -1, err
		}
		return a,nil
	}
}

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	count := make(map[Alternative]int)
	note := make([]Alternative,0)
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
					if p[k][t]==note[i] {
						index_1 = t
					}
					if p[k][t]==note[j] {
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

	ans := make([]Alternative,0)
	max_v := -1
	for _,j := range count{
		if j > max_v {
			max_v = j
		}
	}

	for i,j := range count{
		if j == max_v {
			ans = append(ans, i)
		}
	}

	if len(ans) > 1 {
		ans2 := make([]Alternative,0)
		return ans2,nil
	} else {
		return ans, nil
	}
}