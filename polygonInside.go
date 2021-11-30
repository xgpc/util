package util

import "math"

/*
 *	目标：精确判断某个点是否在多边形内
 *	理论依据：目标点向右延伸，得出和多边形有几个交叉点（临界值判定规则：交叉点两端必须在射线两侧，判断依据：y1<y且y2>=y）
 *		如果是奇数 => 点在多边形内
 *		否则为偶数 => 点在多边形外
 */
func IsInPolygon(x, y float64, area [][2]float64) bool {
	// 1.1 取出多边形的最大和最小坐标
	var minX, minY, maxX, maxY float64 = 0, 0, 0, 0
	for _, px := range area {
		if px[0] < minX {
			minX = px[0]
		}
		if px[0] > maxX {
			maxX = px[0]
		}
		if px[1] < minY {
			minY = px[1]
		}
		if px[1] > maxY {
			maxY = px[1]
		}
	}

	// 1.2 如果P点在大矩形之外，直接返回false
	if x < minX || y < minY || x > maxX || y > maxY {
		return false
	}

	// 2.1 遍历，计算多边形和右射线的交叉点数量
	crossCnt := 0
	areaLen := len(area)

	for i := 0; i < areaLen; i++ {
		p1 := area[i]
		p2 := area[(i+1)%areaLen]

		x1, y1, x2, y2 := p1[0], p1[1], p2[0], p2[1]

		// 交叉线两端的点必须在射线两侧，即：y1和y2必须有一个小于y
		if y1 >= y && y2 >= y {
			continue
		}

		// 如果y高出y1和y2，也一定没有交叉
		if y > y1 && y > y2 {
			continue
		}

		// 根据公式判断
		if (y-y1)*(x2-x1)/(y2-y1)+x1 > x {
			crossCnt++
		}
	}

	return crossCnt%2 == 1
}

/**
 * 粗略判断目标点是否在多边形内，判断规则：
 * 1.精确判断目标在多边形内，返回True
 * 2.继续取出距离最近的一条边，计算距离，并根据误差值判断
 */
func IsInPolygonRough(x, y float64, area [][2]float64, deviation float64) bool {
	if IsInPolygon(x, y, area) {
		return true
	}

	px, py := x, y
	areaLen := len(area)
	var dist float64 = -1

	for i := 0; i < areaLen; i++ {
		a := area[i]
		b := area[(i+1)%areaLen]

		ax, ay, bx, by := a[0], a[1], b[0], b[1]

		abx := bx - ax
		aby := by - ay
		apx := px - ax
		apy := py - ay
		abAp := abx*apx + aby*apy
		distAb2 := abx*abx + aby*aby

		dx, dy := ax, ay
		if distAb2 != 0 {
			t := abAp / distAb2
			if t > 1 {
				dx, dy = bx, by
			} else if t > 0 {
				dx, dy = ax+abx*t, ay+aby*t
			}
		}

		pdx, pdy := dx-px, dy-py

		d := math.Sqrt(pdx*pdx + pdy*pdy)

		if dist == -1 || d < dist {
			dist = d
		}
	}

	// 结合误差判断是否在多边形内
	return dist < deviation
}
