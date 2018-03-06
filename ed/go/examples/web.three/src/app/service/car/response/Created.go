package response

type Created struct {
	Id int `json: "id"`
	Vendor string `json: "vendor"`
	Name string `json: "name"`
}
