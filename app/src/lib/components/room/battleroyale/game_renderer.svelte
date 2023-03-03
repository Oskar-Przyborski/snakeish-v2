<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { get, type Writable } from 'svelte/store';
	import Konva from 'konva';
	import type { Vector2 } from '$lib/vector';
	import type { ClassicGameState, PageState, Player } from './types';
	import { GameFramesRenderer } from '$lib/renderers/game_frames';
	import { renderSnake } from '$lib/renderers/snake';
	import { renderApple } from '$lib/renderers/apple';

	export let store: Writable<PageState>;
	let unsubscribe = () => {};
	onDestroy(unsubscribe);

	let canvasWrapper: HTMLDivElement;
	let stage: Konva.Stage;
	let gridLayer: Konva.Layer;
	let applesLayer: Konva.Layer;
	let snakesLayer: Konva.Layer;
	let gameFrameRenderer: GameFramesRenderer<ClassicGameState> | null;

	onMount(() => {
		initKonva();
		const pageState = get(store);
		gameFrameRenderer = new GameFramesRenderer<ClassicGameState>(
			50,
			pageState.gameState!.frameTime,
			pageState.gameState!,
			(state: ClassicGameState, moveValue: number) => {
				drawPlayers(state.players, stage.width() / state.gridSize, moveValue);
			}
		);
		unsubscribe = store.subscribe(updateState);
	});

	async function updateState(state: PageState) {
		if (state.gameState == null) return;
		const cellSize = stage.width() / state.gameState.gridSize;

		drawGrid(state.gameState.gridSize, cellSize);
		drawApples(state.gameState.apples, cellSize);
		gameFrameRenderer?.setGameData(state.gameState);
	}

	function initKonva() {
		stage = new Konva.Stage({
			container: canvasWrapper,
			width: 550,
			height: 550
		});

		gridLayer = new Konva.Layer();
		applesLayer = new Konva.Layer();
		snakesLayer = new Konva.Layer();
		stage.add(gridLayer, applesLayer, snakesLayer);
	}

	function drawGrid(gridSize: number, cellSize: number) {
		if (gridLayer.hasChildren()) return;

		for (let y = 0; y < gridSize; y++) {
			for (let x = 0; x < gridSize; x++) {
				const darked = y % 2 == 0 ? x % 2 == 0 : x % 2 == 1;
				const color = darked ? getComputedStyle(document.documentElement).getPropertyValue("--grid-1") : getComputedStyle(document.documentElement).getPropertyValue("--grid-2");
				gridLayer.add(
					new Konva.Rect({
						width: cellSize,
						height: cellSize,
						x: x * cellSize,
						y: y * cellSize,
						fill: color
					})
				);
			}
		}

		gridLayer.draw();
	}

	function drawApples(apples: Vector2[], cellSize: number) {
		applesLayer.removeChildren();
		for (let i = 0; i < apples.length; i++) {
			const apple = apples[i];
			applesLayer.add(renderApple(apple, cellSize));
		}
		applesLayer.draw();
	}

	function drawPlayers(players: Player[], cellSize: number, moveValue: number) {
		snakesLayer.removeChildren();
		for (let i = 0; i < players.length; i++) {
			const player = players[i];
			const snake = renderSnake(player.snakeTail, player.direction, cellSize, moveValue, {
				color: player.color,
				name: player.name
			});
			snakesLayer.add(snake);
		}
		snakesLayer.draw();
	}
</script>

<div bind:this={canvasWrapper} />
