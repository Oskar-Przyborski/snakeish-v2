<script lang="ts">
	import type { PageData } from './$types';
	import Panel from '$lib/components/panel.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import TextButton from '$lib/components/buttons/text_button.svelte';
	import Divider from '$lib/components/divider.svelte';
	import ChooseName from './choose_name.svelte';
	import ChooseMode from './choose_mode.svelte';
	import StepItem from './step_item.svelte';
	import ChoosePin from './choose_pin.svelte';

	// const createRoom = () => data.createRoom({ roomName, configName });

	export let data: PageData;
	
	let currStep = 0;
	const prevStep = () => currStep--;
	const nextStep = () => currStep++;

	let ctaBtnState: 'continue' | 'skip' | 'disabled' = 'continue';
</script>

<Panel full>
	<div class="container">
		<div class="top-header">
			<h1>Creating room</h1>
			<Divider />
		</div>
		<div class="bottom">
			<div class="navigation-panel">
				{#each data.steps as step, idx}
					<StepItem {step} active={currStep == idx} prevActive={idx < currStep} />
				{/each}
			</div>
			<Divider vertical />
			<div class="step-view">
				<div class="top-headers">
					<div>Step {currStep + 1}/{data.steps.length}</div>
					<h1>{data.steps[currStep].stepView.title}</h1>
					<div>{data.steps[currStep].stepView.description}</div>
				</div>
				{#if currStep == 0}
					<ChooseName on:continue={nextStep}/>
				{:else if currStep == 1}
					<ChooseMode on:continue={nextStep} />
				{:else if currStep == 2}
					<ChoosePin on:continue={nextStep} />
				{/if}
				<div class="continue-section">
					{#if currStep != 0}
						<TextButton on:click={prevStep}>Back</TextButton>
					{/if}
					<Button on:click={nextStep} disabled={ctaBtnState == 'disabled'}>
						{#if ctaBtnState == 'skip'}
							Skip
						{:else}
							Continue
						{/if}
					</Button>
				</div>
			</div>
		</div>
	</div>
</Panel>

<style lang="scss">
	.container {
		height: 100%;
		display: grid;
		grid-template-rows: max-content 1fr;
		.top-header {
			h1 {
				margin: 0;
				margin-bottom: 1rem;
			}
		}
		.bottom {
			display: grid;
			grid-template-columns: 25% max-content 1fr;
			gap: 1rem;
		}
	}

	.step-view {
		padding: 2rem 1rem;
		padding-bottom: 1rem;
		.top-headers {
			margin-bottom: 2rem;
			h1 {
				font-size: 1.6rem;
				margin: 0.4rem 0;
			}
		}
		.continue-section {
			margin-top: 1.5rem;
			display: flex;
			justify-content: flex-end;
			gap: 1rem;
			align-items: center;
		}
	}
</style>
