<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import type { Writable } from 'svelte/store';
	import type { PageState, Player } from './classic_game';

	export let store: Writable<PageState>;

	let unsubscribe = () => {};
	onMount(() => {
		unsubscribe = store.subscribe(updateState);
	});
	onDestroy(unsubscribe);

	let canvas: HTMLCanvasElement;

	interface CanvasConfig {
		canvas: HTMLCanvasElement;
		ctx: CanvasRenderingContext2D;
		gridSize: number;
		cellSize: number;
	}

	function updateState(state: PageState) {
		if (canvas == null) return;
		if (state.gameState == null) return;

		const ctx = canvas.getContext('2d');
		if (ctx == null) return;
		const cellSize = canvas.width / state.gameState.gridSize;

		const config: CanvasConfig = {
			canvas,
			ctx,
			cellSize,
			gridSize: state.gameState.gridSize
		};

		clearGrid(config);
		drawGrid(config);
		drawApples(config, state.gameState.apples);
		drawSnakes(config, state.gameState.players);
	}

	function clearGrid(config: CanvasConfig) {
		config.ctx.clearRect(0, 0, config.canvas.width, config.canvas.height);
	}

	function drawGrid(config: CanvasConfig) {
		config.ctx.beginPath();
		for (let x = 0; x < config.gridSize + 1; x++) {
			config.ctx.moveTo(x * config.cellSize, 0);
			config.ctx.lineTo(x * config.cellSize, canvas.height);
		}
		for (let y = 0; y < config.gridSize + 1; y++) {
			config.ctx.moveTo(0, y * config.cellSize);
			config.ctx.lineTo(canvas.width, y * config.cellSize);
		}
		config.ctx.strokeStyle = 'rgba(0,0,0,1)';
		config.ctx.stroke();
		config.ctx.closePath();
	}

	function drawApples(config: CanvasConfig, apples: App.Vector2[]) {
		for (let i = 0; i < apples.length; i++) {
			const apple = apples[i];
			let color = 'red';
			// if (apple.value === 3) {
			// 	color = 'gold';
			// }
			drawRectShape(
				config,
				{
					x: apple.x * config.cellSize + config.cellSize * 0.125,
					y: apple.y * config.cellSize + config.cellSize * 0.125
				},
				{
					x: config.cellSize * 0.75,
					y: config.cellSize * 0.75
				},
				color
			);
		}
	}

	function drawRectShape(
		config: CanvasConfig,
		position: App.Vector2,
		size: App.Vector2,
		color: string
	) {
		config.ctx.fillStyle = color;
		config.ctx.fillRect(position.x, position.y, size.x, size.y);
	}

	function drawSnakes(config: CanvasConfig, players: Player[]) {
		if (players.length == 0) return;

		for (let i = 0; i < players.length; i++) {
			const player = players[i];
			DrawPlayerSnake(config, player);
		}
	}

	function DrawPlayerSnake(config: CanvasConfig, player: Player) {
		const snake = player.snakeTail;
		if (snake.length == 0) return;

		for (let i = 0; i < snake.length; i++) {
			const currentElement = snake[i];
			const previousElement = i < snake.length - 1 ? snake[i + 1] : null;
			const nextElement = i > 0 ? snake[i - 1] : null;
			drawSnakeElement(config, previousElement, currentElement, nextElement, player.color);
		}
		drawText(config, player.name, snake[0]);
	}

	function drawSnakeElement(
		config: CanvasConfig,
		prev: App.Vector2 | null,
		curr: App.Vector2,
		next: App.Vector2 | null,
		color: string
	) {
		const prevDiff = prev != null ? { x: prev.x - curr.x, y: prev.y - curr.y } : null;
		const nextDiff = next != null ? { x: next.x - curr.x, y: next.y - curr.y } : null;
		const currentElementTopLeftPos = {
			x: curr.x * config.cellSize,
			y: curr.y * config.cellSize
		};
		switch (JSON.stringify({ prev: prevDiff, next: nextDiff })) {
			case JSON.stringify({ prev: { x: -1, y: 0 }, next: { x: 1, y: 0 } }):
			case JSON.stringify({ prev: { x: 1, y: 0 }, next: { x: -1, y: 0 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: 0, y: 1 }, next: { x: 0, y: -1 } }):
			case JSON.stringify({ prev: { x: 0, y: -1 }, next: { x: 0, y: 1 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: 0, y: 1 }, next: { x: 1, y: 0 } }):
			case JSON.stringify({ prev: { x: 1, y: 0 }, next: { x: 0, y: 1 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.875
					},
					color
				);
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.875,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: -1, y: 0 }, next: { x: 0, y: 1 } }):
			case JSON.stringify({ prev: { x: 0, y: 1 }, next: { x: -1, y: 0 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.875
					},
					color
				);
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.875,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: 0, y: -1 }, next: { x: 1, y: 0 } }):
			case JSON.stringify({ prev: { x: 1, y: 0 }, next: { x: 0, y: -1 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.875
					},
					color
				);
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.875,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: 0, y: -1 }, next: { x: -1, y: 0 } }):
			case JSON.stringify({ prev: { x: -1, y: 0 }, next: { x: 0, y: -1 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.875
					},
					color
				);
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.875,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: null, next: { x: 1, y: 0 } }):
			case JSON.stringify({ prev: { x: 0, y: 0 }, next: { x: 1, y: 0 } }):
			case JSON.stringify({ prev: { x: 1, y: 0 }, next: null }):
			case JSON.stringify({ prev: { x: 1, y: 0 }, next: { x: 0, y: 0 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.875,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: -1, y: 0 }, next: null }):
			case JSON.stringify({ prev: { x: -1, y: 0 }, next: { x: 0, y: 0 } }):
			case JSON.stringify({ prev: null, next: { x: -1, y: 0 } }):
			case JSON.stringify({ prev: { x: 0, y: 0 }, next: { x: -1, y: 0 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.875,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
			case JSON.stringify({ prev: null, next: { x: 0, y: -1 } }):
			case JSON.stringify({ prev: { x: 0, y: 0 }, next: { x: 0, y: -1 } }):
			case JSON.stringify({ prev: { x: 0, y: -1 }, next: null }):
			case JSON.stringify({ prev: { x: 0, y: -1 }, next: { x: 0, y: 0 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.875
					},
					color
				);
				break;
			case JSON.stringify({ prev: { x: 0, y: 1 }, next: null }):
			case JSON.stringify({ prev: { x: 0, y: 1 }, next: { x: 0, y: 0 } }):
			case JSON.stringify({ prev: null, next: { x: 0, y: 1 } }):
			case JSON.stringify({ prev: { x: 0, y: 0 }, next: { x: 0, y: 1 } }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.875
					},
					color
				);
				break;
			case JSON.stringify({ prev: null, next: null }):
				drawRectShape(
					config,
					{
						x: currentElementTopLeftPos.x + config.cellSize * 0.125,
						y: currentElementTopLeftPos.y + config.cellSize * 0.125
					},
					{
						x: config.cellSize * 0.75,
						y: config.cellSize * 0.75
					},
					color
				);
				break;
		}
	}

	function drawText(config: CanvasConfig, text: string, position: App.Vector2) {
		config.ctx.fillStyle = 'white';
		config.ctx.font = config.cellSize / 1.8 + 'px DM Sans';
		config.ctx.textAlign = 'center';

		//fill text centered vertically and horizontally
		const xPx = position.x * config.cellSize + config.cellSize / 2;
		const yPx = position.y * config.cellSize + config.cellSize / 2 + config.cellSize / 6;
		config.ctx.fillText(text, xPx, yPx);
	}
</script>

<canvas
	id="game-render"
	bind:this={canvas}
	width={550}
	height={550}
	style="image-rendering: crisp-edges;"
/>
