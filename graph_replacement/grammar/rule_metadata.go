package grammar

import "slices"

type ruleMetadata struct {
	StepApplicability         StepCondition // Nil == applicable always
	AddsCycle                bool
	AddsTeleport             bool
	AdditionalWeight         int
	EnablesNodes             int
	EnablesNodesUnknown      bool // request full enabled nodes recalculation on apply
	FinalizesDisabledNodes   int
	UnfinalizesDisabledNodes int
}

// StepCondition is an applicability predicate.
type StepCondition func(step int) bool

// --- Basic applicability constructors ---

func Always() StepCondition {
	return func(step int) bool { return true }
}

func OnSteps(steps ...int) StepCondition {
	return func(step int) bool { return slices.Contains(steps, step) }
}

func Even() StepCondition {
	return func(step int) bool { return step%2 == 0 }
}

func Odd() StepCondition {
	return func(step int) bool { return step%2 != 0 }
}

func GreaterThan(x int) StepCondition {
	return func(step int) bool { return step > x }
}

func LessThan(x int) StepCondition {
	return func(step int) bool { return step < x }
}

// --- Combinations constructors ---

func And(conds ...StepCondition) StepCondition {
	return func(step int) bool {
		for _, c := range conds {
			if !c(step) {
				return false
			}
		}
		return true
	}
}

func Or(conds ...StepCondition) StepCondition {
	return func(step int) bool {
		for _, c := range conds {
			if c(step) {
				return true
			}
		}
		return false
	}
}

func Not(cond StepCondition) StepCondition {
	return func(step int) bool { return !cond(step) }
}
