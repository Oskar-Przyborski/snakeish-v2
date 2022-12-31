import { PUBLIC_API_URL } from '$env/static/public';

export default async function serverRequest<T>(
	fetch: (input: RequestInfo | URL, init?: RequestInit | undefined) => Promise<Response>,
	endpoint: string
): Promise<{ isOnline: boolean; data: T | undefined }> {
	try {
		const controller = new AbortController();
		setTimeout(() => controller.abort(), 5000);
		const resp = await fetch(PUBLIC_API_URL + endpoint, { signal: controller.signal });
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
