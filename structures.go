package main

import (
  "fmt"
  "time"
)

type Person struct {
    Name string
    EmailId string
    Age int
    Weight float64
    Country string
    State string
    CreatedOn string
}

type Account struct {
    isLogin bool
    *Person
    timeOfTry string
}

func main(){
    var name,emailId,country,state string
    var age int 
    var weight float64
    fmt.Println("Taking Inputs:\n")
    //createdOn = time.Now().Local().String()
    fmt.Println("Enter Person Name:")
    fmt.Scanln(&name)
    fmt.Println("Enter Email Id:")
    fmt.Scanln(&emailId)
    fmt.Println("Enter Age:")
    fmt.Scanln(&age)
    fmt.Println("Enter Weight:")
    fmt.Scanln(&weight)
    fmt.Println("Enter Country:")
    fmt.Scanln(&country)
    fmt.Println("Enter State:")
    fmt.Scanln(&state)
    //fmt.Println("Name:",name,createdOn)
    //fmt.Println("Current Time is:",dt.Local())
   /* 
    Profile := Person{Name:name,EmailId:emailId,Age:age,Weight:weight,Country:country,State:state,CreatedOn:createdOn}
    fmt.Println(Profile)
    changeAge(Profile,34)
    fmt.Println(Profile)
    */
    Profile := createNewProfile(name,emailId,age,weight,country,state)
    fmt.Println(Profile)
    changeAgeByReference(Profile,45)
    fmt.Println(Profile)
    Profile.Coun6()
    fmt.Println(Profile)
    
    AccountLogin := makeLogin(Profile)
    fmt.Println(AccountLogin.isLogin)
    fmt.Println(AccountLogin.Person)

    fmt.Println(AccountLogin)

  /*
            output:
                    &{hello kskksks 24 65.5 sksk sksk 2019-07-08 08:53:34.257549226 +0000 UTC}
                    &{hello kskksks 45 65.5 sksk sksk 2019-07-08 08:53:34.257549226 +0000 UTC}
  */
}
func createNewProfile(name string,emailId string,age int,weight float64,country string,state string) *Person {
    /*return &Person{
        Name:name,
        EmailId:emailId,
        Age:age,
        Weight:weight,
        Country:country,
        State:state,
        CreatedOn:time.Now().Local().String(),
    }*/
    //We can also Use new Keyword on struct
    profile := new(Person)
    profile.Name=name
    profile.EmailId=emailId
    profile.Age=age
    profile.Weight=weight
    profile.Country=country
    profile.State=state
    profile.CreatedOn=time.Now().Local().String()
    return profile
}
func makeLogin(person *Person) *Account {
    return &Account{
        isLogin:true,
        Person:person,
        timeOfTry: time.Now().Local().String(),
    }
}
func changeAge(person Person,age int) {  //Take copy of Person and change it ,not changing the original value
    person.Age = age
}
//Creating Receiver
func (person *Person) Coun6(){  //*Person is the receiver of the Coun6 method
    person.Country ="New Zealand"
}
func changeAgeByReference(person *Person,age int){
        person.Age = age
}
