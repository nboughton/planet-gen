package table

import (
	"fmt"

	"github.com/nboughton/rollt"
)

// OneRoll represents the oft used one-roll systems spread throughout SWN
type OneRoll struct {
	D4  rollt.Able
	D6  rollt.Able
	D8  rollt.Able
	D10 rollt.Able
	D12 rollt.Able
	D20 rollt.Able
}

// Roll performs all rolls for a OneRoll and returns the results
func (o OneRoll) Roll() [][]string {
	return [][]string{
		{o.D4.Label(), o.D4.Roll()},
		{o.D6.Label(), o.D6.Roll()},
		{o.D8.Label(), o.D8.Roll()},
		{o.D10.Label(), o.D10.Roll()},
		{o.D12.Label(), o.D12.Roll()},
		{o.D20.Label(), o.D20.Roll()},
	}
}

func (o OneRoll) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", o.D4, o.D6, o.D8, o.D10, o.D12, o.D20)
}

// Attributes table to roll on
var Attributes = OneRoll{
	D4: rollt.List{
		Name: "Rule of Law",
		Items: []string{
			"Perilous",
			"Mostly Unsafe",
			"Mostly Safe",
			"Secure",
		},
	},
	D6: rollt.List{
		Name: "Population",
		Items: []string{
			"Established monoculture",
			"Cosmopolitan ",
			"Criminal/Pirates/Rebels",
			"Tribal/Regressive/Feral",
			"Outpost - Military/ Research/ Commercial",
			"Colonists",
		},
	},
	D8: rollt.List{
		Name: "Major Terrain Types",
		Items: []string{
			"Forest/Jungle",
			"Mountain/Hills",
			"Swamp",
			"Garden World",
			"Desert/Frozen/Toxic - Wasteland",
			"Underground Caves/Mines",
			"Water World",
			"Gas Giant",
		},
	},
	D10: rollt.List{
		Name: "Assets",
		Items: []string{
			"High-grade resource",
			"Tourism - unique culture or nature",
			"Unique life-forms draw researchers",
			"Alien Ruins and Tech",
			"A delicacy is harvested here",
			"An renowned place of learning",
			"A life-giving atmosphere/ spring/ resource that can’t be exported",
			"Dangerous fauna for the hunt",
			"A strategic position",
			"Massive manufactories",
		},
	},
	D12: rollt.List{
		Name: "Cultures",
		Items: []string{
			"Xenophobic",
			"Mercantile",
			"Hedonistic",
			"Spartan",
			"Inexplicably alien",
			"Friendly and Welcoming",
			"Non-Organic",
			"Religious Zealots",
			"Hivemind",
			"Aesthetes",
			"Extreme Pride",
			"Cold/calculating/no emotion",
		},
	},
	D20: rollt.List{
		Name: "Adventure Hooks",
		Items: []string{
			"Quarantine",
			"Refugees",
			"Civil War",
			"Invasion",
			"Tyranny/ Exploitation",
			"Scarcity",
			"Gold rush",
			"Impending disaster",
			"A wild/exotic festival",
			"Abandoned",
			"Momentus discovery",
			"Freak weather",
			"Battleground",
			"Pilgrimage",
			"Hostile flora/fauna",
			"Terrorism",
			"Corporate Takeover ",
			"Corruption",
			"A new religion",
			"It’s so nice here… why leave? Stay.",
		},
	},
}

// Names table to roll on
var Names = rollt.List{
	Name: "Name",
	Items: []string{
		"Craka V ",
		"New Yellowstone ",
		"New Alexandria ",
		"Fotti Prime ",
		"Astarte ",
		"I'Tedai ",
		"Chi-You ",
		"Phoebe ",
		"Ch'Deni ",
		"Kazi",
		"Hezitis ",
		"Giveria ",
		"Cholion ",
		"Nulrade ",
		"Duwei ",
		"Leanus ",
		"Dorscind's World ",
		"Goiturn ",
		"Bryke ",
		"1A4 Vinda RO",
		"Dyton",
		"Sihnon",
		"Higgins' Moon",
		"Ariel",
		"Londinium",
		"Liann Jiun",
		"Santo",
		"Triumph",
		"Three Hills",
		"Hera",
		"Lambda V",
		"Morloo IV",
		"Omid IV",
		"Gule IV",
		"Horki II",
		"Arcturus VI",
		"Brako VI",
		"Mu Arae VI",
		"Menkalinan VII",
		"Poxu II",
		"Hrane",
		"Siono",
		"Kote",
		"Gerte",
		"Yedin",
		"Palmary",
		"Zathru",
		"Axus",
		"Calfuu",
		"Kidia",
		"Lungor",
		"Munei",
		"Ekak",
		"Otaw",
		"Olok",
		"Anein",
		"Lonei",
		"Tsunei",
		"Eytaw",
		"Malu",
		"Sihnon",
		"New Melbourne",
		"Bernadette",
		"New Canaan",
		"Lazarus",
		"Parth",
		"Paquin",
		"St. Albans",
		"Iota Felis III",
		"Maenali VI",
		"Cassiopeiae IV",
		"Mega Cerberi",
		"Zorgi III",
		"Regulus Prime",
		"Pegasi III",
		"Pleione IV",
		"Theta Carinae",
		"Sagittae VI",
		"Gana",
		"Nara",
		"Beyscrim",
		"Bora",
		"Anosh",
		"Aros",
		"Myto",
		"Parmea",
		"Ablis",
		"Tala",
		"Chostrastea",
		"Sevozuno",
		"Kallilia",
		"Imiq",
		"Roabos",
		"Euwei",
		"Lluxetis",
		"Vaipra",
		"Comia UT5",
		"Thonoe 142",
		"Mercedes Creon",
		"Thalidae",
	},
}
