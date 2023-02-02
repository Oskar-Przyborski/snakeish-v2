import { writable } from 'svelte/store';

const defaultStateCreator = () => {
	return {
		roomName: '',
		configName: 'classic-casual',
		pinEnabled: false,
		pin: ['', '', '', '']
	};
};

export const store = writable<App.RoomCreationState>(defaultStateCreator());

export const resetState = () => store.set(defaultStateCreator());
