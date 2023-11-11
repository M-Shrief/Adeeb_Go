package component_poet

type Poet struct {
	ID   string `json:"id" redis:"id"` // use redis tags when using redis Hashmaps
	Name string `json:"name" redis:"name"`
	Bio  string `json:"bio" redis:"bio"`
}
