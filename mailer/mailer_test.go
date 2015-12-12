package mailer

import "testing"

func TestParse(t *testing.T) {
	m := map[string]string{"name": "Joao"}
	template := `Ola {{name}}`
	if Parse(m, template) != "Ola Joao" {
		t.Error("parsing failed")
	}
}
