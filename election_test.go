package election_test
import (
	"testing"
	"github.com/oaStuff/leaderElection"
	"fmt"
)

func myCallback(state int){

	if state == election.LEADER {
		fmt.Println("Hey !, I am the leader and can do or setup things to become a leader")
	} else if state == election.FOLLOWER {
		fmt.Println("I am a follower")
	}
}

func TestSomething(t *testing.T) {

	err := election.RegisterCallback(myCallback,"239.0.1.1:9999","eth1")
	if err != nil{
		t.Error(err)
	}

}
