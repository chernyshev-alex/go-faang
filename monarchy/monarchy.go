package monarchy

type Monarchy interface {
	birth(child, parent string)
	death(name string)
	orderOfSuccession() []string
}

type Person struct {
	name     string
	isAlive  bool
	children []*Person
}

func NewPerson(name string) *Person {
	return &Person{
		name:     name,
		isAlive:  true,
		children: make([]*Person, 0),
	}
}

func (p *Person) addChild(child *Person) {
	p.children = append(p.children, child)
}

type MonarchyTree struct {
	king    *Person
	persons map[string]*Person
}

func NewMonarchy(kingName string) MonarchyTree {
	p := NewPerson(kingName)
	return MonarchyTree{
		king:    p,
		persons: map[string]*Person{p.name: p},
	}
}

func (m *MonarchyTree) addPerson(p *Person) {
	m.persons[p.name] = p
}

func (m *MonarchyTree) Birth(child, parent string) {
	if p, ok := m.persons[parent]; ok {
		c := NewPerson(child)
		p.addChild(c)
		m.addPerson(c)
	}
}

func (m *MonarchyTree) Death(name string) {
	if p, ok := m.persons[name]; ok {
		p.isAlive = false
		return
	}
	panic("not found " + name)
}

func (m *MonarchyTree) orderOfSuccession() []string {
	order := make([]string, 0)
	m.dfs(&order, m.king)
	return order
}

func (m *MonarchyTree) dfs(result *[]string, p *Person) {
	if p.isAlive {
		*result = append(*result, p.name)
	}
	for _, c := range p.children {
		m.dfs(result, c)
	}
}
