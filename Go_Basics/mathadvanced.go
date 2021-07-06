package main

import ("fmt"
        "math/big"
        "math"
)

func main(){
    /* big.float is a data type, give more precise output
     this is explicit declaration*/

     var b1, b2, b3, sum big.Float
     b1.SetFloat64(2.5)
     b2.SetFloat64(7.5)
     b2.SetFloat64(3.2)

     sum.Add(&b1, &b2).Add(&sum, &b3)
     fmt.Printf("Big Sum better precision %.10g\n", &sum)

    /* the verb .2f means output should contain only
    2 characters after
    the decimal */
     circleRadius := 18.5
     circumference := circleRadius * math.Pi
     /* printf takes verbs like .2 f etc, println doesn't */
     fmt.Printf("Circumference: %.2f\n", circumference)
}