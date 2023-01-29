<script lang="ts">
	import Button from '$lib/components/button.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import RefreshButton from './refresh_button.svelte';
	import RoomPreview from './room_preview.svelte';

	export let rooms: App.RoomPreview[];
	export let remainingRooms: number = 0;

	export let onRefresh: () => void;
</script>

<Panel>
	<div class="header">
		<h1>Suggested Rooms</h1>
		<RefreshButton callback={onRefresh} />
	</div>
	{#if rooms.length != 0}
		<div class="rooms-grid">
			{#each rooms as room}
				<div class="preview-wrapper">
					<RoomPreview {room} />
				</div>
			{/each}
		</div>
		{#if remainingRooms != 0}
			<div class="remaining-rooms">
				...and {remainingRooms} more <Button
					><Icon icon="ic:sharp-search" inline /> Browse All</Button
				>
			</div>
		{/if}
	{:else}
		<div class="no-rooms-error">
			<h2>There is no any room here!</h2>
			<p>Don't worry, you can create one!</p>
		</div>
	{/if}
</Panel>

<style lang="scss">
	.no-rooms-error {
		text-align: center;
	}
	h1 {
		margin: 0;
	}

	.header {
		display: flex;
		justify-content: space-between;
	}
	.rooms-grid {
		display: grid;
		grid-template-columns: 50% 50%;
		margin: 1rem -0.5rem;
		.preview-wrapper {
			padding: 0.5rem;
		}
	}
	.remaining-rooms {
		font-size: 1.1rem;
		display: flex;
		justify-content: flex-end;
		gap: 1rem;
		align-items: center;
	}
</style>
