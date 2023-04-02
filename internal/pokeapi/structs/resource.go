package structs

type Resource struct {
	Count    int      `json:"count"`
	Next     any      `json:"next"`
	Previous any      `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
