
# Задание
1. При помощи датчика случайных чисел сгенерировать массив данных (по данным практической работы №3) размером 30, 200 элементов.
2. Выполнить поиск определенного заданного значения в неупорядоченном массиве для двух наборов данных (30 и 200) методом полного перебора. Вывести на экран индекс найденного элемента.
3. Выполнить сортировку массива данных (для двух наборов данных 30 и 200 элементов) по возрастанию применяя метод из практической работы №3.
4. Выполнить поиск определенного заданного значения в упорядоченном массиве для двух наборов данных (30 и 200) двумя методами: методом прямого перебора и выбранным методом. Вывести на экран индекс найденного элемента и количество итераций по выбранному методу поиска.
5. Проанализировать полученные результаты. Цель сортировки в п.3 облегчить последующий поиск элементов в упорядоченном массиве при обработке данных.

Метод поиска: Бинарный поиск

# Описание выбранного метода
Бинарный поиск -- алгоритм поиска элементов в отсортированном массиве. Работает следующим образом:
Пусть на вход подаётся массив отсортированный по возрастанию.
1. Берётся элемент массива посередине.
2. Если элемент больше искомого, то часть массива справа от элемента отбрасывается. Иначе отбрасывается часть слева.
3. Если элемент равен искомому возвращается его индекс.
4. К получившемуся после "отбрасывания" массиву применяем шаги 1-3, пока не найдём искомый элемент. Если не найдём, возвращаем ошибку.

В худшем случае имеет сложность `log_n`.

# Реализация
```
func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func SearchFloat64s(a []float64, x float64) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}
```

# Результаты работы
Формат: 
`ФункцияПоиска()`, `ДЛИННА_ВХОДНЫХ_МАССИВОВ`: `Среднее время поиска в наносекундах`

В результате тестирования получены следующие результаты:
```
// Non-sorted
findBruteForce(), LENGTH_SHORT:     115.752520 ns/op
findBruteForce(), LENGTH_LONG:      376.517420 ns/op
// Sorted
findBruteForce(), LENGTH_SHORT:     115.752520 ns/op
findBruteForce(), LENGTH_LONG:      376.517420 ns/op
searchFloat64s(), LENGTH_SHORT:     159.858830 ns/op
searchFloat64s(), LENGTH_LONG:      222.160090 ns/op
```

`LENGTH_SHORT = 30`
`LENGTH_LONG  = 500`, заместо `200`, дабы лучше показать сильные стороны бинарного поиска.
В каждом случаем поиск проводился `100000` раз. Затем вычислялось среднее время поиска.

# Выводы
Бинарный поиск имеет логарифмическую сложность. Логарифмическая функция резко возрастает в начале, поэтому бинарный поиск не слишком хорошо подходит для поиска в массиве небольшой длинны. Так, для массива длинной в `30` элементов бинарный поиск показывает даже худший результат, чем наивный метод полного перебора. С другой стороны, в массиве большой длинны скорость поиска едва возрастает даже при кратном увеличении длинны массива. В данном случае при росте с `30` до `500` элементов время поиска выросло всего лишь в `~1.4` раза, когда для наивного метода время выросло больше, чем в `3` раза. 