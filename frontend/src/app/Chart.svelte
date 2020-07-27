<svelte:options accessors/>
<script>
    import {onMount, afterUpdate} from 'svelte';
    import moment from 'moment';
    import 'chartjs-plugin-streaming';
    import Chart from 'chart.js';

    export let chartId;
    export let label;
    export let borderColor; 
    export let backgroundColor;
    //for development only!
    export let chart;

    let ctx;
    // let chart;
    let type;
    let unit;

    export let datasets = []

    let config = {
        data: {
            datasets,
        },
        options: {
            animation: {
                duration: 0
            },
            scales: {
                // xAxes: [{
                //     type: 'time',
                //     distribution: 'linear',//data are spread according to their time (distances can vary)
                //     distribution: 'series',//data are spread at the same distance from each other
                //     // bounds: 'ticks',
                //     bounds: 'data',
                //     time: {
                //         unit: 'second',
                //         displayFormats: {
                //             quarter: 'h:mm:ss a'
                //         }
                //     },
                //     offset: true,
                //     ticks: {
                //         major: {
                //             enabled: true,
                //             fontStyle: 'bold'
                //         },
                //         source: 'data',
                //         // source: 'labels',
                //         // source: 'auto',
                //         autoSkip: true,
                //         autoSkipPadding: 75,
                //         maxRotation: 0,
                //         sampleSize: 100
                //     },
                //     afterBuildTicks: function(scale, ticks) {
                //         var majorUnit = scale._majorUnit;
                //         var firstTick = ticks[0];
                //         var i, ilen, val, tick, currMajor, lastMajor;

                //         val = moment(ticks[0].value);
                //         if ((majorUnit === 'minute' && val.second() === 0)
                //                 || (majorUnit === 'hour' && val.minute() === 0)
                //                 || (majorUnit === 'day' && val.hour() === 9)
                //                 || (majorUnit === 'month' && val.date() <= 3 && val.isoWeekday() === 1)
                //                 || (majorUnit === 'year' && val.month() === 0)) {
                //             firstTick.major = true;
                //         } else {
                //             firstTick.major = false;
                //         }
                //         lastMajor = val.get(majorUnit);

                //         for (i = 1, ilen = ticks.length; i < ilen; i++) {
                //             tick = ticks[i];
                //             val = moment(tick.value);
                //             currMajor = val.get(majorUnit);
                //             tick.major = currMajor !== lastMajor;
                //             lastMajor = currMajor;
                //         }
                //         return ticks;
                //     }
                // }],
                xAxes: [{
						type: 'realtime',
						realtime: {
							duration: 20000,
							refresh: 3000,
							delay: 2000,
							// onRefresh: onRefresh
						}
					}],
                yAxes: [{
                    gridLines: {
                        drawBorder: false
                    },
                    // scaleLabel: {
                    //     display: true,
                    //     labelString: scaleLabel
                    // }
                }]
            },
            tooltips: {
                intersect: false,
                mode: 'index',
                callbacks: {
                    label: function(tooltipItem, myData) {
                        var label = myData.datasets[tooltipItem.datasetIndex].label || '';
                        if (label) {
                            label += ': ';
                        }
                        label += parseFloat(tooltipItem.value).toFixed(2);
                        return label;
                    }
                }
            }
        }
    };

    window.config = config;
    window.moment = moment;

    onMount(_=>{
        unit = document.getElementById('unit').value;
        ctx = document.getElementsByClassName(chartId)[0].getContext('2d');
        ctx.canvas.width = 1000;
        ctx.canvas.height = 300;
    });

    function initializeChart() {
        chart = new Chart(ctx, config);
    }

    /**
     * Handle button click for randomized dataset
     */ 
    function update() {
        var type = document.getElementById('type').value;
        var dataset = chart.config.data.datasets[0];
        dataset.type = type;

        chart.update();
    }

    export function updateData(data) {
        // var type = document.getElementById('type').value;
        // dataset.type = type;

        chart.config.data.datasets.forEach((dataset, i)=>{
            dataset.data = data[i];
        });

        chart.update();
    }
    
    afterUpdate(initializeChart);

    $: if(chart && type) {
        chart.config.data.datasets.forEach(dataset=>{
            dataset.type = type;
        });
        
    }
</script>

<div id="canvas-holder" style="width:3;height=1;" >
    <canvas class="{chartId}"></canvas>
</div>

Chart Type:
<select id="type" bind:value={type}>
    <option value="line">Line</option>
    <option value="bar">Bar</option>
</select>
<select id="unit" bind:value={unit}>
    <option value="second">Second</option>
    <option value="minute">Minute</option>
    <option value="hour">Hour</option>
    <option value="day" selected>Day</option>
    <option value="month">Month</option>
    <option value="year">Year</option>
</select>
<button id="update" on:click={_=>update()} >Update</button>