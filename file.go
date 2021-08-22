package main

//func (p Person) SaveToFile(person Person) error{
//	var persons Person
//	persons= person
//	file, err := json.MarshalIndent(persons, "", " ")
//	if err != nil {
//		e := apiErrors("fail in format file to json", err)
//		log.Println(e)
//		return err
//	}
//	_ = ioutil.WriteFile("persons.json", file, 0666)
//
//	return nil
//}