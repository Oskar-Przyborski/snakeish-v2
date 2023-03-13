import { WebSocketClient } from '$lib/websocket';
import { get, writable } from 'svelte/store';
import type { ClassicGameState, JoinError, JoinSuccessType, PageState } from './types';
//@ts-ignore
import { ga } from '@beyonk/svelte-google-analytics';

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
	localStorage.setItem('player-name', name);
	localStorage.setItem('player-color', color);
	get(store).websocket?.sendMessage('request-join', { name, color, pin });
};

const joinSuccess = (data: JoinSuccessType) => {
	ga.addEvent('join_game', {
		playerName: data.name,
		modeTag: "classic",
	});

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
