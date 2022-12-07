package chain

var allChains = make([]Chain, 0)

var pointToChains = make(map[int][]int) // список всех точек и чейнов, которые из нее выходят

var coordinate Coord //координата работника текущая

var maxStep = 20

type Coord struct {
	ChainID int
	Dist    int
}

/*
  Тип, определяющий направление движения.
  Содержит точку и чейн, из нее выходящий.
  Например, если есть чейны 20-30, 40-30 и 40-50,
  то чтобы попасть из 20-30 в 40-50 нужно идти к точке 30, а затем в чейн 40-30
	Также хранится счетчик - расстояние в чейнах до нужного. В примере выше это будет 2.
	То есть добраться до чейна 40-50 можно через два прыжка
*/

type Direction struct {
	Point    int
	Chain    int
	Distance int
}

//Список всех направлений.
//Ключ - массив {a,b}, в котором a - чейн из которого идем, b - чейн, в который идем

var allDirections map[[2]int]Direction

type Chain struct {
	ID   int
	AC1  int
	AC2  int
	Dist int
}
