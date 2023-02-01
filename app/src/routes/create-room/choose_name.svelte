<script lang="ts">
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import Panel from '$lib/components/panel.svelte';
	import { state } from '$lib/room_creation_state';
</script>

<Panel>
	<div class="input-wrapper">
		<TextInput placeholder="Room name" bind:value={$state.roomName} />
	</div>
</Panel>

<style lang="scss">
	.input-wrapper {
		font-size: 1.2rem;
	}
</style>
