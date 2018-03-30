package main

import (
	"fmt"
	"strings"

)

const message = "theredfoxtrotsquietlyatmidnight"
const codeword = "bond"

const offset = 97

func main() {
	
	fmt.Println("Original Message:",message)

	encoded_string := encode_string(message,codeword)
	fmt.Println("Encoded Message: ",encoded_string)
	
	decoded_string := decode_string(message,codeword)
	fmt.Print("\nDecoded Messaged: ",decoded_string,"\n")
}

func encode_string(msg string, cde string) string {

	var encoded_message []string

	for i, _ := range message {	
		msg_int := int(msg[i]) - offset
		cde_int := int(cde[(i % len(codeword))]) - offset
		
		encoded_char := string(((msg_int+cde_int) % 26) + offset)
		encoded_message = append(encoded_message,encoded_char)
	}
	string_message := strings.Join(encoded_message,"")
	return string_message
}

func decode_string(msg string, cde string) string {

	var decoded_message []string

	for i, _ := range message {	
		diff := (msg[i] - cde[(i % len(codeword))])
		decoded_char := string(int(cde[(i % len(codeword))] + diff))
		decoded_message = append(decoded_message,decoded_char)	
	}
	string_message := strings.Join(decoded_message,"")
	return string_message
}
