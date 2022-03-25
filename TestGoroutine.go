package main

import(
	"fmt"
	"time"
)

func printInSec(s string, sec int, cnt int){

	start := time.Now()
	for i:=0; i < cnt; i++{
		fmt.Println(s, " index:", i)
		t := time.Now()
		elapsed := t.Sub(start)
		for (int)(elapsed.Seconds()) < sec {
			time.Sleep(time.Duration(100)*time.Millisecond)			
			t := time.Now()
			elapsed = t.Sub(start)
		}
		
		start = t		
	}
}


func mainRoutine(){
	printInSec("direct", 1, 5)
	
	go printInSec("routine1", 1, 10)
	
	go printInSec("routine2", 5, 10)
	
	defer printInSec("last", 1, 10)
	
	fmt.Println("bye")
}

func main(){
	go mainRoutine()
	
	time.Sleep(time.Duration(20)*time.Second)			
}