<script lang="ts">
	import type { PageData } from './$types';
	import Panel from '$lib/components/panel.svelte';
	import Divider from '$lib/components/divider.svelte';
	import ChooseName from './choose_name.svelte';
	import ChooseMode from './choose_mode.svelte';
	import StepItem from './step_item.svelte';
	import ChoosePin from './choose_pin.svelte';
	import { store } from '$lib/room_creation_state';

	export let data: PageData;
	const createRoom = () =>{
		data.createRoom($store);
	}

	let currStep = 0;
	const prevStep = () => currStep--;
	const nextStep = () => currStep++;
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
				<div class="divider">
					<Divider />
				</div>
				{#if currStep == 0}
					<ChooseName on:continue={nextStep} />
				{:else if currStep == 1}
					<ChooseMode on:continue={nextStep} on:back={prevStep} />
				{:else if currStep == 2}
					<ChoosePin on:create-room={createRoom} on:back={prevStep} />
				{/if}
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
			margin-bottom: 1.5rem;
			h1 {
				font-size: 1.6rem;
				margin: 0.4rem 0;
			}
		}
		.divider {
			margin: 1.5rem 0;
		}
	}
</style>
