<script lang="ts">
	import Icon from '@iconify/svelte';
	import Button from '$lib/components/button.svelte';

	export let room: App.RoomPreview;
</script>

<div class="room-preview-container">
	<div class="top">
		<div class="name">{room.name}</div>
		<Button href={`/room/${room.id}`}><Icon icon="ion:enter" inline /> Join</Button>
	</div>
	<div class="bottom">
		<div class="game-mode">
			<div class="tag">{room.gameModeTag.toUpperCase()}</div>
			<div class="name">{room.gameModeName}</div>
		</div>
		<div class="players">
			<Icon icon="ic:sharp-people-alt" inline />
			{room.players}/{room.maxPlayers}
		</div>
	</div>
</div>

<style lang="scss">
	.room-preview-container {
		border: 2.5px solid var(--red);
		border-radius: 12px;
		filter: drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25));

		padding: 1.2rem 1.5rem;
		display: flex;
		flex-direction: column;
		gap: 1.3rem;

		.top {
			display: flex;
			justify-content: space-between;
			align-items: center;
			gap: 1rem;
			font-size: 1.2rem;

			.name {
				font-size: 1.8rem;
				overflow: hidden;
				white-space: nowrap;
				text-overflow: ellipsis;
			}
		}
		.bottom {
			display: flex;
			justify-content: space-between;
			align-items: flex-end;

			.game-mode {
				display: flex;
				flex-direction: column;
				gap: 0.2rem;
				// align-items: center;
				.tag {
					width: max-content;
					padding: 3px 4px;
					background: rgba(0, 0, 0, 0.3);
				}
				.name {
					font-size: 1.65rem;
				}
			}

			.players {
				font-size: 1.6rem;
			}
		}
	}
</style>
