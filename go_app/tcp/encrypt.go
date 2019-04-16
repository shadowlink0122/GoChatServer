package main

import(
	// "fmt"
)

// func upper_anphabet(one_str rune, seed int)(rune){
// 	return (one_str - "A" + seed) % int("Z" - "A" + 1) + "A"
// }

func enc_upper_anphabet(one_str rune, seed int)(string){
	return string((int(one_str) - int('A') + seed) % int(int('Z') - int('A') + 1) + int('A'))
}

func enc_lower_anphabet(one_str rune, seed int)(string){
	return string((int(one_str) - int('a') + seed) % int(int('z') - int('a') + 1) + int('a'))
}

func encCaesar(one_str rune, seed int)(string){
	return string(int(one_str) - seed)
}

func enc(raw_msg string, seed int)(string){
	var enc_str string

	// fmt.Println(raw_msg, seed)

	for _, c := range raw_msg{
		// enc_str = encCaesar(c, seed)
		if('a' <= c && c <= 'z'){
			enc_str += enc_lower_anphabet(c, seed)
		}else if('A' <= c && c <= 'Z'){
			enc_str += enc_upper_anphabet(c, seed)
		}else{
			enc_str += string(c)
		}
	}

	return enc_str
}


