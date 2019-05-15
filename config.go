package main

type Config struct {
	Providers []*Provider
}

type Provider struct {
	Uri          string
	Replacements *Replacement
	Filters      *Filters
	Setters      []*Setter
	Groups       []*Group
}

type Setter struct {
	Name       string
	Attributes map[string]string
	Filters    *Filters
}

type Group struct {
	Name    string
	Filters *Filters
}

type Replacement struct {
	Name       []*Replacer
	Attributes map[string][]*Replacer
}

type Replacer struct {
	Find    string
	Replace string
}

type Filters struct {
	Name       *IncludeExcludeFilter
	Attributes map[string]*IncludeExcludeFilter
}