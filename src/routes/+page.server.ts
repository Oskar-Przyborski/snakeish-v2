import serverRequest from '$lib/backend_request';
import type { PageServerLoad } from './$types';

export const load = (async (context) => {
	return await serverRequest<App.RoomPreview[]>(context.fetch, '/get-rooms')
}) satisfies PageServerLoad;
