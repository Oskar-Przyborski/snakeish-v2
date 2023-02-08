<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import TextInput from '$lib/components/inputs/text_input.svelte';

	const dispatcher = createEventDispatcher();

	//TODO PIN code support
	let name: string;
	let color: string;

	const submit = () => {
		dispatcher('join', { name, color });
	};
</script>

<div class="join-form">
	<h1>Join Game</h1>
	<div class="form">
		<TextInput bind:value={name}>Name</TextInput>
		<TextInput bind:value={color}>Color</TextInput>
	</div>
	<div class="join-btn">
		<Button on:click={submit}>Join</Button>
	</div>
</div>

<style lang="scss">
	.join-form {
		h1 {
			margin: 0;
		}
		.form {
			margin: 2rem 0;
		}
		.join-btn {
			float: right;
		}
	}
</style>
