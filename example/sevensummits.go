// Generated code. DO NOT EDIT.

package sevensummits

// Peak represents a mountain.
type Peak struct {
	Name string
	// Elevation in meters.
	Elevation int
	Continent string
	// Name of mountain range.
	Range           string
	Country         string
	FirstAscentYear int
}

// Peaks is the Messner version of the Seven Summits.
var Peaks = []Peak{
	{"Mount Everest", 8848, "Asia", "Himalaya", "Nepal/China", 1953},
	{"Aconcagua", 6961, "South America", "Andes", "Argentina", 1897},
	{"Denali", 6194, "North America", "Alaska Range", "United States", 1913},
	{"Kilimanjaro", 5895, "Africa", "", "Tanzania", 1889},
	{"Mount Elbrus", 5642, "Europe", "Caucasus Mountains", "Russia", 1874},
	{"Mount Vinson", 4892, "Antarctica", "Sentinel Range", "", 1966},
	{"Puncak Jaya", 4884, "Australasia", "Sudirman Range", "Indonesia", 1962},
}
