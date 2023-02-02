<script lang="ts">
	export let players: { name: string; color: string }[];
</script>

<div class="leaderboard">
	<h1>Leaderboard</h1>
	<ul>
		{#each players as player}
			<li>{player.name}</li>
		{/each}
	</ul>
</div>
