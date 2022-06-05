package main

import (
    "net/http"
    "fmt"
    "os"
    )


func getOsVersion(key string) (value string) {
    value = os.Getenv(key)
    return value
}


func index(w http.ResponseWriter, r *http.Request) {
    fmt.Println("打印Header参数列表：")
    if len(r.Header) > 0 {
       for k,v := range r.Header {
          w.Header().Set(k,v[0])
          fmt.Printf("%s=%sn", k, v[0])
       }
    }
    
    // 获取环境标量
    key := "VERSION"
    value := getOsVersion(key)
    fmt.Printf(value)
    w.Header().Set(key,value)

    str := "<h1 style='color:red'>哈哈哈哈</h1>"
//    w.Header().Set("Name", "my name is smallsoup")
   // w.WriteHeader(201)

    w.Write([]byte(str))
}

func healthz(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ok"))    
    w.WriteHeader(200)    
    

} 

func main() {
    http.HandleFunc("/index", index)
    http.HandleFunc("/healthz", healthz)
    http.ListenAndServe("127.0.0.1:8080",nil)
}
