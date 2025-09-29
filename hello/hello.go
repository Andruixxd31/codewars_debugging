package kata

import (
	"fmt"
	"strings"
)

type Dinglemouse struct {
	name  string
	age   int
	sex   rune
	order []string
	keys  map[string]bool
}

// TODO:
// dont rewrite an already existing attribute -> modify Setters to check if value is set
// Write the string as the methods are calle -> a string var can be used to keep concating

func NewDinglemouse() *Dinglemouse {
	return &Dinglemouse{
		keys: make(map[string]bool),
	}
}

func (d *Dinglemouse) SetAge(age int) *Dinglemouse {
	addIfNew(d.keys, &d.order, "age")

	d.age = age
	return d
}

func (d *Dinglemouse) SetSex(sex rune) *Dinglemouse {
	addIfNew(d.keys, &d.order, "sex")

	d.sex = sex
	return d
}

func (d *Dinglemouse) SetName(name string) *Dinglemouse {
	addIfNew(d.keys, &d.order, "name")

	d.name = name
	return d
}

func addIfNew(keys map[string]bool, order *[]string, key string) {
	if _, ok := keys[key]; !ok {
		keys[key] = true
		*order = append(*order, key)
	}
}

func (d *Dinglemouse) Hello() string {
	sexText := "male"

	if d.sex == 'F' {
		sexText = "female"
	}

	var b strings.Builder
	b.WriteString("Hello.")

	for _, v := range d.order {
		switch v {
		case "name":
			b.WriteString(fmt.Sprintf(" My name is %s.", d.name))
		case "age":
			b.WriteString(fmt.Sprintf(" I am %d.", d.age))
		case "sex":
			b.WriteString(fmt.Sprintf(" I am %s.", sexText))
		}
	}

	return b.String()
}
