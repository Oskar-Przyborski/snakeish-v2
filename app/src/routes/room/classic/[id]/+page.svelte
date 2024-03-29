<script lang="ts">
	import type { PageData } from './$types';
	import {
		connectRoomWebsocket,
		leaveRoom,
		store
	} from '$lib/components/room/classic/classic_game';
	import { onMount, onDestroy } from 'svelte';
	import JoinForm from '$lib/components/room/classic/join_form.svelte';
	import Leaderboard from '$lib/components/room/classic/leaderboard.svelte';
	import GameRenderer from '$lib/components/room/classic/game_renderer.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import Watching from '$lib/components/watching.svelte';

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

	function onKeyDown(event: KeyboardEvent) {
		if (!$store.isPlaying) return;

		const direction = bindings.get(event.code);
		if (direction == null) return;

		event.preventDefault();
		$store.websocket?.sendMessage('change-direction', { direction });
	}
</script>

<svelte:head>
	<title>Play Snakeish - {data.name}, {data.modeName} mode</title>
	<meta name="title" content={`Play Snakeish - ${data.name}, ${data.modeName} mode`} />
	<meta
		name="description"
		content="Play snake online mulitplayer with your friends on classic mode!"
	/>
</svelte:head>

{#if data.id != null && $store.gameState != null}
	<div class="wrapper">
		<Panel>
			<div class="game-renderer">
				<GameRenderer {store} />
			</div>
		</Panel>
		<div class="sidebar">
			<div class="titles">
				<h1>{data.name}</h1>
				<div class="mode">
					<span class="tag">{data.modeTag.toUpperCase()}</span>
					<h2>{data.modeName}</h2>
				</div>
			</div>
			<Panel>
				{#if !$store.isPlaying}
					<JoinForm pinEnabled={data.pinEnabled} />
				{:else}
					<Leaderboard />
				{/if}
			</Panel>
			{#if $store.gameState?.watching}
				<Watching count={$store.gameState?.watching} />
			{/if}
		</div>
	</div>
{:else}
	<p style="text-align:center">Loading</p>
{/if}
<div class="ad-wrapper">
	<div class="example-ad">ad</div>
</div>

<svelte:window on:keydown={onKeyDown} />

<style lang="scss">
	.wrapper {
		padding: 1rem;
		display: grid;
		grid-template-columns: 3fr 2fr;
		gap: 1rem;
	}
	@media (max-width: 996px) {
		.wrapper{
			grid-template-columns: 1fr;
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
			padding: 1rem;
			display: flex;
			flex-flow: row nowrap;
			align-items: center;
			justify-content: space-between;
			gap: 1rem;

			h1 {
				margin: 0;
			}
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
	.ad-wrapper {
		display: grid;
		place-items: center;
		.example-ad {
			height: 100px;
			width: 700px;
			background-color: #666;
			color: white;
			display: grid;
			place-items: center;
		}
	}
</style>
