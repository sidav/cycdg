package replacement

import (
	. "cycdg/graph_replacement/geometry"
	. "cycdg/graph_replacement/grammar"
	"fmt"
)

// This struct is needed for logging only
type AppliedRuleInfo struct {
	ruleName, mandatoryFeatureName, optionalFeatureName string
	appliedAt                                           []Coords
}

func newAppliedRuleInfoInitial(rule *InitialRule, mandatory *FeatureAdder, x, y int) *AppliedRuleInfo {
	ari := &AppliedRuleInfo{
		ruleName:  rule.Name,
		appliedAt: []Coords{NewCoords(x, y)},
	}
	if mandatory != nil {
		ari.mandatoryFeatureName = mandatory.Name
	}
	return ari
}

func newAppliedRuleInfo(rule *ReplacementRule, mandatory, optional *FeatureAdder, coords []Coords) *AppliedRuleInfo {
	ari := &AppliedRuleInfo{
		ruleName:  rule.Name,
		appliedAt: coords,
	}
	if mandatory != nil {
		ari.mandatoryFeatureName = mandatory.Name
	}
	if optional != nil {
		ari.optionalFeatureName = optional.Name
	}
	return ari
}

func (ari *AppliedRuleInfo) StringifyRule() string {
	ruleString := ari.ruleName
	if ari.mandatoryFeatureName != "" {
		ruleString += " (" + ari.mandatoryFeatureName + ")"
	}
	if ari.optionalFeatureName != "" {
		ruleString += " + " + ari.optionalFeatureName
	}
	return ruleString
}

func (ari *AppliedRuleInfo) StringifyCoords() string {
	coordString := ""
	for i := range ari.appliedAt {
		coordString += fmt.Sprintf("%d,%d  ", ari.appliedAt[i].X, ari.appliedAt[i].Y)
	}
	return coordString
}
