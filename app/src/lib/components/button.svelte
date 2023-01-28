<script lang="ts">
	import { goto } from '$app/navigation';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	export let href: string | undefined | null = null;
	export let disabled: boolean = false;

	const click = () => {
		if (href != null) goto(href);
		else dispatch('click');
	};
</script>

<button class="btn" class:disabled on:click={click} tabindex={disabled ? -1 : 0}>
	<slot />
</button>

<style lang="scss">
	.btn {
		white-space: nowrap;
		padding: 0.5em 1.1em;

		background-color: transparent;
		border-radius: 0.4em;
		border: 3px solid #d62246;

		font-size: 1.05em;
		font-family: 'DM Sans', sans-serif;
		color: white;
		text-align: center;
		text-decoration: none;

		display: inline-block;
		cursor: pointer;
		transition: background-color 0.2s, box-shadow 0.2s;
		&:not(.disabled):active,
		&:not(.disabled):hover {
			box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.217);
			border: 3px solid #d62246;
			background-color: #d62246;
		}

		&.disabled {
			border: 3px solid #808080;
			color: #808080;
			cursor: default;
		}
	}
</style>
