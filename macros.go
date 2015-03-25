package crisp

import "log"

type Macro func(forms []Form) Form

type MacroExpander struct {
}

func (me *MacroExpander) Visit(form Form) (Form, error) {
	transformed, err := form.Accept(me)
	if err != nil {
		return nil, err
	}
	return transformed.(Form), nil
}

func (me *MacroExpander) VisitInt(i *Int) (interface{}, error) {
	return i, nil
}

func (me *MacroExpander) VisitString(s *String) (interface{}, error) {
	return s, nil
}

func (me *MacroExpander) VisitSymbol(s *Symbol) (interface{}, error) {
	return s, nil
}

func defn(forms []Form) Form {
	log.Println(forms)
	forms = append([]Form{}, forms...)
	return &List{forms}
}

func (me *MacroExpander) VisitList(list *List) (interface{}, error) {
	if len(list.Children()) == 0 {
		return list, nil
	}
	if list.Children()[0].Kind() == SymbolType {
		return defn(list.Children()[1:]), nil
	}
	return list, nil
}
