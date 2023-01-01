import backendRequest from '$lib/backend_request';
import type { PageServerLoad } from './$types';

export const load = (async (context) => {
	return await backendRequest<App.RoomPreview[]>(context.fetch, '/get-rooms')
}) satisfies PageServerLoad;
