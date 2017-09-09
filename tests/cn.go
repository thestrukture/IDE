package main

import (
	"fmt"
	"github.com/cheikhshift/gos/core"
)

func main() {

	//cmd := exec.Command("docker", "run", "-t", "-i", "-v","/home/strukture:/home/strukture", "ubuntu" ,"/bin/bash")
	//cmd := exec.Command("docker", "create","-t","-i", "-v" ,"/home:/home" ,"ubuntu")
	//cmd := exec.Command("docker", "start", "-a", "-i","b8c4e464184c3a04a67b6e61a7485b618b0e1ea519732d01d98c478d25674b6d")
	//cmd := exec.Command("docker","exec","a5a67c23d1f17f7694d2c24ab0b252106b0522c2f4f0294c9c29394958791356","echo","Test")

	id := core.RunCmdSmart("docker create -t -i -v /home/strukture/src/cseck:/home/strukture/src/cseck ubuntu_wgo")
	fmt.Println(core.RunCmdSmart("docker start " + id))
	fmt.Println(core.RunCmdSmart("docker exec " + id + " ls /"))
}
