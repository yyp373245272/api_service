package main

import "api_service/common"

func main() {
    s:=common.NewServer()
    register(s)
    err := s.Run(":3456")
    if err != nil {
        return 
    }
}
