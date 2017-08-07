package main

import (
  "fmt"
  "time"
  "os"
  "os/exec"
  "io/ioutil"
  "strings"
)


func checkStatus() {
    time.Sleep(10000 * time.Millisecond) //currently 10 seconds. 1s = 1000
  //  fmt.Println("Crontab updated")  //Uncomment to see a Message everytime it checks

     _, err := os.Stat("./renewTmp/renew1")
	if err == nil {
		     // fmt.Printf("File exists")
      		      renewBytes, _ := ioutil.ReadFile("./renewTmp/renew1")
                      renewStr := string(renewBytes[0:len(renewBytes)])
                      contents := strings.SplitN(renewStr," ",3)
                      addToCron(contents[0], contents[1], contents[2])
                      os.Remove("./renewTmp/renew1") //Cant be deferred as this process never ends
    } else {
    //fmt.Printf("File doesnt exist")
}                 

  
    
    checkStatus()
}

func addToCron (domainStr, lifeTimeStr, crtUuid string) {
  addTaskCommand := []string{"addTask.sh", domainStr, lifeTimeStr, crtUuid}
	fmt.Print(addTaskCommand)
        _,err := exec.Command("/bin/sh",addTaskCommand...).Output()
        if err != nil {
                fmt.Printf("El error es: %+v fin", err)
                panic(err)
        }
}



func main() {
    fmt.Println("RenewalManager status is: ACTIVE")
     checkStatus()
    
	  fmt.Println("RenewalManager status is: DOWN")
}

