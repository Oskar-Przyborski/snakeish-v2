<script lang="ts">
	import Button from '$lib/components/button.svelte';
	import backendRequest, { responseToMap } from '$lib/backend_request';
	import RefreshButton from './refresh_button.svelte';
	import RoomsList from './rooms_list.svelte';

	export let rooms: App.RoomPreview[];

	async function refresh() {
		const { data } = await responseToMap<App.RoomPreview>(
			backendRequest<Map<string, App.RoomPreview>>(fetch, '/get-rooms')
		);
		rooms = [];
		if (data != null) {
			data.forEach((v) => rooms.push(v));
		}
	}

	async function createRoom() {
		await backendRequest(fetch, '/create-classic-room', {
			roomName: 'room2',
			speed: 4,
			gridSize: 4,
			collideEnemies: true
		});

		await refresh();
	}
</script>

<div class="available-rooms">
	<div class="rooms">
		<div class="top">
			<h2>Available rooms</h2>
			<RefreshButton callback={refresh} />
		</div>
		<RoomsList {rooms} />
	</div>
	<div class="create-room">
		<h2>Wanna have own room?</h2>
		<Button on:click={createRoom}><b>Create room</b></Button>
	</div>
</div>

<style lang="scss">
	.available-rooms {
		display: grid;
		grid-template-columns: 3fr 1fr;
		overflow: hidden;
		border-radius: 1em;
		margin: 0.5em 0;

		.rooms {
			background-color: #4a525a;
			padding: 1.1em 1.2em;
			padding-left: 0;
			font-size: 1.2em;

			.top {
				display: flex;
				flex-direction: row;
				justify-content: space-between;
				align-items: center;

				padding-left: 1.1em;

				h2 {
					margin: 0;
				}
			}
		}

		.create-room {
			text-align: center;
			background-color: #d62246;
			padding: 1em;
			font-size: 1.2em;
			h2 {
				margin: 0.3em;
			}
		}
	}
	@media (max-width: 992px) {
		.available-rooms {
			font-size: 0.9em;
		}
	}
	@media (max-width: 768px) {
		.available-rooms {
			display: flex;
			flex-direction: column;
		}
	}
	@media (max-width: 575px) {
		.available-rooms {
			display: flex;
			flex-direction: column;
			font-size: 0.9em;
		}
	}
</style>
