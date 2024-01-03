package consolemenus

import "fmt"

type ValueSelectorEntry struct {
	Title, UnitName               string
	MinValue, MaxValue, ValueStep int

	currentValue int
}

func NewValueSelectorEntry(title, unit string, min, def, max, step int) *ValueSelectorEntry {
	value := def
	if value < min {
		value = min
	}
	return &ValueSelectorEntry{
		Title:        title,
		UnitName:     unit,
		MinValue:     min,
		MaxValue:     max,
		ValueStep:    step,
		currentValue: value,
	}
}

func (se *ValueSelectorEntry) getValueStringWithArrows() string {
	str := fmt.Sprintf("%d", se.currentValue)
	if len(se.UnitName) > 0 {
		str = fmt.Sprintf("%d%s", se.currentValue, se.UnitName)
	}
	prefix := "- "
	suffix := " -"
	if se.currentValue > se.MinValue {
		prefix = "< "
	}
	if se.currentValue < se.MaxValue {
		suffix = " >"
	}
	return prefix + str + suffix
}

func (se *ValueSelectorEntry) decrease() {
	se.currentValue -= se.ValueStep
	if se.currentValue < se.MinValue {
		se.currentValue = se.MinValue
	}
}

func (se *ValueSelectorEntry) increase() {
	se.currentValue += se.ValueStep
	if se.currentValue > se.MaxValue {
		se.currentValue = se.MaxValue
	}
}
