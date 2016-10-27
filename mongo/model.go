package mongo

type Model struct {
}

func (m *Model) Mate() map[string]string {
	return make(map[string]string, 0)
}

func (m *Model) checkMate() error {
	dbMate := m.Mate()
	if dbMate["dbAddr"] == "" || dbMate["dbName"] == "" || dbMate["cName"] == "" {
		panic("check model mate ")
	}
	return nil
}

func (m *Model) IndexKeys() []string {
	var keys []string
	return keys
}
