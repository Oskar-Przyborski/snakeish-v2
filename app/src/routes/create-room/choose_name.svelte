<script lang="ts">
	import TextInput from '$lib/components/inputs/text_input.svelte';

	export let value: string;
</script>

<div class="choose-name">
	<div class="input-wrapper">
		<TextInput placeholder="Room name" bind:value />
	</div>
</div>

<style lang="scss">
	.input-wrapper {
		font-size: 1.4rem;
	}
</style>
