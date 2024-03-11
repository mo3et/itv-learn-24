package main

import ("fmt")

func main(){
  p1:=new(int)
  fmt.Printf("p1 --> %#v \n ", p1) //(*int)(0xc)
  fmt.Printf("p1 point to --> %#v \n ", *p1) //0

  p0:= new(string)

  fmt.Printf("p0 to --> %#v \n ", p0) //(*string)()
  fmt.Printf("p0 point to --> %#v \n ", *p0) //0

  var i *int
  p9=new(string)
  fmt.Printf("p0 to --> %#v \n ", p9) //(*string)()
  fmt.Printf("p0 point to --> %#v \n ", *p9) //0

  ii:= new(int)
  if ii==i{
fmt.Println("true.")
	fmt.Printf("ii %#v and i %#v",ii,i)
  }

}
