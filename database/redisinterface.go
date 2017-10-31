package database

import (
  "gopkg.in/redis.v5"
)

type Redis struct{
  IP string
  Port string
  Password string
  DB int
  Client *redis.Client
}

func (r *Redis) Connect() (bool){
  r.Client = redis.NewClient(&redis.Options{
    Addr: r.IP + ":" +  r.Port,
    Password: r.Password,
    DB: r.DB,
  })
  _, err := r.Client.Ping().Result()
  if(err != nil){
    return false
  }else{
    return true
  }
}
func (r *Redis) Disconnect() bool{

  err := r.Client.Close()
  if(err != nil){
    return false
  }
  return true
}

func (r *Redis) Find(key string) string{
  return ""
}

func (r *Redis) AddString(key string, val string) bool{
  return false
}

func (r *Redis) AddInt(key string, value int) bool{
  return false
}

func (r *Redis) Delete(key string) bool{
  return false
}
