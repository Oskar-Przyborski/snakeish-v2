<script lang="ts">
	import type { PageServerData } from './$types';
	import AvailableRooms from './available_rooms.svelte';
	import GameDescription from './game_description.svelte';

	export let data: PageServerData;
	const rooms = data.data ? data.data : [];
</script>

<div class="snakeish-logo">
	<img alt="Snakeish Logo" src="/logo-long-white.png" width="310" height="60" />
</div>
{#if data.isOnline}
	<AvailableRooms {rooms} />
{:else}
	Server is offline.
{/if}
<GameDescription />

<style>
	.snakeish-logo {
		display: grid;
		place-items: center;
		padding: 0.5em 0;
	}
</style>
