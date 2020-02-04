package chapter3

type Eye struct {
	isOpen bool
}

func (e *Eye) Watch() {
	e.isOpen = true
}

type Nose struct {
	isOpen bool
}

func (n *Nose) Breathe() {
	n.isOpen = true
}

type Mouth struct {
	isOpen  bool
	hasFood bool
}

func (m *Mouth) Eat() {
	m.hasFood = true
}

func (m *Mouth) Breathe() {
	m.isOpen = true
}

// 課題4
// 上の3つのstructの機能を持つFaceを実行してください。
// ただし口と鼻両方で呼吸します。
type Face struct{}
