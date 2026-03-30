func carFleet(target int, position []int, speed []int) int {
    cars := make([][2]int, len(position))                        // init array of pairs (position, speed)

    for i := 0; i < len(position); i++ {                         // iterate and populate array of pairs 
        cars[i] = [2]int{position[i], speed[i]}
    }
    sort.Slice(cars, func(i, j int) bool {                       // sort array of pairs by position descending
        return cars[i][0] > cars[j][0]
    })

    stack := []float64{}                                         // init stack for fleets (by arrival time)
    for _, car := range cars {                                   // iterate over cars, starting from nearest
        arrival := float64(target - car[0]) / float64(car[1])    // calculate arrival time of car
        if len(stack) != 0 && arrival <= stack[len(stack)-1] {   // if arrival faster than prev car ahead, fleet forms
            continue                                             // keep slower speed of prev car as fleet
        }
        stack = append(stack, arrival)                           // else, add arrival time as fleet to stack
    }
    return len(stack)                                            // return # of merged fleets
}
// monotonic ascending stack, logn for sorting, n for iteration, time O(nlogn)
