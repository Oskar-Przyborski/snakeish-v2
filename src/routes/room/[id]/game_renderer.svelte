<script lang="ts">
	import { appStateStore } from '$lib/app_state';
	import { onDestroy, onMount } from 'svelte';

	let unsubscribe = () => {};
	onMount(() => {
		unsubscribe = appStateStore.subscribe(updateState);
	});
	onDestroy(unsubscribe);

	let canvas: HTMLCanvasElement;

	interface CanvasConfig {
		canvas: HTMLCanvasElement;
		ctx: CanvasRenderingContext2D;
		gridSize: number;
		cellSize: number;
	}

	function updateState(state: App.AppState) {
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

		drawGrid(config);
		drawApples(config, state.gameState.apples);
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

	function drawApples(config: CanvasConfig, apples: { x: number; y: number }[]) {
		for (let i = 0; i < apples.length; i++) {
			const apple = apples[i];
			let color = 'red';
			// if (apple.value === 3) {
			// 	color = 'gold';
			// }
			DrawRectShape(
				config,
				apple.x * config.cellSize + config.cellSize * 0.125,
				apple.y * config.cellSize + config.cellSize * 0.125,
				config.cellSize * 0.75,
				config.cellSize * 0.75,
				color
			);
		}
	}

	function DrawRectShape(
		config: CanvasConfig,
		x: number,
		y: number,
		width: number,
		height: number,
		color: string
	) {
		config.ctx.fillStyle = color;
		config.ctx.fillRect(x, y, width, height);
	}
</script>

<canvas
	id="game-render"
	bind:this={canvas}
	width={550}
	height={550}
	style="image-rendering: crisp-edges;"
/>
