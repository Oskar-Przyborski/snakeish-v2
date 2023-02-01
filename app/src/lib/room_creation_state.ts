import { writable } from 'svelte/store';

const defaultStateCreator = () => {
	return {
		roomName: '',
		configName: 'classic-casual',
		pinEnabled: false,
		pin: ['', '', '', '']
	};
};

export const state = writable<App.RoomCreationState>(defaultStateCreator());

export const resetState = () => state.set(defaultStateCreator());
