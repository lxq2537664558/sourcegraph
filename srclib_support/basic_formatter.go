package srclib_support

import (
	"encoding/json"

	"sourcegraph.com/sourcegraph/srclib/graph"
)

// Registers def formatter for defs made by basic formatter (PHP, ObjC)
func init() {
	graph.RegisterMakeDefFormatter("basic-php", newBasicFormatter("PHP"))
	graph.RegisterMakeDefFormatter("basic-objc", newBasicFormatter("Objective-C"))
}

// DefData should be kept in sync with the def 'Data' field emitted by the
// basic grapher.
type DefData struct {
	Name      string
	Keyword   string
	Type      string
	Kind      string
	Separator string
}

// Constructs new def formatter for defs made by basic formatterx
func newBasicFormatter(lang string) graph.MakeDefFormatter {
	return func(s *graph.Def) graph.DefFormatter {
		var si DefData
		if len(s.Data) > 0 {
			if err := json.Unmarshal(s.Data, &si); err != nil {
				panic("unmarshal basic def data: " + err.Error())
			}
		}
		return basicFormatter{lang, s, &si}
	}
}

type basicFormatter struct {
	lang string
	def  *graph.Def
	data *DefData
}

func (f basicFormatter) Language() string {
	return f.lang
}

func (f basicFormatter) DefKeyword() string {
	return f.data.Keyword
}

func (f basicFormatter) Kind() string {
	return f.data.Kind
}

func (f basicFormatter) Name(qual graph.Qualification) string {
	if f.data.Name != `` {
		return f.data.Name
	}
	return f.def.Name
}

func (f basicFormatter) Type(qual graph.Qualification) string {
	return f.data.Type
}

func (f basicFormatter) NameAndTypeSeparator() string {
	return f.data.Separator
}
