<script lang="ts">
	import { goto } from '$app/navigation';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	export let href: string | undefined | null = null;
	export let disabled: boolean = false;
	export let outline: boolean = false;

	const click = () => {
		if (disabled) return;
		if (href != null) goto(href);
		else dispatch('click');
	};
</script>

<button class="btn" class:disabled class:outline on:click={click} tabindex={disabled ? -1 : 0}>
	<slot />
</button>

<style lang="scss">
	.btn {
		white-space: nowrap;
		padding: 0.5em 1.1em;

		background-color: #d62246;
		border-radius: 0.4em;
		border: 3px solid #d62246;

		font-size: 1.05em;
		font-family: 'DM Sans', sans-serif;
		color: white;
		text-align: center;
		text-decoration: none;

		display: inline-block;
		cursor: pointer;
		transition: background-color 0.1s, box-shadow 0.2s, border 0.1s;

		&:not(.disabled):hover {
			box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.2);
			background-color: var(--hover-red);
			border: 3px solid var(--hover-red);
		}

		&.disabled {
			background-color: #808080;
			border: 3px solid #808080;
			color: #cdcdcd;
			cursor: default;
		}

		&.outline {
			background-color: transparent;
			&:not(.disabled):hover {
				box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.2);
			}

			&.disabled {
				border: 3px solid #808080;
				color: #808080;
				cursor: default;
			}
		}
	}
</style>
