package vaultsdict

import . "cycdg/graph_replacement/grid_graph/graph_element"

func CreateExampleDictionary() *VaultsDictionary {
	dict := &VaultsDictionary{
		VaultSize: 7,
		vaults:    exampleVaults,
		Walls:     exampleWallParcels,
	}
	return dict
}

var exampleVaults = []*Vault{
	// 4-door
	{
		applicableDoorsMask: DIR_ANY,
		implementsTags:      []string{""},
		charmap: []string{
			".......",
			".......",
			".......",
			".......",
			".......",
			".......",
			".......",
		},
	},
	{
		applicableDoorsMask: DIR_ANY,
		implementsTags:      []string{""},
		charmap: []string{
			"##...##",
			"##...##",
			".......",
			".......",
			".......",
			"##...##",
			"##...##",
		},
	},
	{
		applicableDoorsMask: DIR_ANY,
		implementsTags:      []string{""},
		charmap: []string{
			"###.###",
			"###.###",
			"##...##",
			".......",
			"##...##",
			"###.###",
			"###.###",
		},
	},
	// 3-door
	{
		applicableDoorsMask: DIR_NORTH | DIR_EAST | DIR_SOUTH,
		implementsTags:      []string{""},
		charmap: []string{
			"##...##",
			"##...##",
			"##.....",
			"##.....",
			"##.....",
			"##...##",
			"##...##",
		},
	},
	// 2-door
	{
		applicableDoorsMask: DIR_EAST | DIR_WEST,
		implementsTags:      []string{""},
		charmap: []string{
			"#######",
			".......",
			".......",
			".......",
			".......",
			".......",
			"#######",
		},
	},
	{
		applicableDoorsMask: DIR_NORTH | DIR_EAST,
		implementsTags:      []string{""},
		charmap: []string{
			"#......",
			"#......",
			"#......",
			"#......",
			"#......",
			"#......",
			"#######",
		},
	},
	// 1-door
	{
		applicableDoorsMask: DIR_EAST,
		implementsTags:      []string{""},
		charmap: []string{
			"#######",
			"#######",
			"###....",
			"###....",
			"###....",
			"#######",
			"#######",
		},
	},
	{
		applicableDoorsMask: DIR_EAST,
		implementsTags:      []string{""},
		charmap: []string{
			"#######",
			"##....#",
			"#......",
			"#......",
			"#......",
			"##....#",
			"#######",
		},
	},
}

var exampleWallParcels = []*Wall{
	{
		implementsTags: []TagKind{TagLockedEdge, TagBilockedEdge, TagMasterLockedEdge, TagOneTimeEdge, TagOneWayEdge},
		charmap:        "###+###",
	},
	{
		implementsTags: []TagKind{TagLockedEdge, TagBilockedEdge, TagMasterLockedEdge},
		charmap:        "##+#+##",
	},
	{
		implementsTags: []TagKind{TagLockedEdge, TagBilockedEdge, TagMasterLockedEdge},
		charmap:        "#+###+#",
	},
	{
		implementsTags: []TagKind{TagLockedEdge, TagBilockedEdge, TagMasterLockedEdge},
		charmap:        "##+++##",
	},
	{
		implementsTags: []TagKind{TagSecretEdge},
		resrictNullTag: true,
		charmap:        "##???##",
	},
	{
		implementsTags: []TagKind{TagSecretEdge},
		resrictNullTag: true,
		charmap:        "###?###",
	},
	{
		implementsTags: []TagKind{TagWindowEdge},
		resrictNullTag: true,
		charmap:        "##'''##",
	},
}
