<script lang="ts">
	import type { PageData } from './$types';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import ChooseName from './choose_name.svelte';
	import StepIndicator from '$lib/components/step_indicator.svelte';
	import ChooseMode from './choose_mode.svelte';

	const createRoom = () => data.createRoom({ roomName, configName });

	export let data: PageData;
	let roomName: string;
	let configName = 'classic-casual';

	let currStep = 0;
	const nextStep = () => currStep++;
	const goToStep = (idx: number) => (currStep = idx < currStep ? idx : currStep); //can go ONLY back
</script>

<Panel full>
	<div class="container">
		<div class="navigation-panel">
			<h1>Creating Room</h1>
			{#each data.steps as step, idx}
				<div
					class="step-item"
					class:active={currStep == idx}
					class:prev-active={idx < currStep}
					on:click={() => goToStep(idx)}
					on:keypress={() => goToStep(idx)}
				>
					<div class="icon">
						<Icon icon="eva:checkmark-circle-2-outline" />
					</div>
					<div class="name">
						<div class="title">{step.name}</div>
						<div class="description">{step.description}</div>
					</div>
				</div>
			{/each}
		</div>
		<div class="step-view">
			{#if currStep == 0}
				<ChooseName bind:value={roomName} on:continue={nextStep} />
			{:else if currStep == 1}
				<ChooseMode
					gameModes={data.gameModes}
					bind:selectedConfig={configName}
					on:continue={nextStep}
				/>
			{:else if currStep == 2}
				<div>Password</div>
			{:else if currStep == 3}
				<div>Summary</div>
			{/if}
			<StepIndicator steps={data.steps.length} current={currStep} />
		</div>
	</div>
</Panel>

<style lang="scss">
	.container {
		height: 100%;
		display: grid;
		grid-template-columns: 25% 1fr;
		gap: 1rem;
	}

	.navigation-panel {
		h1 {
			margin: 0;
			margin-bottom: 2rem;
		}
		.step-item {
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
			}
		}
	}

	.step-view {
		display: grid;
		grid-template-rows: 1fr max-content;
		gap: 1rem;
		padding: 0 1rem;
	}
</style>
