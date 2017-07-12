package mongo

import mgo "gopkg.in/mgo.v2"

// Model ...
type Model struct {
}

// Mate ...
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

// Indexs ...
func (m *Model) Indexs() (idexs []mgo.Index) {
	return idexs
}
