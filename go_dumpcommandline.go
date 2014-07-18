/*
Description: iterate through commandline arguments
Author: gil cosson
*/

package main 
import "fmt" 
import "os" 

func main ( ) {
 fmt.Println("dump commandline")
 for i := 0;  i < len ( os.Args ); i ++ {
    fmt.Printf ( "%d] %s\n",i,os.Args [ i ] )
 }
}
