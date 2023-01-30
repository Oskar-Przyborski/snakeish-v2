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
				description: 'Enhance your room with your own name!'
			},
			{
				name: 'Mode',
				description: 'Select game mode, you most like!'
			},
			{
				name: 'PIN Code',
				description: 'Protect your room against uninvited players'
			},
			{
				name: 'Summary',
				description: 'Make sure of everything and jump right into the game!'
			}
		]
	};
}) satisfies PageLoad;
