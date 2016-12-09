package ubus

import "github.com/satori/go.uuid"

type Ubus struct {
	q []eventInfo
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

	u.q = append(u.q, eventInfo{
		ID:    id,
		Token: token,
		Cb:    cb,
		Del:   false,
	})

	return func() {
		for i, v := range u.q {
			if v.ID == id {
				u.q = append(u.q[:i], u.q[i+1:]...)
			}
		}
	}
}

func (u *Ubus) Once(token string, cb func(interface{})) {
	id := uuid.NewV4()

	u.q = append(u.q, eventInfo{
		ID:    id,
		Token: token,
		Cb:    cb,
		Del:   true,
	})
}

func (u *Ubus) Emit(token string, info interface{}) {
	for i, v := range u.q {
		if token == v.Token {
			v.Cb(info)

			if v.Del {
				u.q = append(u.q[:i], u.q[i+1:]...)
			}
		}
	}
}

func NewBus() *Ubus {
	return &Ubus{
		q: []eventInfo{},
	}
}
