import type { PageLoad } from './$types';
import { goto } from '$app/navigation';
import { resetState } from '$lib/room_creation_state';
import { fetchJson } from '$lib/fetchJson';
import modes from '$lib/modes';

export const load = (({ fetch }) => {
	async function createRoom(data: { roomName: string; configName: string | null; pin: string[] }) {
		const { id, modeTag } = await fetchJson<App.RoomPreview>('/rooms/create', {
			fetcher: fetch,
			method: 'POST',
			body: {
				roomName: data.roomName,
				modeName: data.configName ? modes.get(data.configName)?.title : null,
				modeTag: data.configName ? modes.get(data.configName)?.tag : null,
				pin: data.pin
			}
		});

		await goto(`/room/${modeTag}/${id}`, { replaceState: true });
		resetState();
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
				}
			}
		]
	};
}) satisfies PageLoad;
