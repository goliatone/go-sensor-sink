<script>
    import {onMount, afterUpdate} from 'svelte';
    import moment from 'moment';
    import Chart from 'chart.js';

    const chartColors = {
        red: 'rgb(255, 99, 132)',
        orange: 'rgb(255, 159, 64)',
        yellow: 'rgb(255, 205, 86)',
        green: 'rgb(75, 192, 192)',
        blue: 'rgb(54, 162, 235)',
        purple: 'rgb(153, 102, 255)',
        grey: 'rgb(201, 203, 207)'
    };

    let ctx;
    let chart;
    let unit;

    let config = {
        data: {
            datasets: [
            // {
            //     label: 'Humidity',
            //     borderColor: chartColors.red,
            //     backgroundColor: chartColors.orange,
            //     data: getReadingsData(readings, 'humidity'),
            //     type: 'line',
            //     pointRadius: 0,
            //     fill: false,
            //     lineTension: 0,
            //     borderWidth: 2
            // },
            {
                label: 'Temperature',
                borderColor: chartColors.blue,
                backgroundColor: chartColors.purple,
                data:[{
                    t:0,
                    y:70
                }],
                type: 'line',
                pointRadius: 0,
                fill: false,
                lineTension: 0,
                borderWidth: 2
            }]
        },
        options: {
            animation: {
                duration: 0
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    distribution: 'linear',//data are spread according to their time (distances can vary)
                    distribution: 'series',//data are spread at the same distance from each other
                    // bounds: 'ticks',
                    bounds: 'data',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'h:mm:ss a'
                        }
                    },
                    offset: true,
                    ticks: {
                        major: {
                            enabled: true,
                            fontStyle: 'bold'
                        },
                        source: 'data',
                        // source: 'labels',
                        // source: 'auto',
                        autoSkip: true,
                        autoSkipPadding: 75,
                        maxRotation: 0,
                        sampleSize: 100
                    },
                    afterBuildTicks: function(scale, ticks) {
                        var majorUnit = scale._majorUnit;
                        var firstTick = ticks[0];
                        var i, ilen, val, tick, currMajor, lastMajor;

                        val = moment(ticks[0].value);
                        if ((majorUnit === 'minute' && val.second() === 0)
                                || (majorUnit === 'hour' && val.minute() === 0)
                                || (majorUnit === 'day' && val.hour() === 9)
                                || (majorUnit === 'month' && val.date() <= 3 && val.isoWeekday() === 1)
                                || (majorUnit === 'year' && val.month() === 0)) {
                            firstTick.major = true;
                        } else {
                            firstTick.major = false;
                        }
                        lastMajor = val.get(majorUnit);

                        for (i = 1, ilen = ticks.length; i < ilen; i++) {
                            tick = ticks[i];
                            val = moment(tick.value);
                            currMajor = val.get(majorUnit);
                            tick.major = currMajor !== lastMajor;
                            lastMajor = currMajor;
                        }
                        return ticks;
                    }
                }],
                yAxes: [{
                    gridLines: {
                        drawBorder: false
                    },
                    scaleLabel: {
                        display: true,
                        labelString: 'Temperature'
                    }
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
        ctx = document.getElementById('chart1').getContext('2d');
        ctx.canvas.width = 1000;
        ctx.canvas.height = 300;
    });

    function initializeChart() {
    
        /**
         * Attach the graph to the DOM
         */ 
        
        
        chart = new Chart(ctx, config);
    }

    /**
     * Handle button click for randomized dataset
     */ 
    function update() {
        function updateDateset(index, type) {
            var type = document.getElementById('type').value;
            var dataset = chart.config.data.datasets[index];
            dataset.type = type;
            // dataset.data = getReadingsData(readings, type);
        }
        // updateDateset(0, 'humitidy');
        // updateDateset(1, 'temperature');

        chart.update();
    }

    export function updateData(data) {
        // var type = document.getElementById('type').value;
        // dataset.type = type;

        var dataset = chart.config.data.datasets[0];
        dataset.data = data;
        chart.update();
    }
    

    afterUpdate(initializeChart);
</script>

<div id="canvas-holder" style="width:3;height=1;" >
    <canvas id="chart1"></canvas>
</div>

Chart Type:
<select id="type">
    <option value="line">Line</option>
    <option value="bar">Bar</option>
</select>
<select id="unit">
    <option value="second">Second</option>
    <option value="minute">Minute</option>
    <option value="hour">Hour</option>
    <option value="day" selected>Day</option>
    <option value="month">Month</option>
    <option value="year">Year</option>
</select>
<button id="update" on:click={_=>update()} >update</button>