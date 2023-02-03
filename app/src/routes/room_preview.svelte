<script lang="ts">
	import Icon from '@iconify/svelte';

	export let room: App.RoomPreview;
</script>

<a class="room-preview-container" href={`/room/${room.id}`}>
	<div class="top">
		<div class="name">{room.name}</div>
		<span class="enter-icon"><Icon icon="mdi:location-enter" inline /></span>
	</div>
	<div class="bottom">
		<div class="game-mode">
			<div class="tag">{room.modeTag.toUpperCase()}</div>
			<div class="name">{room.modeName}</div>
		</div>
		<div class="players">
			<Icon icon="mdi:account-supervisor" inline />
			{room.players}/{room.maxPlayers}
		</div>
	</div>
</a>

<style lang="scss">
	.room-preview-container {
		color: white;
		text-decoration: none;
		border-radius: 12px;
		padding: 1.2rem 1.5rem;
		border: 2.5px solid var(--red);
		// background-color: #ffffff1a;

		display: flex;
		flex-direction: column;
		gap: 1.3rem;

		cursor: pointer;
		transition: box-shadow 0.1s;
		&:hover,
		&:focus {
			box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.25);
			.top .name {
				text-decoration: underline;
			}
		}

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
			.enter-icon {
				color: white;
				font-size: 2rem;
			}
		}
		.bottom {
			display: flex;
			justify-content: space-between;
			align-items: flex-end;

			.game-mode {
				display: flex;
				flex-direction: row;
				gap: 0.5rem;
				// align-items: center;
				.tag {
					width: max-content;
					padding: 3px 4px;
					background: rgba(0, 0, 0, 0.3);
				}
				.name {
					font-size: 1.3rem;
				}
			}

			.players {
				font-size: 1.3rem;
			}
		}
	}
</style>
