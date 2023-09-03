//demo_14.go
package mapx

import (
	"fmt"
)

//pkcode映射
func GetPkcodeEmoji(emoji string) string {
	tmp := map[string]string{
		"🆚": "00", "🌈": "01", "🌏": "02", "🌚": "03", "🌞": "04", "🌟": "05", "🌱": "06", "🌵": "07", "🌸": "08", "🌼": "09",
		"🌿": "10", "🍄": "11", "🍆": "12", "🍋": "13", "🍖": "14", "🍵": "15", "🍺": "16", "🎁": "17", "🎂": "18", "🎃": "19",
		"🎉": "20", "🎓": "21", "🎤": "22", "🎮": "23", "🎹": "24", "🏀": "25", "🏆": "26", "🏥": "27", "🐂": "28", "🐇": "29",
		"🐉": "30", "🐍": "31", "🐎": "32", "🐒": "33", "🐔": "34", "🐛": "35", "🐟": "36", "🐢": "37", "🐮": "38", "🐯": "39",
		"🐰": "40", "🐱": "41", "🐲": "42", "🐴": "43", "🐵": "44", "🐶": "45", "🐷": "46", "🐸": "47", "🐹": "48", "👀": "49",
		"👂": "50", "👃": "51", "👅": "52", "👊": "53", "👋": "54", "👌": "55", "👍": "56", "👏": "57", "👨": "58", "👩": "59",
		"👮": "60", "👴": "61", "👶": "62", "👻": "63", "👽": "64", "👿": "65", "💄": "66", "💉": "67", "💊": "68", "💋": "69",
		"💏": "70", "💔": "71", "💡": "72", "💣": "73", "💤": "74", "💦": "75", "💩": "76", "💪": "77", "💭": "78", "💯": "79",
		"🔔": "80", "🔞": "81", "🔥": "82", "😂": "83", "😈": "84", "😉": "85", "😋": "86", "😎": "87", "😏": "88", "😐": "89",
		"😘": "90", "😝": "91", "😣": "92", "😤": "93", "😭": "94", "😱": "95", "😳": "96", "🙈": "97", "🙉": "98", "🙊": "99",
	}
	return tmp[emoji]
}

func GetPkuidByEmoji(emoji string) string {
	tmp := GetPkcodeEmoji(emoji[0:4]) + GetPkcodeEmoji(emoji[4:8])
	return tmp
}

func main2() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	fmt.Println("p1 :", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	fmt.Println("p2 :", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	fmt.Println("p3 :", p3)

	p4 := map[int]string{}
	p4[1] = "Tom"
	fmt.Println("p4 :", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	fmt.Println("p5 :", p5)

	p6 := map[int]string{
		1: "Tom",
	}
	fmt.Println("p6 :", p6)

	print(GetPkuidByEmoji("🆚🌈"))
	fmt.Println()
	fmt.Println(len("🆚"))
}