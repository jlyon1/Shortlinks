package database

import (
)

type DB interface{
  Connect() bool
  Disconnect() bool
  Find(key string) string
  AddString(key string, val string) bool
  AddInt(key string, value int) bool
  Delete(key string) bool

}
