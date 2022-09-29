package comsoc

import (
	"errors"
)

/// 候选人
type Alternative int

/// Profile[i][j] 代表第i位投票人 偏好 j=0 > j=1 > j=2 > ...
type Profile [][]Alternative

/// 候选人得分
type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for i := 0; i < len(prefs); i++ {
		if prefs[i] == alt {
			return i
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2 
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	/// 判断alt1在偏好列表中的下标是否小于alt2
	return rank(alt1,prefs) < rank(alt2,prefs)
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count Count) (bestAlts []Alternative){
	max_value := -1
	/// 一次遍历找到最大值
	for _,j := range count {
		if j > max_value {
			max_value = j
		}
	}

	/// 将等于最大值的Alternative加入slice中
	for i,j := range count {
		if j == max_value {
			bestAlts = append(bestAlts, i)
		}
	}

	return bestAlts
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois par préférences
func checkProfile(prefs Profile) error {

	if len(prefs) == 0 {
		return errors.New("Le profile list est NULL!")
	}

	/// 判读所有的长度是否一致
	length := len(prefs[0])

	for i := 0; i < len(prefs); i++ {
		if len(prefs[i]) != length {
			return errors.New("Profils ne sont pas tous complets!")
		}
	}

	/// 对于每个投票者 判断有无重复的成员
	for i := 0; i < len(prefs); i++ {
		set := make(map[Alternative]bool)
		for j := 0; j < length; j++ {
			if set[prefs[i][j]] == false {
				set[prefs[i][j]] = true
			}else {
				return errors.New("Profils ne sont pas tous complets!")
			}
		}
	}

	return nil
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	if len(prefs) == 0 || len(alts) == 0 {
		return errors.New("Le profile list est NULL!")
	}
	
	Number_Alternative := len(alts)

	/// 判断所有候选人是否都只出现了一次
	note := make(map[Alternative]bool)
	for it := range alts {
		if note[alts[it]] == false {
			note[alts[it]] = true
		} else{
			return errors.New("Le Alternative list est NULL!")
		}
	}

	/// 依次判断每次投票结果
	for i := 0; i < len(prefs); i++ {
		/// 如果长度不对 直接返回
		if len(prefs[i]) != Number_Alternative {
			return errors.New("Profils ne sont pas tous complets!")
		}

		/// 记录所有候选人在set中
		set := make(map[Alternative]int)
	    for alt := range alts {
		    set[alts[alt]] = 1
	    }
		/// 出现一个删除一个
		for j := 0; j < len(prefs[i]); j++ {
			set[prefs[i][j]]-- 
		}

		/// 最后set中元素应该全部为0
		for _,j := range set {
			if j != 0 {
				return errors.New("Profils ne sont pas tous complets!");
			}
		}
	}	

	return nil
}