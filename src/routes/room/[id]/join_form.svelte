<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/button.svelte';

	const dispatcher = createEventDispatcher();

	let name: string;
	let color: string;

	const submit = () => {
		dispatcher('join', { name, color });
	};
</script>

<div class="join-form">
	<input type="text" placeholder="Name" bind:value={name} />
	<input type="text" placeholder="Color" bind:value={color} />
	<Button on:click={submit}>Join</Button>
</div>
