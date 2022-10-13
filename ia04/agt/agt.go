package agt

import (
	"bytes"
	"fmt"
	"ia04/comsoc"
	"reflect"
)

type AgentID int

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []comsoc.Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a comsoc.Alternative, b comsoc.Alternative)
	Start()
}

func (a2 *Agent) Prefers(a comsoc.Alternative, b comsoc.Alternative) {
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
	fmt.Println("start")
	// panic("implement me")
}

func NewAgent(ID AgentID, name string, prefs []comsoc.Alternative) *Agent {
	return &Agent{ID: ID, Name: name, Prefs: prefs}
}

func (a2 *Agent) Equal(ag AgentI) bool {
	return a2 ==ag
}

func (a2 *Agent) DeepEqual(ag AgentI) bool {
	return reflect.DeepEqual(a2,ag)
}

func (a2 *Agent) Clone() AgentI {
	re :=  NewAgent(a2.ID, a2.Name, a2.Prefs)
	return re
}

func (a2 *Agent) String() string {
	s := "Agent:"
	buf := bytes.NewBufferString(s)
	_, _ = fmt.Fprint(buf, a2.ID)
	_, _ = fmt.Fprint(buf,", ID:" + a2.Name + ", ")
	_, _ = fmt.Fprint(buf,"has preference: " )
	if len(a2.Prefs) == 0{
		_, _ = fmt.Fprint(buf,"NULL" )
		return buf.String()
	}
	_, _ = fmt.Fprint(buf, a2.Prefs[0])

	for i := 1; i < len(a2.Prefs); i++ {
		_, _ = fmt.Fprint(buf," < ")
		_, _ = fmt.Fprint(buf, a2.Prefs[i])
	}

	return buf.String()
}

