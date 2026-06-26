
type Solution struct {
    Key string
}

func (s *Solution) Encode(strs []string) string {
    // Agar massiv butunlay bo'sh bo'lsa, to'g'ridan-to'g'ri "" qaytadi
    if len(strs) == 0 {
        return ""
    }

    var result []rune
    keyRune := '🍏'

    for _, str := range strs {
        result = append(result, []rune(str)...)
        result = append(result, keyRune) // Har bir elementdan keyin majburiy belgi qo'yamiz
    }
    return string(result)
}

func (s *Solution) Decode(e string) []string {
    // Agar string butunlay bo'sh bo'lsa, demak kiruvchi massiv [] bo'lgan
    if len(e) == 0 {
        return []string{}
    }

    var strs []string
    runes := []rune(e)
    keyRune := '🍏'
    start := 0

    for i := 0; i < len(runes); i++ {
        if runes[i] == keyRune {
            // Belgigacha bo'lgan qismni qirqib olamiz
            strs = append(strs, string(runes[start:i]))
            start = i + 1 // Ko'rsatkichni belgidan keyinga suramiz
        }
    }
    
    // Oxirgi belgidan keyin hech narsa qolmaydi (chunki har bir element oxirida belgi bor)
    // Shuning uchun bu yerda qo'shimcha append kerak emas.

    return strs
}