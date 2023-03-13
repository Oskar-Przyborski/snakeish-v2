<script>
	// @ts-ignore
	import { GoogleAnalytics } from '@beyonk/svelte-google-analytics';
</script>

{#if import.meta.env.PROD}
	<GoogleAnalytics properties={['G-FDJP074REL']} />
{/if}
