package main

type IniProperty struct {
	Name	string
	Value	string
}

type IniSection struct {
	Name		string
	Properties	[]IniProperty
}


type IniParser struct {
	IniSections map[string][]IniProperty
}


func NewIniParser() *IniParser {
	
	iniProperty_1 := &IniProperty {
		Name: "Color1",
		Value: "Red",
	}

	iniProperty_2 := &IniProperty {
		Name: "Color2",
		Value: "Blue",
	}

	sections := map[string][]IniProperty{"Colors": []IniProperty{*iniProperty_1, *iniProperty_2}}

	parser := &IniParser {
		IniSections: sections,
	}

	return parser
}


// parsers

