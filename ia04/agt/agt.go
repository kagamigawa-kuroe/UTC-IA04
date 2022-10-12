package agt

import (
	"bytes"
	"fmt"
	"reflect"
)

type Alternative int
type AgentID int

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative)
	Start()
}

func (a2 *Agent) Prefers(a Alternative, b Alternative) {
	/// 先判断a与b是否已经在slice中
	IndexA := -1
	IndexB := -1
	for i := 0; i < len(a2.Prefs); i++ {
		if a2.Prefs[i] == a {
			IndexA = i
		}
		if a2.Prefs[i] == b {
			IndexB = i
		}
	}
	/// 如果都不存在
	if IndexA != -1 && IndexB != -1 {
		if IndexA > IndexB {
			temp := a2.Prefs[IndexA]
			a2.Prefs[IndexA] = a2.Prefs[IndexB]
			a2.Prefs[IndexB] = temp
		}
		/// 都存在
	} else if IndexA == -1 && IndexB == -1 {
		a2.Prefs = append(a2.Prefs, a)
		a2.Prefs = append(a2.Prefs, b)
		/// 如果存在A 不存在B
	} else if IndexB == -1 {
		a2.Prefs = append(a2.Prefs, b)
	} else {
		/// 如果存在B 不存在A
		a2.Prefs[IndexB] = a
		a2.Prefs = append(a2.Prefs, b)
	}
	return
}

func (a2 *Agent) Start() {
	// panic("implement me")
}

func NewAgent(ID AgentID, name string, prefs []Alternative) *Agent {
	return &Agent{ID: ID, Name: name, Prefs: prefs}
}

func (a *Agent) Equal(ag AgentI) bool {
	return a==ag
}

func (a *Agent) DeepEqual(ag AgentI) bool {
	return reflect.DeepEqual(a,ag)
}

func (a *Agent) Clone() AgentI {
	re :=  NewAgent(a.ID,a.Name,a.Prefs)
	return re
}

func (a *Agent) String() string {
	s := "Agent:"
	buf := bytes.NewBufferString(s)
	_, _ = fmt.Fprint(buf, a.ID)
	_, _ = fmt.Fprint(buf,", ID:" + a.Name + ", ")
	_, _ = fmt.Fprint(buf,"has preference: " )
	if len(a.Prefs) == 0{
		_, _ = fmt.Fprint(buf,"NULL" )
		return buf.String()
	}
	_, _ = fmt.Fprint(buf,a.Prefs[0])

	for i := 1; i < len(a.Prefs); i++ {
		_, _ = fmt.Fprint(buf," < ")
		_, _ = fmt.Fprint(buf,a.Prefs[i])
	}

	return buf.String()
}

