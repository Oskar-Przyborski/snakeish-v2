// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	// interface Error {}
	// interface Locals {}
	// interface PageData {}
	// interface Platform {}
	interface RoomPreview {
		name: string;
		playersCount: number;
		gameModeName: string;
		id: string;
	}
	interface AppState {
		websocket: import('$lib/websocket').WebSocketClient | null;
		roomId: string | null;
		isPlaying: boolean;
		gameState: GameState | null;
	}
	interface GameState {
		players: Player[];
		apples: Vector2[];
		gridSize: number;
	}
	interface Player{		
		name: string;
		color: string;
		snakeTail: Vector2[];
	}
	
	interface Vector2 {
		x: number;
		y: number;
	}
}
