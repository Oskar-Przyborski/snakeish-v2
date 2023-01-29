<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let title: string, description: string, tag: string;
	export let selected: boolean = false;

	const dispatcher = createEventDispatcher();
</script>

<button
	class="gamemode-card"
	class:selected
	on:click={() => dispatcher('select')}
	on:keydown={() => dispatcher('select')}
>
	<div class="tag">{tag.toUpperCase()}</div>
	<h2>{title}</h2>
	<div class="description">{description}</div>
</button>

<style lang="scss">
	.gamemode-card {
		all: unset;
		border: 3px solid #ffffff;
		border-radius: 0.8rem;
		padding: 1rem;
		transition: border 0.1s ease-in-out, box-shadow 0.1s ease-in-out;
		cursor: pointer;

		.tag {
			font-size: 0.75rem;
			font-weight: bold;
			background-color: #0000002b;
			width: max-content;
			padding: 0.3rem;
		}

		h2 {
			margin: 0.3rem 0;
		}

		&:hover {
			box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.384);
		}
		&.selected {
			border: 3px solid #d62246;
			box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.384);
		}
	}
</style>
