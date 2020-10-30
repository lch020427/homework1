package main
import(
	"fmt"
)
func Receiver(v interface{}) {
	v=132
	switch v.(type) {
	case int:
		fmt.Printf("%d是%T型",v,v)
	case string:
		fmt.Printf("%s是%T型",v,v)
	case bool:
		fmt.Printf("%v是%T型",v,v)
	default:
		fmt.Printf("%T型",v)
	}
}
func main(){
	var v interface{}
	Receiver(v)
}