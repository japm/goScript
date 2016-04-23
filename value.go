/*
The MIT License (MIT)
Copyright (c) 2016 Juan Pascual
*/
package goScript


type value interface{
    ToInterface()(interface{})
    ToInt64()(int64)
}


type vInt int64


func buildValue(v interface{})(value){
  return vInt(v.(int64))
}

func  (v vInt) ToInterface()(interface{}){
  return int64(v)
}

func  (v vInt) ToInt64()(int64){
  return int64(v)
}
