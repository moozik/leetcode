//https://leetcode-cn.com/problems/perfect-rectangle/
package p391

import "fmt"

type pos struct {
	x int
	y int
}

func isRectangleCover(rectangles [][]int) bool {

	//总面积
	areaTotal := 0
	//记录最外层顶点位置
	pointVertexMax := rectangles[0]
	//记录顶点map
	pointVertexMap := make(map[pos]int)
	//遍历所有矩形
	for _, recItem := range rectangles {
		//统计面积
		areaTotal += (recItem[2] - recItem[0]) * (recItem[3] - recItem[1])
		//左下顶点
		if recItem[0] < pointVertexMax[0] || recItem[1] < pointVertexMax[1] {
			pointVertexMax[0] = recItem[0]
			pointVertexMax[1] = recItem[1]
		}
		//右上顶点
		if recItem[2] > pointVertexMax[2] || recItem[3] > pointVertexMax[3] {
			pointVertexMax[2] = recItem[2]
			pointVertexMax[3] = recItem[3]
		}
		pointVertexMap[pos{recItem[0], recItem[1]}]++
		pointVertexMap[pos{recItem[2], recItem[3]}]++
		pointVertexMap[pos{recItem[2], recItem[1]}]++
		pointVertexMap[pos{recItem[0], recItem[3]}]++
	}

	fmt.Printf("pointVertexMap,%+v,areaTotal:%+v, pointMax:%+v\n", pointVertexMap, areaTotal, pointVertexMax)
	//判断面积是否吻合
	if (pointVertexMax[2]-pointVertexMax[0])*(pointVertexMax[3]-pointVertexMax[1]) != areaTotal {
		return false
	}
	pointVertexLog := 0
	//判断同一位置顶点数量 除了最外部顶点,应该只能是2或者4
	for posItem, pointVertexCount := range pointVertexMap {
		if pointVertexCount != 2 && pointVertexCount != 4 {
			if pointVertexCount == 1 {
				pointVertexLog++
			}
			//不是顶点
			if posItem.x != pointVertexMax[0] && posItem.x != pointVertexMax[2] {
				return false
			}
			if posItem.y != pointVertexMax[1] && posItem.y != pointVertexMax[3] {
				return false
			}
		}
	}
	return pointVertexLog == 4
}
