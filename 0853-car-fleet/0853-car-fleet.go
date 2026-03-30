func carFleet(target int, position []int, speed []int) int {
    cars := make([][2]int, len(position))

    for i := 0; i < len(position); i++ {
        cars[i] = [2]int{position[i], speed[i]}
    }
    sort.Slice(cars, func(i, j int) bool {
        return cars[i][0] > cars[j][0]
    })

    stack := []float64{}
    for _, car := range cars {
        arrival := float64(target - car[0]) / float64(car[1])
        if len(stack) != 0 && arrival <= stack[len(stack)-1] {
            continue
        }
        stack = append(stack, arrival)
    }
    return len(stack)
}

// initialize outer cars slice of inner arrays of size 2
// iterate over position array, for each position, add position + speed to cars slice as inner array
// sort cars slice by i > j for index 0 of inner array

// initialize stack for arrival times
// for each car in cars:
// calculate arrival = target - position / speed
// if stack not empty and arrival <= top element of stack:
// push top element back onto stack
// else, push arrival onto stack

// at end of iteration loop, length of stack == result (return length of stack)
