package maplib


type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func PrintLocation (l Location) {
	fmt.Println("this is lat: ", l.Latitude)
	fmt.Println("this is lon: ", l.Longitude)
}
