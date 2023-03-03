import type { Vector2 } from '$lib/vector';

export interface PageState {
	isPlaying: boolean;
	websocket: import('$lib/websocket').WebSocketClient | null;
	gameState: ClassicGameState | null;
	playerId: string | null;
	roomId: string | null;
	errors: {
		name: string | undefined;
		pin: string | undefined;
	};
}
export interface ClassicGameState {
	players: Player[];
	apples: Vector2[];
	gridSize: number;
	shrinkSize: number;
	frameTime: number;
}
export interface Player {
	id: string;
	name: string;
	color: string;
	snakeTail: Vector2[];
	score: number;
	direction: Vector2;
}

export interface JoinSuccessType {
	playerId: string;
	name: string;
	color: string;
}
export interface JoinError {
	error: string;
}
