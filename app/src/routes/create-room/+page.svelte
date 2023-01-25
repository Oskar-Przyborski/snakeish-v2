<script lang="ts">
	import backendRequest from '$lib/backend_request';
	import Button from '$lib/components/button.svelte';
	import TextInput from '$lib/components/text_input.svelte';


    let roomName: string;
	async function createRoom() {
		await backendRequest(fetch, '/create-classic-room', {
			roomName,
		});
	}
    </script>

<h1>Create rooom</h1>
<TextInput bind:value={roomName}>Room name</TextInput>
<Button on:click={createRoom}>Create</Button>
