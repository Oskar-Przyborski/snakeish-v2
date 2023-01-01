import type { PageLoad } from './$types';
import backendRequest from '$lib/backend_request';

export const load = (async ({ params, fetch }) => {
	const { isOnline, data } = await backendRequest<App.RoomPreview>(
		fetch,
		`/get-room?id=${params.id}`
	);
	if (!isOnline) return { isOnline: false };

	if (data?.id)
		return {
			isOnline: true,
			...data
		};
}) satisfies PageLoad;
