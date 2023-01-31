import type { PageLoad } from './$types';
import backendRequest from '$lib/backend_request';

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
				iconFull: 'mdi:pencil'
			},
			{
				name: 'Mode',
				description: 'Select game mode, you most like!',
				iconEmpty: 'mdi:gamepad-variant-outline',
				iconFull: 'mdi:gamepad-variant'
			},
			{
				name: 'PIN Code',
				description: 'Protect your room against uninvited players',
				iconEmpty: 'mdi:lock-outline',
				iconFull: 'mdi:lock'
			}
		]
	};
}) satisfies PageLoad;
