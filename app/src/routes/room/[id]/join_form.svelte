<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import TextInput from '$lib/components/text_input.svelte';

	const dispatcher = createEventDispatcher();

	let name: string;
	let color: string;

	const submit = () => {
		dispatcher('join', { name, color });
	};
</script>

<div class="join-form">
	<h1>Join Game</h1>
	<TextInput bind:value={name}>Name</TextInput>
	<TextInput bind:value={color}>Color</TextInput>
	<Button on:click={submit}>Join</Button>
</div>

<style lang="scss">
	.join-form {
		display: flex;
		flex-direction: column;
	}
</style>
