<script lang="ts">
	import { page } from '$app/stores';
</script>

{#if $page.error}
	{#if $page.error.message == 'Internal Error'}
		<h1>Server offline</h1>
		<p>Server is currently offline.</p>
		<ol>
			<li>Wait a few seconds and try reloading the page, maybe the server is still waking up.</li>
			<li>If it doesn't work, the server is probably completely down</li>
		</ol>
	{:else}
		<h1>{$page.status}</h1>
		<pre>{$page.error.code}</pre>
		<pre>{$page.error.message}</pre>
	{/if}
{/if}
