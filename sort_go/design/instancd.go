package design

var instancd []int

func NewInstance() []int {
	
	if instancd != nil {
		return instancd
	}
	instancd = make([]int, 0)
	return instancd
}

func main() {

}
