import { PUBLIC_API_URL } from '$env/static/public';
import { error, type HttpError } from '@sveltejs/kit';
import destr from 'destr';

interface options<D> {
	method?: string;
	body?: object;
	params?: object;
	default?: D;
	onError?: () => void;
	fetcher?: (input: RequestInfo | URL, init?: RequestInit | undefined) => Promise<Response>;
}

export const fetchJson = async <T>(endpoint: string, options?: options<T>) => {
	const fetcher = options?.fetcher ?? fetch;
	const url = PUBLIC_API_URL + endpoint;
	const method = options?.method ?? 'GET';
	const headers: HeadersInit = {};
	let body = null;

	const handleError = (err: HttpError) => {
		if (options && options.onError) options?.onError();
		else throw err;

		return options?.default as T;
	};

	if (options?.method && options?.body) {
		if (['POST', 'PUT', 'PATCH'].includes(options.method.toUpperCase())) {
			headers['content-type'] = 'application/json';
			body = JSON.stringify(options.body);
		}
	}

	try {
		const controller = new AbortController();
		setTimeout(() => controller.abort(), 5000);
		const response = await fetcher(url, { method, body, headers, signal: controller.signal });

		if (response.headers.get('content-type') != 'application/json') {
			return handleError(error(500, { message: 'invalid response content-type' }));
		}

		const respObj = destr(await response.text());
		if (respObj == undefined) {
			return handleError(error(500, { message: 'can not parse json' }));
		}

		return respObj as T;
	} catch {
		return handleError(error(500, 'server offline'));
	}
};