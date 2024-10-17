package style

type Style int

const (
	Slab         Style = 0
	Überhang     Style = 1
	Slopey       Style = 2
	Lang         Style = 3
	Kurz         Style = 4
	Statisch     Style = 5
	Dynamisch    Style = 6
	Stretchy     Style = 7
	Kante        Style = 8
	Kraftig      Style = 9
	Einfach_Hoch Style = 10
)

func (s Style) String() string {
	strings := [...]string{"Slab", "Überhang", "Slopey", "Lang", "Kurz", "Statisch", "Dynamisch", " Stretchy", "Kante", "Kraftig", "Einfach_Hoch"}

	if s < Slab || s > Einfach_Hoch {
		return "UNKNOWN"
	}

	return strings[s]
}
