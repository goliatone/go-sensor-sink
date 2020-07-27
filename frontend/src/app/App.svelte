<script>
	import {onMount} from 'svelte';
	import Chart from './Chart.svelte';
	import moment from 'moment';
	import KalmanFilter from 'kalmanjs';
	import Api from './services/api';
	import makeWebsocketStore from './stores/ws';

	let api = new Api();
	api.get();
	let chart;
	let initialValue = {};

	let datasets = [{
		label:'Temperature °C',
		borderColor:'rgb(255, 99, 132)',
		backgroundColor:'rgb(255, 159, 64)',
		data:[],
		type: 'line',
		pointRadius: 0,
		fill: false,
		lineTension: 0,
		borderWidth: 2
	},{
		label:'Humidity %',
		borderColor:'rgb(54, 162, 235)',
		backgroundColor:'rgb(153, 102, 255)',
		data:[],
		type: 'line',
		pointRadius: 0,
		fill: false,
		lineTension: 0,
		borderWidth: 2
	}];
	
	var kfilter = new KalmanFilter({R: 0.01, Q: 3});

	export let version;
	export let environment;
	
	function makeWsUrl() {
		const qs = new URLSearchParams(location.search);
		//TODO: This should come form template renderer, not user!
		const userId = qs.get('user_id');
		const host = qs.get('wshost') ? qs.get('wshost') : location.host;
		const protocol = location.protocol === 'https:' ? 'wss' : 'ws';
		
		const url = `${protocol}://${host}/ws?user_id=${userId}`;

		console.log('url', url);
		return url;
	}

	const wsStore = makeWebsocketStore(makeWsUrl(),[],  initialValue);
	

	let data = [
		[],
		[]
	];
	

	let tAvg = {
		sum: 0,
		tot: 0,
		add(v) {
			this.tot += 1;
			this.sum = Math.round((this.sum + v) / Math.min(this.tot, 2));
			return this.sum;
		}
	};

	let hAvg = {
		add(v) {
			return Math.round(kfilter.filter(v));
		}
	};

	function maxPush(a, v, max=6) {
  		a.push(v);
  		if(a.length > max){ 
			a.shift();
		}
	}

	wsStore.subscribe(value => {
		if(!value || !value.time || !chart) return;

		//TODO: Should make average for each device, maybe use kalman filter
		maxPush(data[0], {
			t: parseInt(Math.abs(moment(value.time).format('x'))),
            y: tAvg.add(value.t)
		}, 50);

		maxPush(data[1], {
			t: parseInt(Math.abs(moment(value.time).format('x'))),
            y: hAvg.add(value.h)
		}, 50);

		chart.updateData(data);
	});

	onMount(_=>{
		window.chart = chart;
	});
</script>

<main>
	<h3>Humidity & Temperature Dashboard</h3>
	<p>Simple dashboard showing data collected using ESP8266 and DHT22 sensors.</p>
	
	<Chart 
		chartId={"chart"} 
		{datasets}
		bind:this={chart}
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