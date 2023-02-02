import type { PageLoad } from './$types';
import backendRequest from '$lib/backend_request';
import { store } from '$lib/room_creation_state';
import { get } from 'svelte/store';

export const load = (({ fetch }) => {
	async function createRoom(data: { roomName: string; configName: string }) {
		await backendRequest(fetch, '/create-classic-room', data);
	}

	return {
		createRoom,
		steps: [
			{
				name: 'Name',
				description: 'Enhance your room with your own name!',
				iconEmpty: 'mdi:pencil-outline',
				iconFull: 'mdi:pencil',
				stepView: {
					title: "Let's start with name",
					description: 'What would you like to name your room? Let it be short and descriptive'
				},
				getCtaBtnState: async () => {
					const state = get(store);
					if (state.roomName == '') return 'disabled';

					const { data } = await backendRequest<{ available: boolean }>(
						fetch,
						`/room-name-available?name=${state.roomName}`
					);

					if (data?.available) return 'continue';
					return 'disabled';
				}
			},
			{
				name: 'Mode',
				description: 'Select game mode, you most like!',
				iconEmpty: 'mdi:gamepad-variant-outline',
				iconFull: 'mdi:gamepad-variant',
				stepView: {
					title: 'What mode fits you the best?',
					description: 'There are many modes to choose from. Choose your favorite!'
				},
				getCtaBtnState: async () => {
					return 'continue';
				}
			},
			{
				name: 'PIN Code',
				description: 'Protect your room against uninvited players',
				iconEmpty: 'mdi:lock-outline',
				iconFull: 'mdi:lock',
				stepView: {
					title: 'Want to protect your room?',
					description: 'Players will have to enter the PIN code before entering the game.'
				},
				getCtaBtnState: async () => {
					return 'continue';
				}
			}
		]
	};
}) satisfies PageLoad;
