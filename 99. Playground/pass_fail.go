// pass_fail 프로그램은 성적의 합격 여부를 알려줍니다
package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		logrus.Fatal(err.Error())
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
