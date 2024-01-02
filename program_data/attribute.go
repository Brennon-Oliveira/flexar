package program_data

import (
	"fmt"

	"github.com/Brennon-Oliveira/flexar/utils"
)

type Attribute struct {
	Name            string
	Type            string
	Overriden       bool
	Abstract        bool
	Readonly        bool
	PrivacyModifier utils.PrivacyModifier
	Static          bool
	Final           bool
}

func (c *Class) AddAttribute(attribute *Attribute) *Attribute {
	c.Attributes[attribute.Name] = attribute
	return attribute
}

func (c *Class) GetAttribute(name string) *Attribute {
	return c.Attributes[name]
}

func (c *Class) ExistsAttribute(name string) bool {
	_, ok := c.Attributes[name]
	return ok
}

func (a *Attribute) GetString() string {
	str := ""
	if a.Readonly {
		str += "readonly "
	}
	if a.Overriden {
		str += "override "
	}
	if a.Abstract {
		str += "abstract "
	}
	str += fmt.Sprintf("%d", a.PrivacyModifier)
	if a.Static {
		str += " static"
	}
	if a.Final {
		str += " final"
	}
	str += fmt.Sprintf(" %s", a.Name)
	if a.Type != "" {
		str += fmt.Sprintf(": %s", a.Type)
	}
	return str
}
