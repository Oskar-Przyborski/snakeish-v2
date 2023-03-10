<script lang="ts">
	import type { PageData } from './$types';
	import RoomPreview from '$lib/components/room_preview.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import ToggleInput from '$lib/components/inputs/toggle_input.svelte';
	import MultiselectDropdown from '../../lib/components/inputs/multiselect_dropdown.svelte';
	import { fetchJson } from '$lib/fetchJson';
	import { browser } from '$app/environment';

	export let data: PageData;

	let publicOnly = false;
	let search = '';
	let modes = [
		{ name: 'Casual', checked: false },
		{ name: 'Huuge', checked: false },
		{ name: 'Eat&Win', checked: false },
	];

	const refresh = async (
		publicOnly: boolean,
		search: string,
		modes: { name: string; checked: boolean }[]
	) => {
		if (!browser) return;
		
		const params: any = {};
		if (publicOnly) params['public'] = 1;
		if (search != '') params['s'] = search;

		const checkedModes = modes.filter((mode) => mode.checked);
		if (checkedModes.length != 0) params['modes'] = checkedModes.map((m) => m.name).join(',');

		data.rooms = await fetchJson('rooms', { params });
	};

	$: refresh(publicOnly, search, modes);
</script>

<svelte:head>
	<title>Snakeish - Browse Rooms</title>
	<meta name="title" content="Snakeish - Browse Rooms" />
	<meta name="description" content="Search for best room in Snakeish - online multiplayer snake game" />
</svelte:head>

<h1 class="header">Browse</h1>
<div class="browse-all">
	<div class="top-section">
		<div class="search">
			<TextInput placeholder="Search" bind:value={search} />
			<span style="font-size: 1.4rem;">
				<Icon icon="material-symbols:search-rounded" inline />
			</span>
		</div>
		<div class="btns">
			<MultiselectDropdown bind:values={modes}>Modes</MultiselectDropdown>
			<div class="public-only">
				<span style="font-size: 1.4rem;">
					<Icon icon="material-symbols:lock-open-outline-rounded" inline />
				</span>
				<ToggleInput bind:value={publicOnly} />
			</div>
		</div>
	</div>
	{#each data.rooms as room}
		<RoomPreview {room} />
	{:else}
		<Panel margin>
			<div class="no-rooms-error">
				There is no room that matches your filters. But you can create one!
			</div>
		</Panel>
	{/each}
</div>

<style lang="scss">
	.header {
		margin: 1rem;
	}
	.top-section {
		margin: 1rem;
		display: flex;
		align-items: center;
		justify-content: space-between;

		.btns {
			display: flex;
			align-items: center;
			gap: 1.5rem;

			.public-only {
				display: flex;
				align-items: center;
				gap: 0.4rem;
				background: var(--surface);
				padding: 0.35rem 0.6rem;
				border: 2px solid var(--outline);
				border-radius: 0.5rem;
			}
		}

		.search {
			display: flex;
			align-items: center;
			gap: 1rem;
		}
	}

	.no-rooms-error {
		text-align: center;
	}
</style>
