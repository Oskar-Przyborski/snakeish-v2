<script lang="ts">
	import Button from '$lib/components/button.svelte';
	import RefreshButton from './refresh_button.svelte';
	import Icon from '@iconify/svelte';
	import RoomPreview from './RoomPreview.svelte';

	export let rooms: App.RoomPreview[];
	export let onRefresh: () => void;
</script>

<div class="available-rooms">
	<div class="header">
		<h1>Available Rooms</h1>
		<RefreshButton callback={onRefresh} />
	</div>
	<div class="rooms-grid">
		{#each rooms as room}
			<div class="preview-wrapper">
				<RoomPreview {room} />
			</div>
		{/each}
	</div>
	<div class="below-grid">
		<Button><Icon icon="ic:sharp-plus" inline /> Create New Room</Button>
		<div>
			<span>... and 34 more</span>
			<Button><Icon icon="ic:baseline-search" inline /> Browse All</Button>
		</div>
	</div>
</div>

<style lang="scss">
	.available-rooms {
		margin-bottom: 2rem;
	}

	.header {
		display: flex;
		justify-content: space-between;
	}
	.rooms-grid {
		display: grid;
		grid-template-columns: 50% 50%;
		margin: -1rem;
		.preview-wrapper {
			padding: 1rem;
		}
	}

	.below-grid {
		margin-top: 1rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
		font-size: 1.2rem;

		span {
			margin-right: 10px;
		}
	}
</style>
