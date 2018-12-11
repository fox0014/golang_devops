package main
 
import (
    "bufio"
    "log"
    "fmt"
    "strings"
    "os/exec"
    "flag"
)

var(
    aa bool
)

func mytest1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func init() {
    flag.BoolVar(&aa, "aa", false, "this help")
}
 
func main() {
    flag.Parse()
    mycmd1 := "ansible-playbook"
    var mycmd *string
    mycmd = &mycmd1
    var myhello string = "project ----------------------- begin"
    var myend string = "project ----------------------- over"
    var myoption string = "-v /etc/ansible/roles/devops/evlink-idc-api.yml -e my_env_name=prod"
    var argArray []string
    if myoption != "" {
        argArray = strings.Split(myoption, " ")
     } else {
         argArray = make([]string, 0)
   }
    log.Println(myhello)
    fmt.Printf("%c[1;45;33m%s%c[0m\n", 0x1B,"----------------------------------------------", 0x1B)
//    mytest1(argArray...)
    cmd := exec.Command(*mycmd, argArray...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
        return
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
        return
    }
    outputBuf := bufio.NewReader(stdout)
    for {
        output,_,err := outputBuf.ReadLine()
        if err != nil {
            if err.Error() != "EOF" {
               log.Fatal(err)
            }
           fmt.Printf("%c[1;45;33m%s%c[0m\n", 0x1B,"----------------------------------------------", 0x1B)
           log.Print(myend)
           return
         }
//        fmt.Printf("%s\n", string(output))
        fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, string(output), 0x1B)
}
     if err := cmd.Wait(); err != nil {
        log.Print(myend)
        return
       }
    log.Print(myend)
}

