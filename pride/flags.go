package pride

// Flags is a mapping for all pride flags in the library.
var Flags = map[string]func(){
	"trans":       PrintTransPrideFlag,
	"genderqueer": PrintGenderqueerPrideFlag,
	"pansexual":   PrintPansexualPrideFlag,
}
