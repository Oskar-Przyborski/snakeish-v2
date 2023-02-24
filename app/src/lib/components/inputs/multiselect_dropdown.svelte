<script lang="ts">
	import { browser } from '$app/environment';
	import { onDestroy, onMount } from 'svelte';
	import Icon from '@iconify/svelte';
	import CheckboxInput from '$lib/components/inputs/checkbox_input.svelte';

	export let values: { name: string; checked: boolean }[];
	let expanded = false;
	let element: HTMLDivElement;

	function onClickOutside(e: MouseEvent) {
		if (e.target != null && !element.contains(e.target as Node)) {
			expanded = false;
		}
	}

	onMount(() => {
		if (browser) window.addEventListener('click', onClickOutside);
	});

	onDestroy(() => {
		if (browser) window.removeEventListener('click', onClickOutside);
	});

	$: checked = values.filter((i) => i.checked).length;
</script>

<div class="multiselect-dropdown">
	<div class="name">
		<slot />
	</div>
	<div class="mode-tag" bind:this={element}>
		<button on:click={() => (expanded = !expanded)}>
			{#if checked == 0}
				Select
			{:else}
				Selected {checked}
			{/if}
			<Icon icon={`mdi:menu-${expanded ? 'up' : 'down'}`} inline />
		</button>
		<div class="content" class:expanded>
			<ul>
				{#each values as item}
					<li>
						<CheckboxInput bind:checked={item.checked}>{item.name}</CheckboxInput>
					</li>
				{/each}
			</ul>
		</div>
	</div>
</div>

<style lang="scss">
	.multiselect-dropdown{
		display: flex;
		flex-flow: row wrap;
		align-items: center;
		gap: 0.5rem;
	}
	.mode-tag {
		position: relative;

		button {
			all: unset;
			cursor: pointer;
			user-select: none;
			font-size: 1.1rem;
			padding: 0.5rem 1rem;
			background-color: var(--surface);
			border: 2px solid var(--outline);
			border-radius: 0.5rem;
		}

		.content {
			display: none;
			&.expanded {
				display: block;
			}

			position: absolute;
			right: 0;
			// min-width: 100%;

			background: var(--surface);
			border: 2px solid var(--outline);
			border-radius: 0.5rem;
			padding: 0.4rem 0;
			margin-top: 0.2rem;

			ul {
				margin: 0;
				padding: 0;
				list-style: none;
				li {
					padding: 0.2rem 0.9rem;
					margin: 0.1rem 0;
				}
			}
		}
	}
</style>
