import backendRequest, { responseToMap } from '$lib/backend_request';
import type { PageServerLoad } from './$types';

export const load = (async (context) => {
	const { data, isOnline } = await responseToMap<App.RoomPreview>(backendRequest(context.fetch, '/get-rooms'));
	const rooms: App.RoomPreview[] = [];
	if (data != null) {
		data.forEach((v) => rooms.push(v));
	}
	
	return {
		isOnline,
		data: rooms
	}
}) satisfies PageServerLoad;
