package confield

import (
	"os"
	"testing"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type testStruct struct {
	Field F
}

func TestIsSet(t *testing.T) {
	tests := map[string]bool{
		"field: foo":        true,
		"field: ":           false,
		"":                  false,
		"field: $EMPTY":     false,
		"field: $EMPTY|foo": true,
	}
	for yml, expect := range tests {
		s := testStruct{}
		yaml.Unmarshal([]byte(yml), &s)
		if s.Field.IsSet() != expect {
			t.Fatalf("expect %t, get %t", expect, s.Field.IsSet())
		}
	}
}

func TestStringField(t *testing.T) {
	tests := map[string]string{
		"field: foo":              "foo",
		"field: ":                 "",
		"field: $EMPTY":           "",
		"field: $EMPTY|default":   "default",
		`field: "$EMPTY|default"`: "default",
	}
	for yml, expect := range tests {
		s := testStruct{}
		yaml.Unmarshal([]byte(yml), &s)
		if s.Field.String() != expect {
			t.Fatalf("expect %s, get %s", expect, s.Field.String())
		}
	}
}

func TestBoolField(t *testing.T) {
	tests := map[string]bool{
		"field: true":         true,
		"field: True":         true,
		"field: false":        false,
		"field: False":        false,
		"field: invalid":      false,
		"field: ":             false,
		"field: $EMPTY":       false,
		"field: $EMPTY|false": false,
	}
	for yml, expect := range tests {
		s := testStruct{}
		yaml.Unmarshal([]byte(yml), &s)
		if s.Field.Bool() != expect {
			t.Fatalf("expect %t, get %t", expect, s.Field.Bool())
		}
	}
}

func TestIntField(t *testing.T) {
	tests := map[string]int{
		"field: 0":         0,
		"field: 1":         1,
		"field: 100":       100,
		"field: -1":        -1,
		"field: ":          0,
		"field: notint":    0,
		"field: $EMPTY|1":  1,
		"field: $EMPTY|-1": -1,
		"field: $EMPTY|0":  0,
		"field: $EMPTY|":   0,
	}
	for yml, expect := range tests {
		s := testStruct{}
		yaml.Unmarshal([]byte(yml), &s)
		if s.Field.Int() != expect {
			t.Fatalf("expect %d, get %d", expect, s.Field.Int())
		}
	}
}

func TestFloat64Field(t *testing.T) {
	tests := map[string]float64{
		"field: 0":         0,
		"field: 1":         1,
		"field: 100":       100,
		"field: -1":        -1,
		"field: 1.23":      1.23,
		"field: -1.23":     -1.23,
		"field: ":          0,
		"field: notfloat":  0,
		"field: $EMPTY|1":  1,
		"field: $EMPTY|-1": -1,
		"field: $EMPTY|0":  0,
		"field: $EMPTY|":   0,
	}
	for yml, expect := range tests {
		s := testStruct{}
		yaml.Unmarshal([]byte(yml), &s)
		if s.Field.Float64() != expect {
			t.Fatalf("expect %f, get %f", expect, s.Field.Float64())
		}
	}
}

func TestTimeField(t *testing.T) {
	tests := map[string]time.Time{
		"field: 2017-01-02":          time.Date(2017, 1, 2, 0, 0, 0, 0, &time.Location{}),
		"field: 2017-01-02 12:23:34": time.Date(2017, 1, 2, 12, 23, 34, 0, &time.Location{}),
		"field: ":                    time.Time{},
		"field: nottime":             time.Time{},
		"field: $EMPTY":              time.Time{},
		"field: $EMPTY|2017-01-02":   time.Date(2017, 1, 2, 0, 0, 0, 0, &time.Location{}),
	}
	for yml, expect := range tests {
		s := testStruct{}
		yaml.Unmarshal([]byte(yml), &s)
		if !s.Field.Time().Equal(expect) {
			t.Fatalf("expect %s, get %s", expect, s.Field.Time())
		}
	}
}

func TestEnvVar(t *testing.T) {
	yml := "field: $CONFIELD_TEST|default"
	s := testStruct{}
	yaml.Unmarshal([]byte(yml), &s)

	if s.Field.String() != "default" {
		t.Fatalf("expect %s, get %s", s.Field.String(), "default")
	}

	os.Setenv("CONFIELD_TEST", "envvar")
	defer os.Setenv("CONFIELD_TEST", "")

	if s.Field.String() != "envvar" {
		t.Fatalf("expect %s, get %s", s.Field.String(), "envvar")
	}
}
