<script>
	import Container from '$lib/components/container.svelte';
	import Footer from '$lib/components/footer.svelte';
</script>

<div class="app">
	<div class="snakeish-logo">
		<img alt="Snakeish Logo" src="/logo-long-white.png" width="260" />
	</div>
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
		grid-template-rows: max-content 1fr max-content;
	}
	.snakeish-logo {
		text-align: center;
		padding: 0.5em 0;
	}
</style>
