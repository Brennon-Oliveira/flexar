package program_data

import "github.com/Brennon-Oliveira/flexar/utils"

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
