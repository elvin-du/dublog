package log

import(
    stdLog "log"
)

func init(){
    stdLog.SetFlags(stdLog.Lshortfile)
}

func Println(v ...interface{}){
    stdLog.Println(v)
}

func Fatalln(v ...interface{}){
    stdLog.Fatalln(v)
}
