package building

type BuildingProvider interface {
	GetAllBuildings() ([]Building, error)
	GetBuilding(name string) (Building, error)
	FilterBuildings(name string) ([]Building, error)
}
