package main

import(
	// "fmt"
)

func dec_upper_anphabet(one_str rune, seed int)(string){
	return string((int(one_str) - int('Z') - seed) % int(int('Z') - int('A') + 1) + int('Z'))
}

func dec_lower_anphabet(one_str rune, seed int)(string){
	return string((int(one_str) - int('z') - seed) % int(int('z') - int('a') + 1) + int('z'))
}

func decCaesar(one_str rune, seed int)(string){
	return string(int(one_str) + seed)
}

func dec(raw_msg string, seed int)(string){
	var enc_str string

	// fmt.Println(raw_msg, seed)

	for _, c := range raw_msg{
		// decCaesar(c, seed)
		if('a' <= c && c <= 'z'){
			enc_str += dec_lower_anphabet(c, seed)
		}else if('A' <= c && c <= 'Z'){
			enc_str += dec_upper_anphabet(c, seed)
		}else{
			enc_str += string(c)
		}
	}

	return enc_str
}


