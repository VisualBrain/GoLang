// methods
package main

import (
    "fmt"
 //   "strconv"
    )
type Form struct {
    username,password string
}
func (form *Form) authenticate() bool {
    if form.username == "Jatinder" && form.password == "anchor" {
        return true
    }
    return false
}
func (form *Form) authorize(val bool) string {
    if val == true {
                return "Admin"
    } else{
    return "Not Authorized"
    }
}
func main() {
    newTalk := Form{username : "Jatiender",password:"anchor"}
    
    /*--------------------Conversion from bool to string -------------------------*/
    //bool resu = newTalk.authenticate()
    //string rs := fmt.Sprint(newTalk.authenticate())
   // string rs := strconv.FormatBool(newTalk.authenticate())
    fmt.Println("Is authenticated:",newTalk.authenticate())
    fmt.Println(newTalk.authorize(newTalk.authenticate()))
}
