// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	interface Error {
		message: string;
		code: string;
	}
	// interface Locals {}
	// interface PageData {}
	// interface Platform {}
	interface RoomPreview {
		name: string;
		players: number;
		maxPlayers: number;
		modeName: string;
		modeTag: string;
		id: string;
		pinEnabled: boolean;
	}
	interface Vector2 {
		x: number;
		y: number;
	}

	interface RoomCreationState {
		roomName: string;
		configName: string | null;
		pinEnabled: boolean;
		pin: (number | null)[];
	}
}
