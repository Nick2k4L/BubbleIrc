package models

type ServerModel struct {
	test []string
}

// These would be some type of box? With highlighting
// once on enter, we will be able to go back. Backspace
// will allow users to go back to the server list
func InitServerModel() *ServerModel {
	return &ServerModel{
		test: []string{"Server1", "Server2", "Server3"},
	}
}

func (m *ServerModel) View() string {
	var final string
	for _, server := range m.test {
		final += server + "\n"
	}

	return final
}
