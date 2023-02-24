<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
	import Icon from '@iconify/svelte';

	export let checked = false;
	const toggle = () => {
		checked = !checked;
		dispatch('change');
	};
</script>

<div class="checkbox" class:checked on:click={toggle} on:keypress={toggle}>
	<button>
		<span class="icon"><Icon icon="mdi:check-bold" /></span>
	</button>
	<div class="label">
		<slot />
	</div>
</div>

<style lang="scss">
	.checkbox {
		cursor: pointer;

		display: flex;
		flex-flow: row nowrap;
		gap: 0.5rem;

		button {
			all: unset;
			width: 1rem;
			height: 1rem;
			border: 2px solid #ccc;
			border-radius: 0.3rem;
			background-color: #fff;

			.icon {
				color: #fff;
				display: none;
			}
		}

		&.checked {
			button {
				border-color: var(--red);
				background-color: var(--red);
				position: relative;
				.icon {
					display: block;
				}
			}
		}
		.label {
			user-select: none;
			font-size: 1.1rem;
		}
	}
</style>
