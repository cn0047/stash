package one

type Config struct {
	ProjectName string
	Env         string
}

func InitConfig(env string) (Config, error) {
	c := Config{
		ProjectName: "test",
		Env:         env,
	}

	return c, nil
}
