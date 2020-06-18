package bootstrap

import "github.com/kelseyhightower/envconfig"

const(
	serviceName = "MYAPP"
)

type Config struct {
	DatabaseName string `envconfig:"DATABASE_NAME"`
	DatabasePort string `envconfig:"DATABASE_PORT"`
	DatabaseHost string `envconfig:"DATABASE_HOST"`
	DatabaseUser string `envconfig:"DATABASE_USER"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD"`

}
// This is just hold methods
type BS struct {}
func (a BS) Boot() (Config,error)  {
	var conf Config
	err := envconfig.Process(serviceName,&conf)
	return conf,err
}