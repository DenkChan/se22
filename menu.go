package main
 
import "fmt"
/*  C 
int main()
{
    while(true)
    {
        scanf(cmd);
        int ret = strcmp(cmd, "help");
        if(ret == 0)
        {
            dosth();
        }
        int ret = strcmp(cmd, "others");
        if(ret == 0)
        {
            dosth();
        }
    }
}*/

func main() {
    for {
        var command string
        fmt.Println("input:")
        fmt.Scanln(&command)
        if command == "help" {
            fmt.Println("command list:")
            fmt.Println("--quit")
            fmt.Println("--hello")
        } else if command == "quit" {
            break
        } else if command == "hello" {
            fmt.Println("hello")
        } else {
            fmt.Println("command not found, try \"help\"")
        }
 
    }
}