package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0

	for _, item := range birdsPerDay {
		count += item
	}

	return count
}

const daysInWeek int = 7

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := week - 1
	start := startIndex * daysInWeek
	end := start + daysInWeek

	return TotalBirdCount(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, item := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] = item + 1
		}
	}

	return birdsPerDay
}
