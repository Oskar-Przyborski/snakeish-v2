<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	export let href: string | undefined | null = null;

</script>

{#if href != null}
	<a {href}>
		<button>
			<slot />
		</button>
	</a>
{:else}
	<button on:click={() => dispatch('click')}>
		<slot />
	</button>
{/if}

<style lang="scss">
	button {
		margin: 0.6em;
		padding: 0.5em 1.1em;

		background-color: transparent;
		border-radius: 0.4em;
		border: 3px solid #8be8cb;

		font-size: 1.05em;
		font-family: 'DM Sans', sans-serif;
		color: white;
		text-align: center;
		text-decoration: none;

		display: inline-block;
		cursor: pointer;
		transition: background-color 0.2s, color 0.2s;
		&:active,
		&:hover {
			background-color: #8be8cb;
			color: black;
		}
	}
</style>
