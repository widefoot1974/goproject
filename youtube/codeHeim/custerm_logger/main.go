package main

import (
	// "youtube/codeHeim/custerm_logger/logger"
	logger "youtube/codeHeim/custerm_logger/zaplogger"
)

func main() {

	// fmt.Println("This is a simple fmt print!")
	// log.Println("This is a simple log print!")

	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	// log.Println("This is another simple log print!")

	// // log.Panic("Something has gone wrong!")

	// log.Fatal("Something has gone bad terribly!")

	logger.SetLevel(logger.DebugLevel)

	logger.Debug("This is an debug print!")
	logger.Info("This is an info print!")
	logger.Warning("This is an warning print!")
	logger.Error("This is an error print!")
}
