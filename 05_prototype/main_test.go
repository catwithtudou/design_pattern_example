package _5_prototype

import (
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

var manager *PrototypeManager

func init(){
	manager=NewPrototypeManager()
	manager.Set("t1",&Type1{name: "type1"})
	manager.Set("t2",&Type2{name:"type2"})
}

type Type1 struct{
	name string
}

func (t *Type1)Clone()CloneType{
	tc:=*t
	return &tc
}

type Type2 struct{
	name string
}

func (t *Type2)Clone()CloneType{
	tc:=*t
	return &tc
}

func TestClone(t *testing.T){
	t1:=manager.Get("t1")

	t2:=t1.Clone()



	if t1==t2{
		t.Fatal("error!")
	}
}

func TestClonePrototypeManager(t *testing.T) {
	c:=manager.Get("t1").Clone()

	t1:=c.(*Type1)
	if t1.name !="type1"{
		t.Fatal("error")
	}
}