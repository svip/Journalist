package chatlog

const (
	NormalMessage = iota
	ToMessage
)

type Person struct {
	IRCName string
	Name    string
}

type Line struct {
	Person   *Person
	Line     string
	Receiver *Person
	Type     int
}

type Chatlog []Line

func (chatlog Chatlog) GetPerson(name string) *Person {
	for _, line := range chatlog {
		if line.Person.IRCName == name {
			return line.Person
		}
	}
	return nil
}

