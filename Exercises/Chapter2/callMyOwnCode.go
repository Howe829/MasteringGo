package main

import "C"

//#include<stdio.h>
//int goCall(int x){
//    printf("This is will call in Go, written by Howe!%d\n", x);
//    return 0;
//}
import "C"

func main() {
	C.goCall(45)
}
