<script lang="ts">
	import { page } from '$app/stores';
	import Button from '$lib/components/buttons/button.svelte';
</script>

{#if $page.error}
	{#if $page.error.code == 'ROOM_NOT_FOUND'}
		<h1>404</h1>
		<pre>Room not found</pre>
	{:else}
		<h1>{$page.status}</h1>
		<pre>{$page.error.code}</pre>
		<pre>{$page.error.message}</pre>
	{/if}
{/if}
<Button outline href="/">Back to home</Button>
