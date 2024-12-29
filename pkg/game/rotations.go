package game

const RotationsCount int = 4
const ShapePoints int = 4
const ShapeCount int = 7

type Point struct {
	Row    int
	Column int
}

type Shape struct {
	DefaultPosition Point
	RotationsTable  [RotationsCount][ShapePoints]Point
}

var JShape = Shape{
	DefaultPosition: Point{Row: -1, Column: 3},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 2},
		},
		// rotation 1
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
		},
		// rotation 2
		{
			// four points
			Point{Row: 0, Column: 0},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
		},
		// rotation 3
		{
			Point{Row: 0, Column: 1},
			Point{Row: 0, Column: 2},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
		},
	},
}

var LShape = Shape{
	DefaultPosition: Point{Row: -1, Column: 3},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 0},
		},
		// rotation 1
		{
			// four points
			Point{Row: 0, Column: 0},
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
		},
		// rotation 2
		{
			// four points
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 0, Column: 2},
		},
		// rotation 3
		{
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
	},
}

var SShape = Shape{
	DefaultPosition: Point{Row: -1, Column: 3},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 2, Column: 0},
			Point{Row: 2, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
		},
		// rotation 1
		{
			// four points
			Point{Row: 0, Column: 0},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
		},
		// rotation 2 - eq to 0
		{
			// four points
			Point{Row: 2, Column: 0},
			Point{Row: 2, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
		},
		// rotation 3 - eq to 1
		{
			// four points
			Point{Row: 0, Column: 0},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
		},
	},
}

var ZShape = Shape{
	DefaultPosition: Point{Row: -1, Column: 3},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
		// rotation 1
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 0},
		},
		// rotation 2 - eq to 0
		{
			// four points
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
		// rotation 3 - eq to 1
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 0},
		},
	},
}

var TShape = Shape{
	DefaultPosition: Point{Row: -1, Column: 3},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
		},
		// rotation 1
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
		},
		// rotation 2
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 0},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
		},
		// rotation 3
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
		},
	},
}

var OShape = Shape{
	DefaultPosition: Point{Row: -1, Column: 4},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
		// rotation 1 - eq to 0
		{
			// four points
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
		// rotation 2 - eq to 0
		{
			// four points
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
		// rotation 3 - eq to 0
		{
			// four points
			Point{Row: 1, Column: 1},
			Point{Row: 1, Column: 2},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
		},
	},
}

var IShape = Shape{
	DefaultPosition: Point{Row: -2, Column: 3},
	RotationsTable: [RotationsCount][ShapePoints]Point{
		// rotation 0
		{
			// four points
			Point{Row: 2, Column: 0},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
			Point{Row: 2, Column: 3},
		},
		// rotation 1
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
			Point{Row: 3, Column: 1},
		},
		// rotation 2 - eq to 0
		{
			// four points
			Point{Row: 2, Column: 0},
			Point{Row: 2, Column: 1},
			Point{Row: 2, Column: 2},
			Point{Row: 2, Column: 3},
		},
		// rotation 3 - eq to 1
		{
			// four points
			Point{Row: 0, Column: 1},
			Point{Row: 1, Column: 1},
			Point{Row: 2, Column: 1},
			Point{Row: 3, Column: 1},
		},
	},
}

var ROTATIONS = [ShapeCount]*Shape{
	&JShape,
	&LShape,
	&SShape,
	&ZShape,
	&IShape,
	&OShape,
	&TShape,
}
