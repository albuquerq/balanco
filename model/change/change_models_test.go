package change

import "testing"

func TestEmployee(t *testing.T) {
	ec := Employee()

	ec.SetName("Natan de Souza Albuquerque").SetPhone("87 99937-9443")

	changes := ec.Changes()

	if len(changes) == 0 {
		t.Fail()
	}

	t.Log(changes)

	values, err := ec.ChangesByProps("name", "phone")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(values)
	}

	values, err = ec.ChangesByProps("desconhecida")
	if err != nil {
		t.Log(err, values, len(values))
	} else {
		t.Fail()
	}
}
