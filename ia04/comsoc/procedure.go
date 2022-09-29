package comsoc

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