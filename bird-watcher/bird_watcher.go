package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int
	for i := range birdsPerDay {
		total += birdsPerDay[i]
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7
	list := birdsPerDay[start:end]
	return TotalBirdCount(list)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 != 0 {
			continue
		}

		birdsPerDay[i]++
	}

	return birdsPerDay
}
