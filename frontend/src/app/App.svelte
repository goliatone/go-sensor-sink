<script>
	import Chart from './Chart.svelte';
	import moment from 'moment';
	import makeWebsocketStore from './stores/ws';

	let tChart;
	let hChart;
	let initialValue = {};

	export let version;
	export let environment;
	
	function makeWsUrl() {
		const qs = new URLSearchParams(location.search);
        const userId = qs.get('user_id');
		let url = `ws://localhost:3131/ws?user_id=${userId}`;
		console.log('url', url);
		return url;
	}

	const wsStore = makeWebsocketStore(makeWsUrl(),[],  initialValue);
	

	let tData = [{t: 1595219532618, y: 27 }];
	let hData = [{t: 1595219532618, y: 27 }];

	wsStore.subscribe(value => {
		if(!value || !value.time || !tChart) return;

		tData.push({
			t: parseInt(Math.abs(moment(value.time).format('x'))),
            y: value.t
		});

		tChart.updateData(tData);

		hData.push({
			t: parseInt(Math.abs(moment(value.time).format('x'))),
            y: value.h
		});

		hChart.updateData(hData);
	});
</script>

<main>
	<h3>Humidity & Temperature Dashboard</h3>
	<p>Simple dashboard showing data collected using ESP8266 and DHT22 sensors.</p>
	
	<Chart 
		chartId={"tChart"} 
		label="Temperature °C"
		borderColor="rgb(255, 99, 132)" 
		backgroundColor="rgb(255, 159, 64)" 
		bind:this={tChart}
	/>

	<Chart 
		chartId={"hChart"} 
		label="Humidity %"
		borderColor="rgb(54, 162, 235)" 
		backgroundColor="rgb(153, 102, 255)" 
		bind:this={hChart}
	/>

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