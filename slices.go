package main

import "fmt"

func main() {

  s := make([]string,3)
  fmt.Println("emp: ",s)

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])
  s = append(s,"d", "e", "f")
  fmt.Println("apd:", s)

  c := make([]string,len(s))
  copy(c,s)
  fmt.Println("cpy:", c)

  l := s[2:5]
  fmt.Println("sl1:", l)

  twoD := make([][]int, 3)
  for i:=0;i<3;i++{
    twoD[i] = make([]int,i+1)
    for j:=0;j<i+1;j++{
      twoD[i][j] = i+j
    }
  }
  fmt.Println("2d: ", twoD)
}
