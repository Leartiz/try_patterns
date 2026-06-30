package main

import (
	"fmt"
	"sort"
)

// NOTE:
/*
	Дан набор интервалов `(start, finish)`.
	Нужно выбрать **максимум** интервалов, которые **не пересекаются**
	(если один заканчивается в момент `t`, другой может начинаться в `t`).

	```
	(1, 4), (3, 5), (0, 6), (5, 7), (8, 9), (5, 9)
	```

	Подсказка: по какому полю сортировать и что брать на каждом шаге.
*/

type Interval struct {
	Start  int
	Finish int
}

func maxNonOverlappingActivities(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Finish < intervals[j].Finish
	})

	result := 1
	lastFinish := intervals[0].Finish

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= lastFinish {
			result += 1
			lastFinish = intervals[i].Finish
		}
	}
	return result
}

func main() {
	// NOTE:
	/*
		 1__4
		   3_5
		0_____6
		     5_7
			    8_9
		     5____9

		Может быть брать минимальный размером интервал?
		И смотреть где заканчивается?
	*/

	intervals := []Interval{
		{1, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9},
	}

	{
		res := maxNonOverlappingActivities(intervals)
		fmt.Println(res)
	}
}
