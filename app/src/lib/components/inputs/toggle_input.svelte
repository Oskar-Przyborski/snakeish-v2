<script lang="ts">
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher()

	export let value: boolean;

	const toggle = () => {
		value = !value;
		dispatch("change")
	};
</script>

<button class="switch" class:checked={value} on:click={toggle}>
	<div class="slider" />
</button>

<style lang="scss">
	.switch {
		all: unset;
		margin: 0 0.25rem;
		cursor: pointer;
		vertical-align: baseline;
		width: 2em;
		border-radius: 2em;
		padding: 0.25em;
		background-color: #555;
		transition: background-color 0.1s ease-in-out;

		&:hover {
			background-color: #444;
		}

		.slider {
			width: 1em;
			height: 1em;
			background-color: white;
			border-radius: 50%;
			transition: margin-left 0.1s ease-in-out;
		}

		&.checked {
			background-color: var(--red);
			&:hover {
				background-color: var(--hover-red);
			}
			.slider {
				margin-left: 1em;
			}
		}
		&:focus-visible {
			outline: auto;
		}
	}
</style>
