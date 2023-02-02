import { writable } from 'svelte/store';
import { WebSocketClient } from './websocket';

export const appStateStore = writable<App.AppState>({
	roomId: null,
	websocket: null,
	gameState: null,
	isPlaying: false
});

export const connectRoomWebsocket = (id: string) => {
	const updateGameState = (newState: App.GameState) => {
		appStateStore.update((state) => {
			state.gameState = newState;
			return state;
		});
	};

	appStateStore.update((state) => {
		state.roomId = id;
		const wsConnection = new WebSocketClient('/ws-connect-room?id=' + id);
		wsConnection.addListener<App.GameState>('game-update', updateGameState);
		state.websocket = wsConnection;

		return state;
	});
};

export const leaveRoom = () => {
	appStateStore.update((state) => {
		state.roomId = null;
		state.isPlaying = false;
		state.websocket?.ws.close();
		state.websocket = null;

		return state;
	});
};

export const joinPlayer = (name: string, color: string) => {
	appStateStore.update((state) => {
		state.websocket?.sendMessage('join-player', { name, color });
		state.isPlaying = true;

		return state;
	});
};

export const leaveGame = () => {
	appStateStore.update((state) => {
		state.websocket?.sendMessage('leave-game', {});
		state.isPlaying = false;

		return state;
	});
};
