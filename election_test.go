package election_test
import (
	"fmt"
	"testing"
	"com.aihe/leaderElection"
)

func ExampleHello()  {
	fmt.Println("hello")
	//Output: hello
}

func TestSomething(t *testing.T) {
	err := election.RegisterCallback(func(state int) {

	},"239.0.1.1:9999","eth1")

	if err != nil{
		t.Error(err)
	}
}
