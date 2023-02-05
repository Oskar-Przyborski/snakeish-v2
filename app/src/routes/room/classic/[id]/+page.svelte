<script lang="ts">
	import type { PageData } from './$types';
	import { connectRoomWebsocket, requestJoin, leaveRoom, store, leaveGame } from './classic_game';
	import { onMount, onDestroy } from 'svelte';
	import JoinForm from './join_form.svelte';
	import Leaderboard from './leaderboard.svelte';
	import GameRenderer from './game_renderer.svelte';
	import Panel from '$lib/components/panel.svelte';

	export let data: PageData;

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
		leaveRoom();
	});

	function onJoin(event: CustomEvent<{ name: string; color: string }>) {
		requestJoin(event.detail.name, event.detail.color);
	}

	function onKeyDown(event: KeyboardEvent) {
		if (!$store.isPlaying) return;

		const direction = bindings.get(event.code);
		if (direction == null) return;

		event.preventDefault();
		$store.websocket?.sendMessage('change-direction', { direction });
	}
</script>

{#if data.id != null && $store.gameState != null}
	<div class="wrapper">
		<Panel>
			<div class="game-renderer">
				<GameRenderer {store} />
			</div>
		</Panel>
		<div class="sidebar">
			<Panel>
				<div class="titles">
					<h1>{data.name}</h1>
					<div class="mode">
						<span class="tag">{data.modeTag.toUpperCase()}</span>
						<h2>{data.modeName}</h2>
					</div>
				</div>
			</Panel>
			<Panel>
				{#if !$store.isPlaying}
					<JoinForm on:join={onJoin} />
				{:else}
					<Leaderboard players={$store.gameState.players} on:leave={leaveGame} />
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
