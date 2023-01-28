<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	export let href: string | undefined | null = null;
	export let disabled: boolean = false;
</script>

{#if href != null}
	<a {href}>
		<div class="btn" class:disabled>
			<slot />
		</div>
	</a>
{:else}
	<button class:disabled class="btn" on:click={() => dispatch('click')}>
		<slot />
	</button>
{/if}

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
