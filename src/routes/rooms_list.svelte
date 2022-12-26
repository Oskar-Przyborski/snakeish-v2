<script lang="ts">
	import Button from '$lib/components/button.svelte';
	import Icon from '@iconify/svelte';

	export let rooms: App.RoomPreview[];
</script>

<div class="rooms-list">
	{#if rooms.length > 0}
		{#each rooms as room}
			<div class="list-item">
				<div class="left-side">
					<div class="name"><b>{room.name}</b></div>
					<div class="props">
						<div><Icon icon="fa6-solid:users" inline={true} /> {room.playersCount}</div>
						<div><Icon icon="fa6-solid:gamepad" inline={true} /> {room.gameModeName}</div>
					</div>
				</div>
				<Button>Join</Button>
			</div>
		{/each}
	{:else}
		<div class="empty-err"><b>😭 List is empty...</b></div>
	{/if}
</div>

<style lang="scss">
	.rooms-list {
		padding: 1em 0;
		display: flex;
		flex-direction: column;
		gap: 1em;

		.list-item {
			display: flex;
			justify-content: space-between;
			align-items: center;

			padding-left: 1.1em;
			padding-right: 0.6em;

			background-color: #d62246;
			border-top-right-radius: 1em;
			border-bottom-right-radius: 1em;

			.left-side {
				font-size: 1.25em;
				display: flex;
				flex-direction: row;
				gap: 1.5em;
				.name {
					max-width: 300px;
					text-overflow: ellipsis;
					overflow: hidden;
					white-space: nowrap;
				}
				.props {
					gap: 1.5em;
					display: flex;
					flex-direction: row;
				}
			}
		}
		.empty-err {
			text-align: center;
			font-size: 1.5em;
		}
	}
	@media (max-width: 992px) {
		.rooms-list {
			.list-item {
				font-size: 0.9em;
				.left-side {
					.name {
						max-width: 150px;
					}
				}
			}
		}
	}
	@media (max-width: 575px) {
		.rooms-list {
			.list-item {
				.left-side {
					flex-direction: column;
					justify-content: space-evenly;
					gap: 0.35em;
					padding: 0.6em 0;
					.name {
						max-width: 180px;
					}
				}
			}
		}
	}
</style>
