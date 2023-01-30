<script lang="ts">
	import Icon from '@iconify/svelte';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	const onClick = () => dispatch('click');

	export let step: { name: string; description: string };
	export let prevActive: boolean;
	export let active: boolean;
</script>

<svelte:element
	this={prevActive ? 'button' : 'div'}
	class="step-item"
	class:active
	class:prev-active={prevActive}
	on:click={onClick}
	on:keypress={onClick}
>
	<div class="icon">
		<Icon icon="eva:checkmark-circle-2-outline" />
	</div>
	<div class="name">
		<div class="title">{step.name}</div>
		<div class="description">{step.description}</div>
	</div>
</svelte:element>

<style lang="scss">
	.step-item {
		all: unset;
		margin: 1.5rem 0;

		display: flex;
		gap: 1rem;
		align-items: center;

		color: #fffb;

		.name {
			.title {
				font-weight: bold;
				font-size: 1.5rem;
			}
		}
		.icon {
			font-size: 2.3rem;
			display: grid;
			place-items: center;
		}

		&.active {
			color: #fff;
			.icon {
				color: var(--red);
			}
		}

		&.prev-active {
			cursor: pointer;
			.icon {
				color: #d62246b0;
			}

			&:hover {
				color: #fff9;
				.icon {
					color: #d6224690;
				}
			}
			&:focus-visible {
				outline-color: rgb(255, 255, 255);
				outline-width: 0.8px;
				outline-style: auto;
				outline-offset: 0px;
			}
		}
	}
</style>
