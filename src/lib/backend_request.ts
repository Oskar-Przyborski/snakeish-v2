import { PUBLIC_API_URL } from '$env/static/public';

export default async function backendRequest<T>(
	fetch: (input: RequestInfo | URL, init?: RequestInit | undefined) => Promise<Response>,
	endpoint: string,
	body: unknown = null
): Promise<{ isOnline: boolean; data: T | undefined }> {
	const withBody = body != null;

	const resBody = withBody ? JSON.stringify(body) : null;
	const resMethod = withBody ? 'POST' : 'GET';
	const resHeaders = {
		'Content-Type': withBody ? 'application/json' : 'application/x-www-form-urlencoded'
	};

	try {
		const controller = new AbortController();
		setTimeout(() => controller.abort(), 5000);
		const resp = await fetch(PUBLIC_API_URL + endpoint, {
			signal: controller.signal,
			method: resMethod,
			body: resBody,
			headers: resHeaders
		});
		return {
			isOnline: true,
			data: (await resp.json()) satisfies T
		};
	} catch {
		return {
			isOnline: false,
			data: undefined
		};
	}
}
