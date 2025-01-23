package vaultsdict

import "fmt"

type VaultsDictionary struct {
	VaultSize int      // Should be manually defined. Does NOT include surrounding walls
	vaults    []*Vault // Should NOT include the surrounding walls
	Walls     []*Wall  // Surrounding walls presets
}

func (vd *VaultsDictionary) GetRandomVaultByParams(usualDoorsFlags int) *Vault {
	list := vd.getListOfVaultsMeetingCondition(
		func(v *Vault) bool {
			return rotationsUntilDirFlagsAreEqual(v.applicableDoorsMask, usualDoorsFlags) != -1
			return rotationsUntilFlagsIncluded(v.applicableDoorsMask, usualDoorsFlags) != -1
		})
	if len(list) == 0 {
		panic(fmt.Sprintf("No vault found for flags! Directions: %v", dirFlagsToVectors(usualDoorsFlags)))
	}
	vlt := list[rnd(len(list))]
	vlt.Rotate(rotationsUntilDirFlagsAreEqual(vlt.applicableDoorsMask, usualDoorsFlags))
	return vlt
}

func (vs *VaultsDictionary) getListOfVaultsMeetingCondition(condition func(v *Vault) bool) []*Vault {
	list := make([]*Vault, 0)
	for _, vlt := range vs.vaults {
		if condition(vlt) {
			list = append(list, vlt)
		}
	}
	return list
}
