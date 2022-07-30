package breezometer

type Metadata struct {
}

type IconCode int

// Measure is generic type for most of the measures that contain value and unit
type Measure struct {
	Value float64 `json:"value"`
	Units string  `json:"units"`
}
