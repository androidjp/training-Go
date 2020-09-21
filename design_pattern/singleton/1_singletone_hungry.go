package singleton

type config struct {
	rotate bool
}

func (cfg *config) SetRotate(isRotate bool) {
	cfg.rotate = isRotate
}

func (cfg *config) IsRotate() bool {
	return cfg.rotate
}

var cfg *config = new(config)

func GetConfig() *config {
	return cfg
}
