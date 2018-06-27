package conf

import (
  `io/ioutil`
  `log`

  `gopkg.in/yaml.v2`
)

type config struct {
  DBUrl string `yaml:"dbUrl"`
}

var (
  c config
  done = false
)

func GetConf() *config {

  if done == true {
    return &c
  }

  f, err := ioutil.ReadFile("conf/conf.yaml")

  if err != nil {
    log.Println(err)
    return nil
  }

  err = yaml.Unmarshal(f, &c)

  if err != nil {
    log.Fatalln(err)
    return nil
  }

  done = true

  return &c
}
