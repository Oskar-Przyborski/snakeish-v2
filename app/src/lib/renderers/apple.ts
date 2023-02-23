import type { Vector2 } from '$lib/vector';
import Konva from 'konva';

export const renderApple = (
	position: Vector2,
	cellSize: number,
	options?: {
		color?: string;
		leafColor?: string;
	}
): Konva.Group => {
	const apple = new Konva.Group({
		x: position.x * cellSize,
		y: position.y * cellSize
	});
	apple.add(
		new Konva.Rect({
			x: cellSize * 0.125,
			y: cellSize * 0.125,
			width: cellSize * 0.75,
			height: cellSize * 0.75,
			fill: options?.color ?? '#d62246',
			cornerRadius: cellSize * 0.2,
			shadowOpacity: 0.2,
			shadowOffsetY: 2
		}),
		new Konva.Path({
			x: cellSize * 0.55,
			y: cellSize * 0.28,
			data: 'M-0.5 0Q0.5 0 0.5-1 -0.5-1 -0.5 0',
			fill: options?.leafColor ?? 'green',
			scaleX: cellSize * 0.35,
			scaleY: cellSize * 0.35,
			shadowOpacity: 0.2,
			shadowOffsetY: 0.15
		})
	);
	return apple;
};
