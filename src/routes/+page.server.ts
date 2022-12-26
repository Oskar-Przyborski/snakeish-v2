import { PUBLIC_API_URL } from '$env/static/public';
import type { PageServerLoad } from './$types';

export const load = (async (context) => {
	try {
		const resp = await context.fetch(PUBLIC_API_URL + '/get-rooms');
		const data: App.RoomPreview[] = await resp.json();
		return {
			isOnline: true,
			rooms: data
		};
	} catch {
		return {
			isOnline: false,
			rooms: []
		};
	}
}) satisfies PageServerLoad;
