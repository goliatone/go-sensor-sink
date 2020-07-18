<script>
    import {onMount, afterUpdate} from 'svelte';
    import moment from 'moment';
    import Chart from 'chart.js';

    export let data;

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

    function generateData() {

        function unitLessThanDay() {
            return unit === 'second' || unit === 'minute' || unit === 'hour';
        }

        function beforeNineThirty(date) {
            return date.hour() < 9 || (date.hour() === 9 && date.minute() < 30);
        }

        // Returns true if outside 9:30am-4pm on a weekday
        function outsideMarketHours(date) {
            if (date.isoWeekday() > 5) return true;
            
            if (unitLessThanDay() && (beforeNineThirty(date) || date.hour() > 16)) {
                return true;
            }
            
            return false;
        }

        function randomNumber(min, max) {
            return Math.random() * (max - min) + min;
        }

        function randomBar(date, lastClose) {
            var open = randomNumber(lastClose * 0.95, lastClose * 1.05).toFixed(2);
            var close = randomNumber(open * 0.95, open * 1.05).toFixed(2);
            return {
                t: date.valueOf(),
                y: close
            };
        }

        var date = moment('Jan 01 1990', 'MMM DD YYYY');
        var now = moment();
        var data = [];
        var lessThanDay = unitLessThanDay();
        for (; data.length < 600 && date.isBefore(now); date = date.clone().add(1, unit).startOf(unit)) {
            if (outsideMarketHours(date)) {
                if (!lessThanDay || !beforeNineThirty(date)) {
                    date = date.clone().add(date.isoWeekday() >= 5 ? 8 - date.isoWeekday() : 1, 'day');
                }
                if (lessThanDay) {
                    date = date.hour(9).minute(30).second(0);
                }
            }
            data.push(randomBar(date, data.length > 0 ? data[data.length - 1].y : 30));
        }

        return data;
    }

    let config = {
        data: {
            datasets: [{
                label: 'CHRT - Chart.js Corporation',
                borderColor: chartColors.red,
                backgroundColor: chartColors.orange,
                data: generateData(),
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
                    distribution: 'series',
                    offset: true,
                    ticks: {
                        major: {
                            enabled: true,
                            fontStyle: 'bold'
                        },
                        source: 'data',
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
                        labelString: 'Closing price ($)'
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
        var type = document.getElementById('type').value;
        var dataset = chart.config.data.datasets[0];
        dataset.type = type;
        dataset.data = generateData();
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