# databundler

Embed CSV data as a Golang package.

## Install

```
go get github.com/mmcloughlin/databundler
```

## Example

We'll use the [Seven Summits](https://en.wikipedia.org/wiki/Seven_Summits) as
our [example](example). The CSV we are going to embed into a package looks like

```
Mount Everest,8848,8848,Asia,Himalaya,Nepal/China,1953
Aconcagua,6961,6961,South America,Andes,Argentina,1897
Denali,6194,6144,North America,Alaska Range,United States,1913
Kilimanjaro,5895,5885,Africa,,Tanzania,1889
Mount Elbrus,5642,4741,Europe,Caucasus Mountains,Russia,1874
Mount Vinson,4892,4892,Antarctica,Sentinel Range,,1966
Puncak Jaya,4884,4884,Australasia,Sudirman Range,Indonesia,1962
```

We specify a *schema* in YAML format.

```yaml
name: Peak
doc: Peak represents a mountain.
collection_doc: Peaks is the Messner version of the Seven Summits.
fields:
  - name: Name
    type: string
  - name: Elevation
    doc: Elevation in meters.
    type: int
  - null # ignore Prominence
  - name: Continent
    type: string
  - name: Range
    doc: Name of mountain range.
    type: string
  - name: Country
    type: string
  - name: FirstAscentYear
    type: int
```

Build the Go source with the following commands.

```sh
databundler -pkg sevensummits -data sevensummits.csv -schema schema.yaml -output sevensummits.go
gofmt -w sevensummits.go
```

Resulting in the following code.

```go
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
```