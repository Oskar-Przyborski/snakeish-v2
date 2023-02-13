<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import type { Writable } from 'svelte/store';
	import Konva from 'konva';
	import { Add, Diff, Multiply, Vector2 } from '$lib/vector';
	import type { PageState, Player } from './types';
	import { FrameManager } from './classic_game';

	export let store: Writable<PageState>;
	let unsubscribe = () => {};
	onDestroy(unsubscribe);

	let canvasWrapper: HTMLDivElement;
	let stage: Konva.Stage;
	let gridLayer: Konva.Layer;
	let applesLayer: Konva.Layer;
	let snakesLayer: Konva.Layer;
	let wantedFrameTime = 1000 / 45;
	let snakesMoveFrame = 0;
	let snakesMoveInterval: NodeJS.Timeout;
	let frameManager: FrameManager;

	onMount(() => {
		initKonva();
		unsubscribe = store.subscribe(updateState);
		frameManager = new FrameManager()
	});

	async function updateState(state: PageState) {
		if (state.gameState == null) return;
		const cellSize = stage.width() / state.gameState.gridSize;

		drawGrid(state.gameState.gridSize, cellSize);
		drawApples(state.gameState.apples, cellSize);
		clearInterval(snakesMoveInterval);
		snakesMoveFrame = 0;
		snakesMoveInterval = setInterval(() => {
			frameManager.registerFrame()

			if (state.gameState == null) return;
			snakesMoveFrame++;
			const framesPerMove =
				(state.gameState?.frameTime / wantedFrameTime) * (wantedFrameTime / frameManager.avgDelta());
			const moveValue = snakesMoveFrame / framesPerMove;
			drawPlayers(state.gameState?.players, cellSize, moveValue);
		}, 1000 / wantedFrameTime);
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
		// console.log(moveValue)
		//draw snake
		for (let i = 0; i < players.length; i++) {
			const player = players[i];
			if (player.snakeTail.length == 0) return;

			const adjustedSnake = [...player.snakeTail];

			const adjustedHeadPos = Add(Multiply(player.direction, moveValue - 0.9), player.snakeTail[0]);
			adjustedSnake[0] = adjustedHeadPos;

			if (player.snakeTail.length > 1) {
				const last = player.snakeTail[player.snakeTail.length - 1];
				const beforeLast = player.snakeTail[player.snakeTail.length - 2];
				const adjustedBackPos = Add(last, Multiply(Diff(beforeLast, last), moveValue - 0.3));
				adjustedSnake[adjustedSnake.length - 1] = adjustedBackPos;
			}

			if (player.snakeTail.length > 1) {
				const points: number[] = [];
				adjustedSnake.forEach((el: Vector2) =>
					points.push(el.x * cellSize + cellSize / 2, el.y * cellSize + cellSize / 2)
				);
				snakesLayer.add(
					new Konva.Line({
						points,
						bezier: false,
						stroke: player.color,
						strokeWidth: cellSize * 0.65,
						closed: false,
						lineCap: 'round',
						lineJoin: 'round',
						shadowOffsetY: 3,
						shadowOpacity: 0.2
					})
				);
			} else {
				snakesLayer.add(
					new Konva.Circle({
						x: adjustedHeadPos.x * cellSize + cellSize / 2,
						y: adjustedHeadPos.y * cellSize + cellSize / 2,
						width: cellSize * 0.75,
						height: cellSize * 0.75,
						fill: player.color,
						shadowOffsetY: 3,
						shadowOpacity: 0.2
					})
				);
			}
			const nameContainer = new Konva.Group({
				x: adjustedHeadPos.x * cellSize,
				y: adjustedHeadPos.y * cellSize
			});
			const nameText = new Konva.Text({
				width: cellSize,
				height: cellSize,
				text: player.name,
				fill: 'white',
				fontFamily: 'DM Sans',
				fontStyle: 'bold',
				fontSize: cellSize / 4.5,
				align: 'center',
				verticalAlign: 'middle',
				shadowBlur: 2,
				shadowOffsetY: 1
			});
			nameContainer.add(nameText);
			snakesLayer.add(nameContainer);
		}
		snakesLayer.draw();
	}
</script>

<div bind:this={canvasWrapper} />
