<script lang="ts">
	import Footer from '$lib/components/footer.svelte';
	import Sidenav from './sidenav.svelte';
	import { page } from '$app/stores';
</script>

<div class="app">
	<Sidenav route={$page.url.pathname} />
	<div class="content-wrapper">
		<div class="content">
			<main><slot /></main>
			<Footer />
		</div>
		<div class="ad-wrapper">
			<div class="example-ad">ad</div>
		</div>
	</div>
</div>

<style lang="scss">
	.app {
		height: 100%;
		display: grid;
		grid-template-columns: 220px 1fr;
	}
	.content-wrapper {
		width: 100%;
		overflow-y: scroll;
		display: grid;
		grid-template-columns: 1fr 15%;
	}
	.ad-wrapper {
		display: grid;
		place-items: center;
		position: sticky;
		top: 0;
		height: 100vh;
		.example-ad {
			height: 500px;
			width: 150px;
			background-color: #666;
			color: white;
			display: grid;
			place-items: center;
		}
	}
</style>
