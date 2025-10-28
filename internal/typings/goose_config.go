package typings

type Database struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
}

type Config struct {
	Migrations string   `json:"migrations"`
	Schemas    string   `json:"schemas"`
	Database   Database `json:"database"`
}
