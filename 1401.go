func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	closestX := max(x1, min(xCenter, x2))
	closestY := max(y1, min(yCenter, y2))
	dx, dy := xCenter-closestX, yCenter-closestY
	return dx*dx+dy*dy <= radius*radius
}
