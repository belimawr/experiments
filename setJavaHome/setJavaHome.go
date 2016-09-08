package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

const baseJava string = "/usr/lib/jvm/"

func main() {
	java := os.Getenv("JAVA_HOME")

	if java != "" {
		fmt.Println("JAVA_HOME", java)
	}
	files, _ := ioutil.ReadDir(baseJava)
	var javas []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Found: ", file.Name())
			javas = append(javas, baseJava+file.Name())
		}
	}

	for i, f := range javas {
		fmt.Printf("[%d] %s\n", i, f)
	}

	var val int
	fmt.Print("Choose an option: ")
	fmt.Scanf("%d", &val)
	fmt.Println("User input: ", javas[val])

	os.Setenv("JAVA_HOME", javas[val])
	path := os.Getenv("PATH")
	os.Setenv("OLD_PATH", path)
	path = javas[val] + "/bin:" + path
	err := os.Setenv("PATH", path)
	if err != nil {
		panic(err)
	}
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}
