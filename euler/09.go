package main

//some math
//a^2 + b^2 = c^2
// a + b + c = 1000
//so c = 1000 - (a + b)
// a ^ 2 + b ^ 2 = (1000 - (a+b))^2
// a ^ 2 + b ^ 2 = 1000^2 - 2000*(a+b) + (a+b)^2
// a ^ 2 + b ^ 2 = 1000000 - 2000a - 2000b + a^2 + 2ab + b^2
// 2000a + 2000b = 1000000 + 2ab
// 2000(a+b) - 2ab = 1000000
// 2(1000(a+b) - 1) = 1000000
// 1000(a+b) - 1 = 500000
// 1000(a+b) = 5000001
// a+b = 5000001 / 1000
