<script lang="ts">
	import Panel from '$lib/components/panel.svelte';
	import RefreshButton from './refresh_button.svelte';
	import RoomPreview from './room_preview.svelte';

	export let rooms: App.RoomPreview[];
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
</style>
