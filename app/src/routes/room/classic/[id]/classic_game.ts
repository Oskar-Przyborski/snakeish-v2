import { WebSocketClient } from "$lib/websocket";
import { writable } from "svelte/store";

export const store = writable<PageState>({
    roomId: null,
    isPlaying: false,
    gameState: null,
    playerId: null,
    websocket: null
})

export const connectRoomWebsocket = (roomId: string) => {
	const updateGameState = (newState: ClassicGameState) => {
		store.update((state) => {
			state.gameState = newState;
			return state;
		});
	};
	
	store.update((state) => {
		state.roomId = roomId

		const wsConnection = new WebSocketClient(`/connect/classic/${state.roomId}`);
		wsConnection.addListener<ClassicGameState>('game-update', updateGameState);
		state.websocket = wsConnection;

		return state;
	});
};

export const leaveRoom = () => {
	store.update((state) => {
		state.websocket?.ws.close();

		state.roomId = null;
		state.isPlaying = false;
		state.websocket = null;

		return state;
	});
};

export const requestJoin = (name: string, color: string) => {
	store.update((state) => {
		state.websocket?.sendMessage('request-join', { name, color });
		state.isPlaying = true; //TODO this will need to be changed

		return state;
	});
};

export const leaveGame = () => {
	store.update((state) => {
		state.websocket?.sendMessage('request-leave', {});
		state.isPlaying = false;

		return state;
	});
};

export interface PageState {
    isPlaying: boolean;
	websocket: import("$lib/websocket").WebSocketClient | null;
    gameState: ClassicGameState | null;
    playerId: string | null
    roomId: string | null;
}
export interface ClassicGameState {
	players: Player[];
	apples: App.Vector2[];
	gridSize: number;
}
export interface Player {
    id: string
    name: string;
    color: string;
    snakeTail: App.Vector2[];
    score: number
}
