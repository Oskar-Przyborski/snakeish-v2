<script lang="ts">
	import Button from '$lib/components/buttons/button.svelte';
	import Panel from '$lib/components/panel.svelte';
	import RoomPreview from '$lib/components/room_preview.svelte';
	import Icon from '@iconify/svelte';

	export let rooms: App.RoomPreview[];
	export let remainingRooms: number;
</script>

{#if rooms.length != 0}
	<div>
		{#each rooms as room}
			<RoomPreview {room} />
		{/each}
	</div>
	{#if remainingRooms != 0}
		<div class="remaining-rooms">
			...and {remainingRooms} more
			<Button href="/browse" outline>
				<Icon icon="material-symbols:search-rounded" inline /> Browse All
			</Button>
		</div>
	{/if}
{:else}
	<Panel margin>
		<div class="no-rooms-error">
			<h2>There is no any room here 😬</h2>
			<p>Don't worry, you can <a href="/create-room">create one</a>!</p>
		</div>
	</Panel>
{/if}

<style lang="scss">
	.no-rooms-error {
		text-align: center;
		color: var(--text-light);
	}
	.remaining-rooms {
		margin: 1rem;
		font-size: 1.1rem;

		display: flex;
		justify-content: flex-end;
		align-items: center;
		gap: 1rem;
	}
</style>
