import backendRequest from '$lib/backend_request';
import type { PageServerLoad } from './$types';

export const load = (async (context) => {
	const { isOnline, data } = await backendRequest<{
		rooms: App.RoomPreview[];
		remainingRooms: number;
	}>(context.fetch, '/get-suggested-rooms');

	return { isOnline, rooms: data?.rooms ?? [], remainingRooms: data?.remainingRooms ?? 0 };
}) satisfies PageServerLoad;
