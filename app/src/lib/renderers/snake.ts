import Konva from 'konva';
import { Add, Diff, Multiply, type Vector2 } from '../vector';

export const renderSnake = (
	snakeTail: Vector2[],
	direction: Vector2,
	cellSize: number,
	cellOffsetValue: number,
	options: {
		name: string;
		color: string;
	}
): Konva.Group => {
	const snakeGroup = new Konva.Group();
	if (snakeTail.length == 0) return snakeGroup;

	const adjustedSnake = [...snakeTail];
	const adjustedHeadPos = Add(Multiply(direction, cellOffsetValue - 0.8), snakeTail[0]);
	adjustedSnake[0] = adjustedHeadPos;

	if (snakeTail.length > 1) {
		const last = snakeTail[snakeTail.length - 1];
		const beforeLast = snakeTail[snakeTail.length - 2];
		const adjustedBackPos = Add(last, Multiply(Diff(beforeLast, last), cellOffsetValue - 0.1));
		adjustedSnake[adjustedSnake.length - 1] = adjustedBackPos;
	}

	if (snakeTail.length > 1) {
		const points: number[] = [];
		adjustedSnake.forEach((el: Vector2) =>
			points.push(el.x * cellSize + cellSize / 2, el.y * cellSize + cellSize / 2)
		);
		snakeGroup.add(
			new Konva.Line({
				points,
				bezier: false,
				stroke: options.color,
				strokeWidth: cellSize * 0.65,
				closed: false,
				lineCap: 'round',
				lineJoin: 'round',
				shadowOffsetY: 3,
				shadowOpacity: 0.2
			})
		);
	} else {
		snakeGroup.add(
			new Konva.Circle({
				x: adjustedHeadPos.x * cellSize + cellSize / 2,
				y: adjustedHeadPos.y * cellSize + cellSize / 2,
				width: cellSize * 0.75,
				height: cellSize * 0.75,
				fill: options.color,
				shadowOffsetY: 3,
				shadowOpacity: 0.2
			})
		);
	}
	const nameContainer = new Konva.Group({
		x: adjustedHeadPos.x * cellSize,
		y: adjustedHeadPos.y * cellSize
	});
	const nameText = new Konva.Text({
		width: cellSize,
		height: cellSize,
		text: options.name,
		fill: 'white',
		fontFamily: 'DM Sans',
		fontStyle: 'bold',
		fontSize: cellSize / 4.5,
		align: 'center',
		verticalAlign: 'middle',
		shadowBlur: 2,
		shadowOffsetY: 1
	});
	nameContainer.add(nameText);
	snakeGroup.add(nameContainer);

	return snakeGroup;
};
