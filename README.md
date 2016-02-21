# This is a library implementing dynamic leadership election using bully algorithm in GO.
### disclaimer
I am not an expert implementing bully algorithm. This was just a 'play project' learing Golang.

#### Installation

```sh
go get github.com/oaStuff/leaderElection
```
### Usage
The usage is very simple. All node (application running on a seperate machine or a separate process) that wish to participate in coordinated operation and allowing only one process become leader must do 
* register a callback with the library
* specify the multicast address to use
* specify the network interface to use

### Example
```go
func myCallback(state int){
	if state == election.LEADER {
		fmt.Println("Hey !, I am the leader and can do or setup things to become a leader")
	} else if state == election.FOLLOWER {
		fmt.Println("I am a follower")
	}
}

func main() {
    bChan := make(chan bool)
    err := election.RegisterCallback(myCallback,"239.0.1.1:9999","eth1")
	if err != nil{
		t.Error(err)
	}
	
	<-bChan
}
```
