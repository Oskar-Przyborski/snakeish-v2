<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { get, type Writable } from 'svelte/store';
	import Konva from 'konva';
	import type { Vector2 } from '$lib/vector';
	import type { ClassicGameState, PageState, Player } from './types';
	import { GameFramesRenderer } from '$lib/renderers/game_frames';
	import { renderSnake } from '$lib/renderers/snake';

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
				// console.log("frame update", moveValue)
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
		gameFrameRenderer?.setGameData(state.gameState)
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
				const color = darked ? '#2e3134' : '#36393c';
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
			const container = new Konva.Group({
				x: apple.x * cellSize,
				y: apple.y * cellSize
			});
			container.add(
				new Konva.Rect({
					x: cellSize * 0.125,
					y: cellSize * 0.125,
					width: cellSize * 0.75,
					height: cellSize * 0.75,
					fill: '#d62246',
					cornerRadius: cellSize * 0.2,
					shadowOpacity: 0.2,
					shadowOffsetY: 2
				}),
				new Konva.Path({
					x: cellSize * 0.55,
					y: cellSize * 0.28,
					data: 'M-0.5 0Q0.5 0 0.5-1 -0.5-1 -0.5 0',
					fill: 'green',
					scaleX: 17,
					scaleY: 17,
					shadowOpacity: 0.2,
					shadowOffsetY: 0.15
				})
			);
			applesLayer.add(container);
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
			})
			snakesLayer.add(snake);
		}
		snakesLayer.draw();
	}
</script>

<div bind:this={canvasWrapper} />
