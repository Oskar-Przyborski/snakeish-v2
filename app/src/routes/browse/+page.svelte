<script lang="ts">
	import type { PageData } from './$types';
	import RoomPreview from '../../lib/components/room_preview.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import ToggleInput from '$lib/components/inputs/toggle_input.svelte';
	import ModeDropdown from './mode_dropdown.svelte';
	import { fetchJson } from '$lib/fetchJson';

	export let data: PageData;

	let publicOnly: boolean = false;
	let search: string = '';

	const refresh = async () => {
		const params: any = {};
		if (publicOnly) params['public'] = 1;
		if (search != '') params['s'] = search;

		data.rooms = await fetchJson('rooms', { params });
	};
</script>

<h1 class="header">Browse</h1>
<div class="browse-all">
	<div class="top-section">
		<div class="search">
			<TextInput placeholder="Search" bind:value={search} on:change={refresh} />
			<span style="font-size: 1.4rem;">
				<Icon icon="mdi:magnify" inline />
			</span>
		</div>
		<div class="btns">
			<ModeDropdown />
			<div class="public-only">
				<span style="font-size: 1.4rem;">
					<Icon icon="mdi:lock-open-variant-outline" inline />
				</span>
				<ToggleInput bind:value={publicOnly} on:change={refresh} />
			</div>
		</div>
	</div>
	{#each data.rooms as room}
		<RoomPreview {room} />
	{:else}
		<Panel>
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
				// flex-flow: column nowrap;
				align-items: center;
				gap: 0.4rem;
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
