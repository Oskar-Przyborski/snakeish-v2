<script lang="ts">
    export let value: string = "";
</script>

<div class="text-input">
	<input type="text" required bind:value={value} />
	<div class="name"><slot /></div>
</div>

<style lang="scss">
	.text-input {
		margin: 0.9em 0;
		position: relative;
		height: 36px;

		input {
			all: unset;
			position: absolute;
			width: 94%;
			height: 100%;
			color: black;

			padding: 0 3%;
			background-color: #f2f2f2;
			border-radius: 0.3em;
		}

		input:focus ~ .name,
		input:valid ~ .name {
			transform: translate(3px, -110%);
			font-size: 0.95em;
			color: white;
		}

		.name {
			position: absolute;
			color: #393939;
			transform: translate(10px, 35%);
			transition: transform 0.2s, font-size 0.2s, color 0.2s;
		}
	}
</style>
