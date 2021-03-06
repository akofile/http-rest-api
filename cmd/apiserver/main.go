package main
import(
	"flag"
	"log"
	"github.com/BurntSushi/toml"
	"github.com/akofile/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configpath", "configs/apiserver.toml", "path to config file")
}

func main(){
	flag.Parse()
	config := apiserver.NewConfig()
	_,err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err:=s.Start; err!=nil{
		log.Fatal(err)
	}
}
