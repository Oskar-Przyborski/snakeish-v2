<script lang="ts">
	import type { PageData } from './$types';
	import { onMount, onDestroy } from 'svelte';
	import {
		connectRoomWebsocket,
		leaveRoom,
		joinPlayer,
		appStateStore,
		leaveGame
	} from '$lib/app_state';
	import JoinForm from './join_form.svelte';
	import Leaderboard from './leaderboard.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import GameRenderer from './game_renderer.svelte';
	import Panel from '$lib/components/panel.svelte';

	export let data: PageData;

	let appState: App.AppState;
	const unsubscribe = appStateStore.subscribe((newState) => (appState = newState));

	const bindings = new Map<string, string>([
		['KeyW', 'up'],
		['ArrowUp', 'up'],
		['KeyS', 'down'],
		['ArrowDown', 'down'],
		['KeyD', 'right'],
		['ArrowRight', 'right'],
		['KeyA', 'left'],
		['ArrowLeft', 'left']
	]);

	onMount(() => {
		if (data.id == null) return;
		connectRoomWebsocket(data.id);
	});

	onDestroy(() => {
		unsubscribe();
		leaveRoom();
	});

	function onJoin(event: CustomEvent<{ name: string; color: string }>) {
		joinPlayer(event.detail.name, event.detail.color);
	}

	function onKeyDown(event: KeyboardEvent) {
		if (!appState.isPlaying) return;

		const direction = bindings.get(event.code);
		if (direction == null) return;

		event.preventDefault();
		appState.websocket?.sendMessage('change-direction', { direction });
	}
</script>

{#if data.id != null && appState.gameState != null}
	<div class="wrapper">
		<Panel>
			<div class="game-renderer">
				<GameRenderer />
			</div>
		</Panel>
		<div class="sidebar">
			<Panel>
				<div class="titles">
					<h1>{data.name}</h1>
					<div class="mode">
						<span class="tag">{data.gameModeTag.toUpperCase()}</span>
						<h2>{data.gameModeName}</h2>
					</div>
				</div>
			</Panel>
			<Panel>
				{#if !appState.isPlaying}
					<JoinForm on:join={onJoin} />
				{:else}
					<Leaderboard players={appState.gameState.players} on:leave={leaveGame} />
				{/if}
			</Panel>
		</div>
	</div>
{:else}
	<p style="text-align:center">Loading</p>
{/if}

<svelte:window on:keydown={onKeyDown} />

<style lang="scss">
	.wrapper {
		display: grid;
		grid-template-columns: 3fr 2fr;
		gap: 1rem;
		h1 {
			margin: 0rem 0;
		}
	}
	.game-renderer {
		display: grid;
		place-items: center;
	}
	.sidebar {
		display: grid;
		grid-template-rows: max-content 1fr;
		gap: 1rem;

		.titles {
			display: flex;
			flex-flow: row nowrap;
			align-items: center;
			justify-content: space-between;
			.mode {
				.tag {
					font-size: 0.75rem;
					font-weight: bold;
					background-color: #0000002b;
					width: max-content;
					padding: 0.3rem;
				}
				h2 {
					margin: 0.3rem 0;
				}
			}
		}
	}
</style>
