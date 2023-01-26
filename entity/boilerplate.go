package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"-"`
}

type Boilerplate struct {
	BoilerplateKey
	Name string `json:"name,omitempty"`
}
