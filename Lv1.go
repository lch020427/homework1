package  main
import(
	"fmt"
)
type Author struct {
	Name string
	VIP bool
	Icon string
	Signature string
	Focus int
}
type All struct {
	Author
	Else []string
}
func main(){
	a:=Author{
		Name:"fengge",
		VIP:true,
		Icon:"",
		Signature: "",
		Focus: 120,
	}
	fmt.Println(a)
}