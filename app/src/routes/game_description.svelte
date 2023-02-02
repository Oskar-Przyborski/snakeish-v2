<script>
	import Panel from '$lib/components/panel.svelte';
</script>

<Panel>
	<h1>Snakeish</h1>
	<p>
		Online multiplayer traditional snake game that contains two game modes - classic and
		battle-royale.
	</p>
	<h2>Classic</h2>
	<p>
		Classic mode is a classic snake game. Use your keyboard arrows or WASD to move your snake.<br />
		Eat apples to grow you snake, but don&apos;t hit walls, yourself or other snakes, or you will die.<br
		/>
		Gameplay is highly customizable for example size of the grid or speed of the game.
	</p>
	<h2>Battle-royale</h2>
	<p>
		Battle-royale mode is something different and even more challenging.<br />
		The goal is to be the last snake alive, but it&apos;s not that easy.<br />
		At the beggining of the game you wait for the lobby to be full. Then the game starts.<br />
		You need to eat apples to grow you snake, because every 15 seconds shortest snake in the game will
		die.<br />
		Also every 15 seconds the map will shrink and if you go inside that dead zone your snake will be
		shortening.<br />
		If you go outside the map, hit yourself or other snakes, you will die.<br />
		<br />
	</p>
</Panel>
