/* Q13 -> func Delete(map, key) bool, error
if map is nil return error
if key exists and delete return true, nil
if key does not exist return false and nil

func Update(map,key,value) bool,error
if map is nil then return error
if key does not exist return false and also error as key does not exist hence cannot Update
if key exists update the value and return true and nil*/

package main

import (
	"errors"
	"fmt"
)

func main() {

	mymap := make(map[string]any) // declaration but not instantiation

	mymap["omprakash"] = 453
	mymap["vedprakash"] = 786
	mymap["hariom"] = 564
	mymap["rahul"] = 766
	// fmt.Println(mymap)
	bool1, err := DeleteMapEle(mymap, "hario")
	fmt.Println("After Delating the element the Element status:", bool1, err)
	bool2, err1 := UpdateMapEle(mymap, "hariom",4655)
	fmt.Println("After Updating the element the Element status:" ,bool2,err1)

}

func DeleteMapEle(map1 map[string]any, str string) (bool, error) {

	if map1 == nil {
		return false, errors.New("map is empty you can be dewlaete element")
	}

	_, ok := map1[str]
	if ok {
		delete(map1, str)
		return ok, nil
	} else {
		return ok, errors.New("that key is not avilable so you can't delete that key")
	}

}

func UpdateMapEle(map2 map[string]any,key string ,val int) (bool,error){
	if map2==nil{
		return false ,errors.New("You can not update this key because map is empty")
	}

	_,ok:=map2[key]

	if ok{
           map2[key]=val
		   return true,errors.New("the particular key successfully updated")
	}else{
		return false,errors.New("you can't update that map key because that key is not avilable")
	}
	

}
