package main

func main() {
	sum := 2;

	fibPrev := 1;
	fibCurrent := 2;
	fibNew := 0;

	for fibCurrent < 4000000 {
		fibNew = fibPrev + fibCurrent;
		fibPrev = fibCurrent;
		fibCurrent = fibNew;

		if fibCurrent % 2 == 0 {
			sum += fibCurrent;
		}
	}

	println(sum);
}