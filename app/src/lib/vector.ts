export class Vector2 {
	constructor(x: number, y: number) {
		this.x = x;
		this.y = y;
	}
	x: number;
	y: number;

	Equals(b: Vector2): boolean {
		return this.x == b.x && this.y == b.y;
	}
}

export function Diff(a: Vector2, b: Vector2): Vector2 {
	return new Vector2(a.x - b.x, a.y - b.y);
}

export function Add(a: Vector2, b: Vector2): Vector2 {
	return new Vector2(a.x + b.x, a.y + b.y);
}

export function Multiply(vec: Vector2, by: number): Vector2 {
	return new Vector2(vec.x * by, vec.y * by);
}

export function Equal(a: Vector2, b: Vector2): boolean {
	return a.x == b.x && a.y == b.y;
}

export function Average(a: Vector2, b: Vector2): Vector2 {
	return new Vector2((a.x + b.x) / 2, (a.y + b.y) / 2);
}

export function QuadraticBezier(t: number, p0: Vector2, p1: Vector2, p2: Vector2): Vector2 {
	const x = (p0.x - 2 * p1.x + p2.x) * t ** 2 + 2 * (p1.x - p0.x) * t + p0.x;
	const y = (p0.y - 2 * p1.y + p2.y) * t ** 2 + 2 * (p1.y - p0.y) * t + p0.y;
	return new Vector2(x, y);
}

export function GetPerpendicularRightDirection(p1: Vector2, p2: Vector2, dist: number): Vector2 {
	const multiplier = dist / Math.sqrt((p1.y - p2.y) ** 2 + (p2.x - p1.x) ** 2);
	const vector = new Vector2(p1.y - p2.y, p2.x - p1.x);

	return Multiply(vector, multiplier);
}
export function RemovedNeighboursDuplicates(list: Vector2[]): Vector2[] {
	const result: Vector2[] = [];
	for (let i = 1; i < list.length; i++) {
		const prevPoint = list[i - 1];
		const currPoint = list[i];
		if (!prevPoint.Equals(currPoint)) {
			result.push(prevPoint);
		}
	}
	result.push(list[list.length - 1]);

	return result;
}
