package agt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ia04/comsoc"
	"log"
	"net/http"
	"sync"
	"time"
)

type Bureau struct {
	sync.Mutex
	addr  string
	ID    int
	Name  string
	P    comsoc.Profile
	Alts []comsoc.Alternative
}

func NewBureau(mutex sync.Mutex, addr string, ID int, name string, p comsoc.Profile, alts []comsoc.Alternative) *Bureau {
	return &Bureau{Mutex: mutex, addr: addr, ID: ID, Name: name, P: p, Alts: alts}
}

type Request struct {
	Alts []comsoc.Alternative `json:"prefer"`
}

type Response struct {
	Result []comsoc.Alternative `json:"result"`
}

func (*Bureau) decodeRequest(r *http.Request) (req Request, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (b *Bureau) vote(w http.ResponseWriter, r *http.Request) {
	b.Lock()
	req, err := b.decodeRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	b.P = append(b.P, req.Alts)
	fmt.Println(req)
	b.Unlock()
}

func (b *Bureau) dovote(w http.ResponseWriter, r *http.Request) {
	var resp Response
	b.Lock()
	p,_ := comsoc.MajoritySCF(b.P)
	resp.Result = p

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
	b.Unlock()
}

func (b *Bureau) Start() {

	mux := http.NewServeMux()
	mux.HandleFunc("/vote", b.vote)
	mux.HandleFunc("/dovote", b.dovote)

	// création du serveur http
	s := &http.Server{
		Addr:           b.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on", b.addr)
	go log.Fatal(s.ListenAndServe())
}

///////////////////////////////////////////////////////

type Voteur struct {
	Agent
	url   string
}

func NewVoteur(agent Agent, url string) *Voteur {
	return &Voteur{Agent: agent, url: url}
}

func (v *Voteur) treatResponse(r *http.Response) []comsoc.Alternative {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp Response
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.Result
}

func (v *Voteur) doRequest() (res []comsoc.Alternative, err error) {
	req := Request{
		Alts: v.Prefs,
	}

	url := v.url + "/vote"
	data, _ := json.Marshal(req)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = v.treatResponse(resp)

	return
}

func (v *Voteur) StartVote() {
	log.Printf("démarrage de un Voteur %d",v.ID)

	_, err := v.doRequest()

	if err != nil {
		log.Fatal(v.ID, "error:", err.Error())
	}
	//} else {
	//	log.Printf("the result of vote :")
	//	log.Println(res)
	//}
}

func (v *Voteur) RequestAnswer() []comsoc.Alternative{
	var res Response

	url := v.url + "/dovote"

	data := make([]byte,0)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))


	if err != nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	json.Unmarshal(buf.Bytes(), &res)

	return res.Result
}

