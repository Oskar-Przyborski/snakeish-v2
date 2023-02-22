<script lang="ts">
	import type { PageData } from './$types';
	import RoomPreview from '../room_preview.svelte';
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
		if (search != '') params['s'] = search;;
		
		data.rooms = await fetchJson('rooms', { params });
	};
</script>

<div class="browse-all">
	<Panel margin>
		<div class="top-section">
			<h1>Browse Rooms</h1>
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
	</Panel>
	<Panel margin>
		{#if data.rooms.length > 0}
			<div class="list">
				{#each data.rooms as room}
					<div class="preview-wrapper">
						<RoomPreview {room} />
					</div>
				{/each}
			</div>
		{:else}
			There is no room that matches your filters. But you can create one!
		{/if}
	</Panel>
</div>

<style lang="scss">
	.top-section {
		display: flex;
		align-items: center;
		justify-content: space-between;
		h1 {
			margin: 0;
		}

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
	.list {
		display: grid;
		grid-template-columns: 50% 50%;
		margin: 1rem -0.5rem;

		.preview-wrapper {
			padding: 0.5rem;
		}
	}
</style>
