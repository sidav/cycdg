package consolemenus

import "fmt"

type ValueSelectorEntry struct {
	Title, UnitName               string
	MinValue, MaxValue, ValueStep int
	valPointer                    *int
}

func NewPointerSelectorEntry(valuePointer *int, title, unit string, min, def, max, step int) *ValueSelectorEntry {
	*valuePointer = def
	if *valuePointer < min {
		*valuePointer = min
	}
	return &ValueSelectorEntry{
		Title:      title,
		UnitName:   unit,
		MinValue:   min,
		MaxValue:   max,
		ValueStep:  step,
		valPointer: valuePointer,
	}
}

func (se *ValueSelectorEntry) GetValue() int {
	return *se.valPointer
}

func (se *ValueSelectorEntry) getValueStringWithArrows() string {
	str := fmt.Sprintf("%d", se.GetValue())
	if len(se.UnitName) > 0 {
		str = fmt.Sprintf("%d%s", se.GetValue(), se.UnitName)
	}
	prefix := "- "
	suffix := " -"
	if se.GetValue() > se.MinValue {
		prefix = "< "
	}
	if se.GetValue() < se.MaxValue {
		suffix = " >"
	}
	return prefix + str + suffix
}

func (se *ValueSelectorEntry) decrease() {
	*se.valPointer -= se.ValueStep
	if *se.valPointer < se.MinValue {
		*se.valPointer = se.MinValue
	}
}

func (se *ValueSelectorEntry) increase() {
	*se.valPointer += se.ValueStep
	if *se.valPointer > se.MaxValue {
		*se.valPointer = se.MaxValue
	}
}
