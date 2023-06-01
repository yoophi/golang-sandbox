package main

import (
	"log"
	"os"
)

func main() {
	fpLog, err := os.OpenFile("myApp.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if err != nil {
		panic(err)
	}

	defer fpLog.Close()

	// 표준 logger 의 출력을 지정한다.
	log.SetOutput(fpLog)
	run()
	log.Println("End of program")
}

func run() {
	log.Println("run()")
}
