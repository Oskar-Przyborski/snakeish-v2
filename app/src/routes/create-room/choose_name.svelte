<script lang="ts">
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import { state } from '$lib/room_creation_state';
</script>

<div class="choose-name">
	<div class="input-wrapper">
		<TextInput placeholder="Room name" bind:value={$state.roomName} />
	</div>
</div>

<style lang="scss">
	.input-wrapper {
		font-size: 1.4rem;
	}
</style>
