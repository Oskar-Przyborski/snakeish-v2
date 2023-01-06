import { PUBLIC_WS_URL } from '$env/static/public';
import { writable } from 'svelte/store';

export const websocketClient = writable<WebSocketClient | null>(null);

export class WebSocketClient {
	constructor(path: string) {
		this.listeners = new Map<string, (payload: unknown) => void>();

		this.ws = new WebSocket(PUBLIC_WS_URL + path);
		this.ws.onmessage = (ev) => {
			const { event, payload } = decodeEventData(ev.data);
			const callback = this.listeners.get(event);
			if (callback != null) {
				callback(payload);
			}
		};
	}

	ws: WebSocket;
	listeners: Map<string, (payload: unknown) => void>;

	addListener<T>(event: string, callback: (payload: T) => void) {
		this.listeners.set(event, callback as (payload: unknown) => void);
	}

	sendMessage(event: string, payload: unknown) {
		this.ws.send(encodeEventData(event, payload));
	}
}

function decodeEventData(data: string) {
	data = data.trim();
	let event = '';
	let msgData = '';
	for (let i = 0; i < data.length; i++) {
		const char = data[i];
		if (char == ';') {
			event = data.slice(0, i);
			msgData = data.slice(i+1);
			break;
		}
	}

	const payload = JSON.parse(msgData);

	return {
		event,
		payload
	};
}

function encodeEventData(eventName: string, payload: unknown): string {
	return eventName + ';' + JSON.stringify(payload);
}
