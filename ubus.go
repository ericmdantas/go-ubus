package ubus

type Ubus struct {
	q []eventInfo
}

type eventInfo struct {
	ID    string
	Token string
	Cb    func(interface{})
	Del   bool
}

type destroyCb func()

func (u *Ubus) On(token string, cb func(interface{})) destroyCb {
	id := "1"

	u.q = append(u.q, eventInfo{
		ID:    id,
		Token: token,
		Cb:    cb,
		Del:   false,
	})

	return func() {
		for _, v := range u.q {
			if v.ID == id {
				// todo
			}
		}
	}
}

func (u *Ubus) Once(token string, cb func(interface{})) {
	id := "1"

	u.q = append(u.q, eventInfo{
		ID:    id,
		Token: token,
		Cb:    cb,
		Del:   true,
	})
}

func (u *Ubus) Emit(token string, info interface{}) {
	for _, v := range u.q {
		if token == v.Token {
			v.Cb(info)
		}
	}
}

func NewUbus() *Ubus {
	return &Ubus{
		q: []eventInfo{},
	}
}
