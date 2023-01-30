<script lang="ts">
	import type { PageData } from './$types';
	import Panel from '$lib/components/panel.svelte';
	import ChooseName from './choose_name.svelte';
	import StepIndicator from '$lib/components/step_indicator.svelte';
	import ChooseMode from './choose_mode.svelte';
	import StepItem from './step_item.svelte';
	import ChoosePin from './choose_pin.svelte';

	const createRoom = () => data.createRoom({ roomName, configName });

	export let data: PageData;
	let roomName: string;
	let configName = 'classic-casual';
	let pin = ['', '', '', ''];

	let currStep = 0;
	const nextStep = () => currStep++;
	const goToStep = (idx: number) => (currStep = idx < currStep ? idx : currStep); //can go ONLY back
</script>

<Panel full>
	<div class="container">
		<div class="navigation-panel">
			<h1>Creating Room</h1>
			{#each data.steps as step, idx}
				<StepItem
					on:click={() => goToStep(idx)}
					{step}
					active={currStep == idx}
					prevActive={idx < currStep}
				/>
			{/each}
		</div>
		<div class="step-view">
			{#if currStep == 0}
				<ChooseName bind:value={roomName} on:continue={nextStep} />
			{:else if currStep == 1}
				<ChooseMode
					modes={data.modes}
					bind:selectedConfig={configName}
					on:continue={nextStep}
				/>
			{:else if currStep == 2}
				<ChoosePin on:continue={nextStep} bind:pin />
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
	}

	.step-view {
		display: grid;
		grid-template-rows: 1fr max-content;
		gap: 1rem;
		padding: 0 1rem;
	}
</style>
