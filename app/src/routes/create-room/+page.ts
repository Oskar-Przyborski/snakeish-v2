import type { PageLoad } from './$types';
import { PUBLIC_API_URL } from '$env/static/public';
import { goto } from '$app/navigation';

export const load = (({ fetch }) => {
	async function createRoom(data: { roomName: string; configName: string | null; pin: string[] }) {
		const resp = await fetch(PUBLIC_API_URL + '/create-classic-room', {
			body: JSON.stringify(data),
			method: 'POST',
			headers: {
				"content-type": "application/json"
			}
		});
		if (resp.status != 200) {
			//TODO handle error
			throw Error('Error while creating room');
		}
		const json = await resp.json()
		goto('/room/'+json.id);
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
