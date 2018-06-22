package ubus

import "github.com/satori/go.uuid"

type Ubus struct {
	q map[string]eventInfo
}

type eventInfo struct {
	ID    uuid.UUID
	Token string
	Cb    func(interface{})
	Del   bool
}

type destroyCb func()

func (u *Ubus) On(token string, cb func(interface{})) destroyCb {
	id := uuid.NewV4()

	u.q[token] = eventInfo{
		ID:    id,
		Token: token,
		Cb:    cb,
		Del:   false,
	}

	return func() {
		for _, v := range u.q {
			if v.ID == id {
				delete(u.q, token)
			}
		}
	}
}

func (u *Ubus) Once(token string, cb func(interface{})) {
	id := uuid.NewV4()

	u.q[token] = eventInfo{
		ID:    id,
		Token: token,
		Cb:    cb,
		Del:   true,
	}
}

func (u *Ubus) Emit(token string, info interface{}) {
	for _, v := range u.q {
		if token == v.Token {
			v.Cb(info)

			if v.Del {
				delete(u.q, token)
			}
		}
	}
}

func (u *Ubus) Off(tokens []string) {
	for _, token := range tokens {
		for _, ev := range u.q {
			if ev.Token == token {
				delete(u.q, token)
			}
		}
	}
}

func NewBus() *Ubus {
	return &Ubus{
		q: make(map[string]eventInfo),
	}
}
