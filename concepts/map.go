package main

import "fmt"

type LatLng struct {
	lat , lng float64
}

var latLngMap map[string]LatLng

// A map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
// The make function returns a map of the given type, initialized and ready for use.

func makeMap() map[string]LatLng {
	latLngMap = make(map[string]LatLng)
	return latLngMap

}

func mapLiterals() {
	latLngMap := map[string]LatLng {
		"KallaBhattha, Chakwa, Near sarva UP GraminBank": LatLng{
			lat: 1.001,
			lng: 1.001,
		},
	}
	fmt.Println(latLngMap)
}

func addKeyValueToMap(key string, value LatLng, m map[string]LatLng) {
	m[key] = value
}

func deleteKeyOfMap(key string, m map[string]LatLng) {
	delete(m, key)
}

func main() {
	addressLines := "chakwa , balrampur, up, 271201"
	m := makeMap()
	addKeyValueToMap(addressLines,LatLng{1.001, 1.001}, m )
	mapLiterals()
	fmt.Println(m)
	deleteKeyOfMap(addressLines, m)
	fmt.Println(m)
}