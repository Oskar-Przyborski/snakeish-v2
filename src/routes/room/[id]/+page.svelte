<script lang="ts">
	import type { PageData } from './$types';
	import { websocketClient, WebSocketClient } from '$lib/websocket';
	import { onMount, onDestroy } from 'svelte';
	import { get } from 'svelte/store';

	onMount(() => {
		const wsConnection = new WebSocketClient('/ws-connect-room?id=' + data.id);
		wsConnection.addListener<{ name: string; color: string; id: string }[]>(
			'players-update',
			(payload) => {
				data.playersCount = payload.length;
			}
		);
		websocketClient.set(wsConnection);
	});

	onDestroy(() => {
		get(websocketClient)?.ws.close();
	});

	export let data: PageData;
</script>

<div>
	{#if data.id != null}
		room: {data.name}<br />
		gameMode: {data.gameModeName}<br />
		playersCount: {data.playersCount}<br />
	{:else}
		room not found
	{/if}
</div>
