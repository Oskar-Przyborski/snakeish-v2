import type { Vector2 } from '$lib/vector';
import { WebSocketClient } from '$lib/websocket';
import { get, writable } from 'svelte/store';

export const store = writable<PageState>({
	roomId: null,
	isPlaying: false,
	gameState: null,
	playerId: null,
	websocket: null,
	errors: {
		name: undefined,
		pin: undefined
	}
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
		wsConnection.addListener<JoinError>('join-error', joinError);

		return state;
	});
};

export const requestJoin = (name: string, color: string, pin: (number | null)[]) => {
	get(store).websocket?.sendMessage('request-join', { name, color, pin });
};

const joinSuccess = (data: JoinSuccessType) => {
	store.update((state) => {
		state.isPlaying = true;
		state.playerId = data.playerId;
		state.errors = { name: undefined, pin: undefined };
		return state;
	});
};
const joinError = (data: JoinError) => {
	switch (data.error) {
		case 'incorrect-pin':
			store.update((state) => {
				state.errors.pin = 'Incorrect PIN code!';
				return state;
			});
			break;
		case 'player-name-already-taken':
			store.update((state) => {
				state.errors.name = 'Name is already taken!';
				return state;
			});
			break;
	}
};

export const leaveRoom = () => {
	store.update((state) => {
		state.websocket?.ws.close();

		state.roomId = null;
		state.isPlaying = false;
		state.websocket = null;
		state.gameState = null;

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
	errors: {
		name: string | undefined;
		pin: string | undefined;
	};
}
export interface ClassicGameState {
	players: Player[];
	apples: Vector2[];
	gridSize: number;
}
export interface Player {
	id: string;
	name: string;
	color: string;
	snakeTail: Vector2[];
	score: number;
}

interface JoinSuccessType {
	playerId: string;
	name: string;
	color: string;
}
interface JoinError {
	error: string;
}
