package model

import(
    "dublog/db"
//    "dublog/"
)

type homeModel byte

var _homeModel homeModel = 1

func HomeModel()*homeModel{
    return &_homeModel
}

func (*homeModel)Index()(map[string]string,error){
    db.NewContext()
    return map[string]string{"user":"elvindu"},nil
}
