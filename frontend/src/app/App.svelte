<script>
	import Chart from "./Chart.svelte";

	var randomScalingFactor = function() {
        return Math.round(Math.random() * 100);
	};
	
	let data = [
		randomScalingFactor(),
		randomScalingFactor(),
		randomScalingFactor(),
		randomScalingFactor(),
		randomScalingFactor(),
	];

	function updateSourceData() {
		console.log('update data source');
		data = [
			randomScalingFactor(),
			randomScalingFactor(),
			randomScalingFactor(),
			randomScalingFactor(),
			randomScalingFactor(),
		]	
	}

	export let name;
	export let version;
	export let environment;
</script>

<main>
	<h3>Humidity And Temperature Dashboard</h3>
	<p>Simple dashboard showing data collected using ESP8266 and DHT22 sensors.</p>
	<Chart {data}/>
	<button on:click={_=>updateSourceData()}>Update Source Data</button>
	<footer>
		<strong>{environment}-{version}</strong>
	</footer>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>