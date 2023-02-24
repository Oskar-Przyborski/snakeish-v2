<script lang="ts">
	import Button from '$lib/components/buttons/button.svelte';
	import Icon from '@iconify/svelte';

	export let room: App.RoomPreview;
</script>

<div class="room-preview-container">
	<div class="left">
		<div class="name">{room.name}</div>
		<div class="game-mode">
			<div class="tag">{room.modeTag.toUpperCase()}</div>
			<div class="name">{room.modeName}</div>
		</div>
	</div>
	<div class="right">
		<div class="icon">
			<Icon icon="mdi:account-supervisor" inline />
			{room.players}/{room.maxPlayers}
		</div>
		<div class="join">
			{#if room.pinEnabled}
				<span class="icon"><Icon icon="mdi:lock-outline" inline /></span>
			{/if}
			<Button href={`/room/${room.modeTag}/${room.id}`}>
				{#if room.players == room.maxPlayers}
					Watch
				{:else}
					Join
				{/if}
			</Button>
		</div>
	</div>
</div>

<style lang="scss">
	.room-preview-container {
		margin: 1rem;
		padding: 1.2rem 1.5rem;

		color: var(--text);
		text-decoration: none;

		background-color: var(--surface);
		border-radius: 12px;
		border: 2px solid var(--outline);

		display: flex;
		justify-content: space-between;
		gap: 1.3rem;

		.left {
			display: flex;
			align-items: center;
			justify-content: start;
			flex-flow: row wrap;
			gap: 1.1rem;
		}
		.right {
			display: flex;
			align-items: center;
			justify-content: end;
			flex-flow: row wrap;
			gap: 1.1rem;
		}
		.name {
			font-size: 1.6rem;

			overflow: hidden;
			white-space: nowrap;
			text-overflow: ellipsis;
		}
		.game-mode {
			display: flex;
			flex-direction: row;
			gap: 0.5rem;
			// align-items: center;
			.tag {
				width: max-content;
				padding: 3px 4px;
				background: var(--outline);
			}
			.name {
				font-size: 1.3rem;
			}
		}
		.icon {
			font-size: 1.3rem;
		}

		.join {
			display: flex;
			align-items: center;
			gap: 0.5rem;
		}
	}
</style>
