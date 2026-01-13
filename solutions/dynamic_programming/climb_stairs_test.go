package dynamic_programming

import "testing"

func TestClimbStairs(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"1 step", 1, 1},
        {"2 steps", 2, 2},
        {"5 steps", 5, 8},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := ClimbStairs(test.input)
            
            if result != test.expected {
                t.Errorf("ClimbStairs(%d) = %d, expect %d", 
                    test.input, result, test.expected)
            }
        })
    }
}
