package chain

import "math"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Chain {
	return allChains
}

func (s *Service) Show() Coord {
	return coordinate
}

func (s *Service) GetPoints() map[int][]int {
	return pointToChains
}

func (s *Service) Erase() {
	allChains = make([]Chain, 0)

	pointToChains = make(map[int][]int)

	coordinate = Coord{ChainID: 0, Dist: 0}
}

func (s *Service) Add(ac1 int, ac2 int, dist int) {
	//добавляем чейн в общий список при создании нового чейна
	chainID := len(allChains)
	allChains = append(allChains, Chain{
		ID:   chainID,
		AC1:  ac1,
		AC2:  ac2,
		Dist: dist,
	})
	//формируем вспомогательную мапу, которая хранит список всех точек с их чейнами
	if _, ok := pointToChains[ac1]; ok {
		pointToChains[ac1] = append(pointToChains[ac1], chainID) //если такая точка уже есть в списке, добавляем для нее новый чейн
	} else {
		pointToChains[ac1] = make([]int, 1) //если такой точки нет, добавляем ее в список
		pointToChains[ac1][0] = chainID
	}

	if _, ok := pointToChains[ac2]; ok {
		pointToChains[ac2] = append(pointToChains[ac2], chainID) //если такая точка уже есть в списке, добавляем для нее новый чейн
	} else {
		pointToChains[ac2] = make([]int, 1) //если такой точки нет, добавляем ее в список
		pointToChains[ac2][0] = chainID
	}
}

// вычисляем новую координату воркера
func (s *Service) Move(chainID int, dist int) (int, int, error) {
	//проверяем, находится ли новый сигнал в том же сегменте, где стоит воркер
	if coordinate.ChainID == chainID {
		//проверяем, что перемещение меньше шага
		if int(math.Abs(float64(coordinate.Dist-dist))) <= maxStep {
			return chainID, dist, nil
		} else {
			if dist > coordinate.Dist {
				return chainID, coordinate.Dist + maxStep, nil
			} else {
				return chainID, coordinate.Dist - maxStep, nil
			}
		}
	} else {
		//если сигнал пришел из другого сегмента, берем информацию куда двигаться из таблицы соответствий
		direction := allDirections[[2]int{coordinate.ChainID, chainID}]
		currChain, _ := GetChain(coordinate.ChainID)
		nextChain, _ := GetChain(direction.Chain)
		var halfStep int //переменная для расчета части шага между сегментами

		//если двигаемся в сторону AC1
		if direction.Point == currChain.AC1 {
			//если воркер стоит близко к точке перехода, то он перейдет на другой сегмент
			if coordinate.Dist < maxStep {
				halfStep = maxStep - coordinate.Dist
			} else {
				//если он не дойдет в этом шаге до следующего сегмента, то остаемся тут
				return coordinate.ChainID, coordinate.Dist - maxStep, nil
			}
		} else {
			//если двигаемся в сторону AC2
			//если воркер стоит близко к точке перехода, то он перейдет на другой сегмент
			if currChain.Dist-coordinate.Dist < maxStep {
				halfStep = maxStep - (currChain.Dist - coordinate.Dist)
			} else {
				//если он не дойдет в этом шаге до следующего сегмента, то остаемся тут
				return coordinate.ChainID, coordinate.Dist + maxStep, nil
			}
		}
		//начинаем расчет координаты воркера, если он перешел на соседний сегмент
		//проверяем, если сигнал пришел из соседнего чейна, то в нем и надо остановится
		if direction.Distance == 1 {
			//вычисляем остаток от пройденного пути больше, или меньше того места, куда надо поставить воркера
			if direction.Point == nextChain.AC1 {
				if halfStep >= dist {
					return direction.Chain, dist, nil
				} else {
					return direction.Chain, halfStep, nil
				}
			} else {
				//если заходим в чейн со стороны AC2
				if halfStep >= (nextChain.Dist - dist) {
					return direction.Chain, dist, nil
				} else {
					return direction.Chain, nextChain.Dist - halfStep, nil
				}
			}
		} else {
			//если идем не в соседний, а дальше, то прошагиваем на всю длину
			if direction.Point == nextChain.AC1 {
				return direction.Chain, halfStep, nil
			} else {
				//если заходим в чейн со стороны AC2
				return direction.Chain, nextChain.Dist - halfStep, nil
			}
		}

	}

}

// изменяем координату воркера
func (s *Service) Set(chainID int, dist int) {
	coordinate.ChainID = chainID
	coordinate.Dist = dist
}

// вычисляем новую карту расположения чейнов друг относительно друга

func (s *Service) Calc() {
	allDirections = make(map[[2]int]Direction)

	for _, a := range allChains {
		for _, b := range allChains {
			if a.ID != b.ID {
				point, chain, steps := findNext(a.ID, b.ID)
				allDirections[[2]int{a.ID, b.ID}] = Direction{Point: point, Chain: chain, Distance: steps}
			}
		}
	}
}
