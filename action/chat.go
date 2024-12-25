package actions

import (
	"bufio"
	"fmt"
	"os"
)

func Chat(openaiApikey string){
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Talk to me: ")
		msg,readerErr := reader.ReadString('\n')
		if readerErr != nil{
			fmt.Println("Something went wrong",readerErr)
		}
		if msg == "exit" {
			break
		}
		
		output := chatWithOpenAI(msg)
		fmt.Println(output)
	}




}
func chatWithOpenAI(message string) string {


}
