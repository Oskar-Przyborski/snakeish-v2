<script>
	import Container from '$lib/components/container.svelte';
	import Footer from '$lib/components/footer.svelte';
</script>

<div class="app">
	<Container>
		<slot />
	</Container>
	<Footer />
</div>

<style lang="scss">
	.app {
		padding: 0 1rem;
		height: 100%;
		display: grid;
		grid-template-rows: 1fr max-content;
	}
</style>
