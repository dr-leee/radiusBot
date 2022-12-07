package chain

import (
	"log"
)

func GetChain(idx int) (*Chain, error) {
	return &allChains[idx], nil
}

func findNext(currChainID int, targetChainID int) (finalPoint int, finalChain int, steps int) {
	currChain, err := GetChain(currChainID)

	if err != nil {
		log.Panic(err)
	}
	finalPoint = -1
	finalChain = -1
	steps = -1
	//начинаем перебирать всех соседей со стороны ac1
	for _, v := range pointToChains[currChain.AC1] {
		if v != currChainID {
			road := make(map[int]int)
			road[currChain.AC1] = currChain.AC1
			currSteps := followMe(currChain.AC1, v, targetChainID, 1, road)
			if steps == -1 && currSteps > steps {
				finalPoint = currChain.AC1
				finalChain = v
				steps = currSteps
			}
			if currSteps != -1 && currSteps < steps {
				finalPoint = currChain.AC1
				finalChain = v
				steps = currSteps
			}
		}

	}
	//продолжаем перебирать соседей уже со стороны ac2
	for _, v := range pointToChains[currChain.AC2] {
		if v != currChainID {
			road := make(map[int]int)
			road[currChain.AC2] = currChain.AC2
			currSteps := followMe(currChain.AC2, v, targetChainID, 1, road)
			if steps == -1 && currSteps > steps {
				finalPoint = currChain.AC2
				finalChain = v
				steps = currSteps
			}
			if currSteps != -1 && currSteps < steps {
				finalPoint = currChain.AC2
				finalChain = v
				steps = currSteps
			}
		}

	}
	//возвращаем результат
	return
}

// возвращает количество шагов до нужного чейна, либо -1, если в этом направлении нельзя добраться до нужного
func followMe(point int, chain int, targetChainID int, steps int, road map[int]int) (dist int) {
	if chain == targetChainID {
		return steps
	} else {
		//если это не нужный нам чейн, идем вглубь
		dist = -1
		currChain, err := GetChain(chain)

		if err != nil {
			log.Panic(err)
		}
		//двигаемся дальше по чейну
		if currChain.AC1 != point {
			//проверяем, что мы ранее не были на этой точке.
			//Если были, то возвращаемся назад, по пройденному пути не идем
			if _, ok := road[currChain.AC1]; !ok {
				for _, v := range pointToChains[currChain.AC1] {
					if v != currChain.ID {
						road[currChain.AC1] = currChain.AC1
						currDistance := followMe(currChain.AC1, v, targetChainID, steps+1, road)
						//ищем самый короткий маршрут, либо выдаем, что его нет (-1)
						if dist == -1 && currDistance > 0 {
							dist = currDistance
						}
						if currDistance != -1 && currDistance < dist {
							dist = currDistance
						}

					}
				}
			}

			return
		} else {
			//проверяем, что мы ранее не были на этой точке.
			//Если были, то возвращаемся назад, по пройденному пути не идем
			if _, ok := road[currChain.AC2]; !ok {

				for _, v := range pointToChains[currChain.AC2] {
					if v != currChain.ID {
						road[currChain.AC2] = currChain.AC2
						currDistance := followMe(currChain.AC2, v, targetChainID, steps+1, road)
						//ищем самый короткий маршрут, либо выдаем, что его нет (-1)
						if dist == -1 && currDistance > 0 {
							dist = currDistance
						}
						if currDistance != -1 && currDistance < dist {
							dist = currDistance
						}
					}
				}
			}
			return
		}
		//		return followMe(, targetChainID)

		return -1
	}

}
