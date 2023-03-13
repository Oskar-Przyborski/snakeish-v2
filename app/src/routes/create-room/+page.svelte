<script lang="ts">
	import type { PageData } from './$types';
	import Panel from '$lib/components/panel.svelte';
	import Divider from '$lib/components/divider.svelte';
	import ChooseName from '$lib/components/create-room/choose_name.svelte';
	import ChooseMode from '$lib/components/create-room/choose_mode.svelte';
	import StepItem from '$lib/components/create-room/step_item.svelte';
	import ChoosePin from '$lib/components/create-room/choose_pin.svelte';
	import { resetState, store } from '$lib/components/create-room/room_creation_state';
	import { goto } from '$app/navigation';
	import { fetchJson } from '$lib/fetchJson';
	import modes from '$lib/modes';
	//@ts-ignore
	import { ga } from '@beyonk/svelte-google-analytics';

	export let data: PageData;
	const createRoom = async () => {
		if ($store.configName == null) return;
		const { id, modeTag, modeName, name } = await fetchJson<App.RoomPreview>('/rooms/create', {
			fetcher: fetch,
			method: 'POST',
			body: {
				roomName: $store.roomName,
				modeName: modes.get($store.configName)?.title,
				modeTag: modes.get($store.configName)?.tag,
				pin: $store.pinEnabled ? $store.pin : null
			}
		});

		ga.addEvent('create_room', {
			modeName,
			modeTag,
			roomName: name
		});

		await goto(`/room/${modeTag}/${id}`, { replaceState: true });
		resetState();
	};

	let currStep = 0;
	const prevStep = () => currStep--;
	const nextStep = () => currStep++;
</script>

<svelte:head>
	<title>Snakeish - Create Room</title>
	<meta name="title" content="Snakeish - Create Room" />
	<meta name="description" content="Create Room in Snakeish - online multiplayer snake game" />
</svelte:head>

<h1 class="header">Create</h1>
<Panel margin>
	<div class="container">
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
</Panel>

<style lang="scss">
	.header {
		margin: 1rem;
	}
	.container {
		display: grid;
		grid-template-columns: 25% max-content 1fr;
		gap: 1rem;
	}

	.step-view {
		padding: 1rem 1rem;
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
