package mongo

type Model struct {
}

func (m *Model) dbAddr() string {
	return ""
}

func (m *Model) dbName() string {
	return ""
}

func (m *Model) cName() string {
	return ""
}

func (m *Model) indexKeys() []string {
	var keys []string
	return keys
}
