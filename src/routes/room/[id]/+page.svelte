<script lang="ts">
	import type { PageData } from "./$types";

    export let data: PageData;
</script>
<div>
    {#if data.id!=null}
        room: {data.name}<br>
        gameMode: {data.gameModeName}<br>
        playersCount: {data.playersCount}<br>
    {:else}
        room not found
    {/if}
</div>