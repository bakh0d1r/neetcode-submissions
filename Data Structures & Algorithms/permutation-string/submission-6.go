func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }
    
    // Step 1: Populate the frequency map for s1
    set := map[byte]int{}
    for i := range s1 {
        set[s1[i]]++
    }
    
    n := len(s1)
    m := len(set)
    match := 0
    
    // Step 2: Initialize the first window of size len(s1)
    for i := 0; i < n; i++ {
        char := s2[i]
        if _, exists := set[char]; exists {
            set[char]--
            if set[char] == 0 {
                match++
            }
        }
    }
    
    if match == m {
        return true
    }
    
    // Step 3: Slide the window across s2
    for l := 1; l <= len(s2)-n; l++ {
        outChar := s2[l-1]   // Character leaving the window
        inChar := s2[l+n-1]  // Character entering the window
        
        // Handle the character leaving the window
        if _, exists := set[outChar]; exists {
            if set[outChar] == 0 {
                match--
            }
            set[outChar]++
        }
        
        // Handle the character entering the window
        if _, exists := set[inChar]; exists {
            set[inChar]--
            if set[inChar] == 0 {
                match++
            }
        }
        
        if match == m {
            return true
        }
    }
    
    return false
}