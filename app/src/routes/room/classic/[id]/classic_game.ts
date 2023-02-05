import { WebSocketClient } from '$lib/websocket';
import { get, writable } from 'svelte/store';

export const store = writable<PageState>({
	roomId: null,
	isPlaying: false,
	gameState: null,
	playerId: null,
	websocket: null
});

export const connectRoomWebsocket = (roomId: string) => {
	const updateGameState = (newState: ClassicGameState) => {
		store.update((state) => {
			state.gameState = newState;
			return state;
		});
	};

	store.update((state) => {
		state.roomId = roomId;

		const wsConnection = new WebSocketClient(`/connect/classic/${state.roomId}`);
		state.websocket = wsConnection;

		wsConnection.addListener<ClassicGameState>('game-update', updateGameState);
		wsConnection.addListener<JoinSuccessType>('join-success', joinSuccess);

		return state;
	});
};

export const requestJoin = (name: string, color: string) => {
	get(store).websocket?.sendMessage('request-join', { name, color });
};

const joinSuccess = (data: JoinSuccessType) => {
	store.update((state) => {
		state.isPlaying = true;
		state.playerId = data.playerId;
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

export const leaveGame = () => {
	store.update((state) => {
		state.websocket?.sendMessage('request-leave', {});
		state.isPlaying = false;

		return state;
	});
};

export interface PageState {
	isPlaying: boolean;
	websocket: import('$lib/websocket').WebSocketClient | null;
	gameState: ClassicGameState | null;
	playerId: string | null;
	roomId: string | null;
}
export interface ClassicGameState {
	players: Player[];
	apples: App.Vector2[];
	gridSize: number;
}
export interface Player {
	id: string;
	name: string;
	color: string;
	snakeTail: App.Vector2[];
	score: number;
}

interface JoinSuccessType {
	playerId: string;
	name: string;
	color: string;
}
